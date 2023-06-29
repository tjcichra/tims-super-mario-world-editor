package helper

import (
	"fmt"
)

func DecompressLZ2(compressedData []byte) []byte {
	var decompressedData []byte
	index := 0

	for index < len(compressedData) {
		headerByte := compressedData[index]

		if headerByte == 255 {
			break
		}

		commandBits := headerByte >> 5
		var length uint16

		if commandBits == 0b111 {
			commandBits = headerByte >> 2 & 0b111

			length1 := headerByte & 0b11
			index++
			length2 := compressedData[index]

			length = twoBytesToWord(length1, length2)
		} else {
			length = uint16(headerByte & 0b11111)
		}

		index++

		// index is now after the header byte(s)
		switch commandBits {
		case 0b000:
			// Direct Copy
			endingIndex := index + int(length+1)
			decompressedData = append(decompressedData, compressedData[index:endingIndex]...)
			index = endingIndex
		case 0b001:
			// Byte Fill
			byteToFill := compressedData[index]
			decompressedData = append(decompressedData, repeatedSlice(byteToFill, length+1)...)
			index++
		case 0b010:
			// Word Fill
			byte1ToFill := compressedData[index]
			index++
			byte2ToFill := compressedData[index]

			decompressedData = append(decompressedData, alternatingRepeatedSlice(byte1ToFill, byte2ToFill, length+1)...)
			index++
			//TODO
		case 0b011:
			// Increasing Fill
			byteToFill := compressedData[index]
			decompressedData = append(decompressedData, increasingRepeatedSlice(byteToFill, length+1)...)
			index++
		case 0b100:
			// repeat
			//TODO
		default:
			fmt.Println("Unknown command bit encountered when decompressing graphics")
		}
	}

	return decompressedData
}

func twoBytesToWord(mostSignificant byte, leastSignificant byte) uint16 {
	return uint16(mostSignificant)<<8 | uint16(leastSignificant)
}

func repeatedSlice(value byte, n uint16) []byte {
	arr := make([]byte, n)

	for i := 0; i < int(n); i++ {
		arr[i] = value
	}

	return arr
}

func alternatingRepeatedSlice(value1 byte, value2 byte, n uint16) []byte {
	arr := make([]byte, n)

	for i := 0; i < int(n); i++ {
		arrvalue := value1
		if i%2 == 1 {
			arrvalue = value2
		}

		arr[i] = arrvalue
	}

	return arr
}

func increasingRepeatedSlice(value byte, n uint16) []byte {
	arr := make([]byte, n)

	for i := 0; i < int(n); i, value = i+1, value+1 {
		arr[i] = value
	}

	return arr
}
