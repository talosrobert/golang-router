#!/bin/bash

# BGP Open Message with 0 Optional Parameters
declare -r msg1="\xff\xff\xff\xff\xff\xff\xff\xff\xff\xff\xff\xff\xff\xff\xff\xff\x00\x1d\x01\x04\xfd\xe9\x00\x0a\x04\x03\x02\x01\x00"
# BGP Open Message with 1 Optional Parameters
declare -r msg2="\xff\xff\xff\xff\xff\xff\xff\xff\xff\xff\xff\xff\xff\xff\xff\xff\x00\x20\x01\x04\xfd\xe9\x00\x0a\x04\x03\x02\x01\x01\x01\x01\x01"
# BGP Open Message with 2 Optional Parameters
declare -r msg3="\xff\xff\xff\xff\xff\xff\xff\xff\xff\xff\xff\xff\xff\xff\xff\xff\x00\x23\x01\x04\xfd\xe9\x00\x0a\x04\x03\x02\x01\x02\x01\x01\x01\x01\x01\x01"
# BGP Open Message with 3 Optional Parameters
declare -r msg4="\xff\xff\xff\xff\xff\xff\xff\xff\xff\xff\xff\xff\xff\xff\xff\xff\x00\x26\x01\x04\xfd\xe9\x00\x0a\x04\x03\x02\x01\x03\x01\x01\x01\x01\x01\x01\x01\x01\x01"
# BGP Open Message with 4 Optional Parameters
declare -r msg5="\xff\xff\xff\xff\xff\xff\xff\xff\xff\xff\xff\xff\xff\xff\xff\xff\x00\x29\x01\x04\xfd\xe9\x00\x0a\x04\x03\x02\x01\x04\x01\x01\x01\x01\x01\x01\x01\x01\x01\x01\x01\x01"

declare -a msgs=(msg1 msg2 msg3 msg4 msg5)
for msg in "${msgs[@]}"; do
	echo -e "${!msg}" | nc -6 ::1 1179
done

