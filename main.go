package main

import (
	"encoding/binary"
	"errors"
	"log"
	"net"
)

func handleBgpOpenMessage(header *BgpMessageHeader, buf []byte) error {
	var bgpVersion uint8 = buf[0]
	var asNumber uint16 = binary.BigEndian.Uint16(buf[1:3])
	var holdTime uint16 = binary.BigEndian.Uint16(buf[3:5])
	var bgpIdentifier uint32 = binary.BigEndian.Uint32(buf[5:9])
	var optionalParametersLength uint8 = buf[9]
	var optionalParameters []*BgpOptionalParameter
	if optionalParametersLength > 0 {
		optionalParameters, _ = NewBgpOptionalParameters(optionalParametersLength, buf[10:])
		log.Printf("INFO parsing optional parameters %v", optionalParameters)
	}

	message, err := NewBgpOpenMessage(header, bgpVersion, asNumber, holdTime, bgpIdentifier, optionalParametersLength, optionalParameters)
	if err != nil {
		log.Printf("ERROR failed to parse BGP Open message %v", err)
		return err
	}

	log.Printf("INFO BGP Open Message %v", message)
	return nil
}

func handleSession(conn net.Conn) error {
	defer conn.Close()

	for {
		buf := make([]byte, 4096)
		_, err := conn.Read(buf)
		if err != nil {
			log.Printf("ERROR failed to read TCP connection data into buffer %v", err)
			return err
		}

		var messageLength uint16 = binary.BigEndian.Uint16(buf[16:18])
		messageType := BgpMessageType(buf[18])
		log.Printf("INFO BGP header message length %d message type %v", messageLength, messageType)

		header, err := NewBgpMessageHeader(messageLength, messageType)
		if err != nil {
			log.Printf("ERROR failed to parse incoming BGP message %v", err)
			return err
		}

		log.Printf("INFO successfully created new BGP message header %v", header)
		switch messageType {
		case OPEN:
			return handleBgpOpenMessage(header, buf[19:])
		case UPDATE:
			return nil
		case KEEPALIVE:
			return nil
		case NOTIFICATION:
			return nil
		default:
			err := errors.New("invalid message type found in BGP header")
			log.Printf("ERROR %v", err)
			return err
		}
	}
}

func startRouter(addr string) error {
	listener, err := net.Listen("tcp6", addr)
	if err != nil {
		log.Printf("ERROR failed to start tcp server on %v", err)
	}
	log.Printf("INFO successfully bound to address %v", listener.Addr())
	defer listener.Close()

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Printf("ERROR failed to accept connection %v", err)
			return err
		}
		go handleSession(conn)
	}
}

func main() {
	startRouter(":1179")
}
