package main

import (
	"encoding/binary"
	"encoding/hex"
	"fmt"
	"math"
	"math/bits"
)

func main() {
	lines := testInput()
	fmt.Println("Part 1")
	part1(lines)
}

func part1(lines []string) {
	for _, line := range lines {
		fmt.Println("---------")
		fmt.Println(line)
		processPacket([]byte(line), 0)
	}
}

func processPacket(line []byte, startPos int) (newPos int) {
	p, newPos := processPacketHeader(line, startPos)
	fmt.Println("")
	fmt.Println("Version", p.version, "and type", p.typeId)
	switch p.typeId {
	case 4:
		fmt.Println("* Literal packet")
		var literalBytes []byte
		literalBytes, newPos = processLiteral(p, newPos)
		fmt.Println(literalBytes)
	default:
		fmt.Println("* Operator packet")
		newPos = processOperator(p, newPos)
	}
	return
}

func processOperator(packet tPacket, pos int) int {
	lengthType, pos := extractBits8(packet.contents, pos, 1)
	fmt.Println("length type", lengthType)
	switch lengthType {
	case 0:
		var length uint16
		length, pos = extractBits16(packet.contents, pos, 15)
		fmt.Println("Length of subs", length, "and pos is", pos)

		endPoint := int(length) + pos
		fmt.Println("Endpoint is", endPoint)

		for pos < endPoint {
			fmt.Println("Reading packet from", pos)
			pos = processPacket(packet.contents, pos)
			fmt.Println("After reading packet, pos is", pos)
		}

	case 1:
		var numSubs uint16
		numSubs, pos = extractBits16(packet.contents, pos, 11)
		fmt.Println("Num subs", numSubs)

		for numSubs > 0 {
			numSubs--
			fmt.Println("Processing sub packet...")
			pos = processPacket(packet.contents, pos)
		}

	}

	fmt.Println("Leaving operator packet, with pos", pos)
	return pos
}

func processLiteral(packet tPacket, pos int) ([]byte, int) {
	var b uint8
	var all []byte
	var okToContinue bool = true

	fmt.Println(" ... looking at literal from pos", pos)

	for okToContinue {
		b, pos = extractBits8(packet.contents, pos, 5)
		all = append(all, b)
		if b&0b00010000 > 0 {
			okToContinue = true
		} else {
			okToContinue = false
		}
	}

	/*
		// Align pos to a 4-bit boundary
		for pos%4 > 0 {
			pos++
		}
	*/

	return all, pos
}

type tPacket struct {
	version  uint8
	typeId   uint8
	contents []byte
}

type tPacket4 struct {
	tPacket
	literal []byte
}

func processPacketHeader(packet []byte, startPos int) (tPacket, int) {
	b, _ := hex.DecodeString(string(packet))
	p := tPacket{}

	pos := startPos
	p.version, pos = extractBits8(b, pos, 3)
	p.typeId, pos = extractBits8(b, pos, 3)
	p.contents = b
	return p, pos
}

func extractBits16(src []byte, pos int, numBits int) (retval uint16, newPos int) {
	newPos = pos + numBits

	if numBits > 16 {
		panic("Can only read up to 16 bits")
	}

	// Extract 4 bytes, for up to 16 bits across a boundary (should be 3 bytes but we only work in powers of 2)
	byteNum := int(pos / 8)
	modifiedPos := pos % 8
	var b uint32
	src = append(src, []byte{0, 0, 0, 0, 0, 0, 0}...)
	b = binary.BigEndian.Uint32(src[byteNum : byteNum+4])

	mask := uint32(math.Pow(2, float64(numBits))) - 1
	retval = uint16(uint32(bits.RotateLeft32(b, -1*(32-(numBits+modifiedPos)))) & mask)

	/*
		fmt.Printf("Src %#v\n", src)
		fmt.Printf("byteNum %#v\n", byteNum)
		fmt.Printf("b %#v\n", b)
		fmt.Printf("mask %#v\n", mask)
		fmt.Printf("retval %#v\n", retval)
	*/
	return
}

func extractBits8(src []byte, pos int, numBits int) (retval uint8, newPos int) {
	newPos = pos + numBits

	if numBits > 8 {
		panic("Can only read up to 8 bits")
	}

	// Extract 2 bytes, for up to 8 bits across a boundary
	byteNum := int(pos / 8)
	modifiedPos := pos % 8
	var b uint16
	if len(src) < byteNum+2 {
		src = append(src, []byte{0, 0}...)
	}
	b = binary.BigEndian.Uint16(src[byteNum : byteNum+2])

	mask := uint16(math.Pow(2, float64(numBits))) - 1
	retval = uint8(uint16(bits.RotateLeft16(b, -1*(16-(numBits+modifiedPos)))) & mask)

	/*
		fmt.Printf("Src %#v\n", src)
		fmt.Printf("byteNum %#v\n", byteNum)
		fmt.Printf("b %#v\n", b)
		fmt.Printf("mask %#v\n", mask)
		fmt.Printf("retval %#v\n", retval)
	*/
	return
}

func testInput() []string {
	return []string{
		"D2FE28",
		"38006F45291200",
		"EE00D40C823060",
		"8A004A801A8002F478",
		"620080001611562C8802118E34",
		"C0015000016115A2E0802F182340",
		"A0016C880162017C3686B18A3D4780",
	}
}
