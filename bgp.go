package main

import (
	"errors"
	"log"
)

type BgpMessageType uint8

const (
	OPEN         BgpMessageType = 1
	UPDATE       BgpMessageType = 2
	NOTIFICATION BgpMessageType = 3
	KEEPALIVE    BgpMessageType = 4
)

type BgpMessageHeader struct {
	marker        [16]byte
	messageLength uint16
	messageType   BgpMessageType
}

func NewBgpMessageHeader(messageLength uint16, messageType BgpMessageType) (*BgpMessageHeader, error) {
	if messageLength < 19 || messageLength > 4096 {
		log.Fatalf("ERROR invalid BGP message header length --> %d", messageLength)
		return nil, errors.New("the value of the Length field MUST always be at least 19 and no greater than 4096")
	}

	if messageType != OPEN && messageType != UPDATE && messageType != NOTIFICATION && messageType != KEEPALIVE {
		log.Fatalf("ERROR invalid BGP message type %s", messageType)
		return nil, errors.New("invalid BGP message type")
	}

	var marker [16]byte
	for i := range marker {
		marker[i] = 0xFF
	}

	return &BgpMessageHeader{
		marker:        marker,
		messageLength: messageLength,
		messageType:   messageType,
	}, nil
}

type BgpOptionalParameter struct {
	parameterType   uint8
	parameterLength uint8
	parameterValue  []byte
}

func NewBgpOptionalParameters(optionalParametersLength uint8, buf []byte) ([]*BgpOptionalParameter, error) {
	var optionalParameters []*BgpOptionalParameter
	for i := range buf {
		t := buf[0]
		l := buf[1]
		v := buf[2 : 2+l]
		optionalParameters = append(optionalParameters, &BgpOptionalParameter{t, l, v})
		log.Printf("INFO NewBgpOptionalParameters BGP optional parameters %v", optionalParameters)
		buf = buf[i+2+int(l):]
		optionalParametersLength -= 1
		if optionalParametersLength < 1 {
			break
		}
	}
	return optionalParameters, nil
}

type BgpOpenMessage struct {
	messageHeader        *BgpMessageHeader
	bgpVersion           uint8
	asNumber             uint16
	holdTime             uint16
	bgpIdentifier        uint32
	optionalParamsLength uint8
	optionalParameters   []*BgpOptionalParameter
}

func NewBgpOpenMessage(messageHeader *BgpMessageHeader, bgpVersion uint8, asNumber uint16, holdTime uint16, bgpIdentifier uint32, optionalParamsLength uint8, optionalParameters []*BgpOptionalParameter) (*BgpOpenMessage, error) {
	return &BgpOpenMessage{
		messageHeader:        messageHeader,
		bgpVersion:           bgpVersion,
		asNumber:             asNumber,
		holdTime:             holdTime,
		bgpIdentifier:        bgpIdentifier,
		optionalParamsLength: optionalParamsLength,
		optionalParameters:   optionalParameters,
	}, nil
}
