package main

import (
	"bytes"
	"encoding/hex"
	"fmt"
	"github.com/magicmonkey/adventofcode/2021/util"
	"math"
)

func main() {
	//lines := testInput()
	lines := util.ReadInputFile()
	fmt.Println("Parts 1 and 2")
	part1(lines)
}

type tPacket struct {
	version uint64
	typeId  uint64
	value   uint64
	op      uint64
	packets []tPacket
}

var versionSum uint64

func part1(lines []string) {
	for _, line := range lines {
		fmt.Println("---------")
		//fmt.Println(line)

		versionSum = 0

		content, err := hex.DecodeString(line)
		if err != nil {
			panic(err)
		}
		b := bytes.NewBuffer(content)
		p, _ := processPacket(b, 0)
		fmt.Println("=== Version sum", versionSum)

		fmt.Println("=== Value", evaluate(p))
	}
}

func evaluate(packet tPacket) (val uint64) {
	switch packet.typeId {

	case 0: // Sum
		for _, p := range packet.packets {
			val += evaluate(p)
		}

	case 1: // Multiply
		val = 1
		for _, p := range packet.packets {
			v := evaluate(p)
			val *= v
		}

	case 2: // Minimum
		val = math.MaxInt64
		for _, p := range packet.packets {
			v := evaluate(p)
			if v < val {
				val = v
			}
		}

	case 3: // Maximum
		val = 0
		for _, p := range packet.packets {
			v := evaluate(p)
			if v > val {
				val = v
			}
		}

	case 4: // Literal value
		val = packet.value

	case 5: // Greater than
		v1 := evaluate(packet.packets[0])
		v2 := evaluate(packet.packets[1])
		if v1 > v2 {
			val = 1
		} else {
			val = 0
		}

	case 6: // Less than
		v1 := evaluate(packet.packets[0])
		v2 := evaluate(packet.packets[1])
		if v1 < v2 {
			val = 1
		} else {
			val = 0
		}

	case 7: // Equal to
		v1 := evaluate(packet.packets[0])
		v2 := evaluate(packet.packets[1])
		if v1 == v2 {
			val = 1
		} else {
			val = 0
		}

	}
	return
}

func processPacket(b *bytes.Buffer, startPos uint32) (packet tPacket, pos uint32) {
	pos = startPos
	var version, typeId uint64
	version, pos = extractBits(b.Bytes(), pos, 3)
	typeId, pos = extractBits(b.Bytes(), pos, 3)
	versionSum += version

	switch typeId {
	case 4:
		packet, pos = processLiteral(b, pos)
	default:
		packet, pos = processOperator(b, pos)
	}

	packet.version = version
	packet.typeId = typeId
	return
}

func processOperator(b *bytes.Buffer, startPos uint32) (packet tPacket, pos uint32) {
	pos = startPos
	lengthType, pos := extractBits(b.Bytes(), pos, 1)

	var p tPacket
	switch lengthType {

	case 0:
		var totalLength uint64
		totalLength, pos = extractBits(b.Bytes(), pos, 15)
		end := pos + uint32(totalLength)
		for pos < end {
			p, pos = processPacket(b, pos)
			packet.packets = append(packet.packets, p)

		}

	case 1:
		var numSubs uint64
		numSubs, pos = extractBits(b.Bytes(), pos, 11)
		for numSubs > 0 {
			numSubs--
			p, pos = processPacket(b, pos)
			packet.packets = append(packet.packets, p)

		}
	}

	return
}

func processLiteral(b *bytes.Buffer, startPos uint32) (packet tPacket, pos uint32) {
	pos = startPos
	var val uint64
	okToContinue := true
	for okToContinue {
		val = val << 4
		var nextBits uint64
		nextBits, pos = extractBits(b.Bytes(), pos, 5)
		if nextBits&0b00010000 > 0 {
			okToContinue = true
		} else {
			okToContinue = false
		}
		val |= (nextBits & 0b00001111)
	}
	packet.value = val
	return
}

func extractBits(bytes []byte, bit_start uint32, length uint32) (bits uint64, endPos uint32) {
	if length > 64 {
		panic("length too long" + string(length))
	}
	var res uint64 = 0
	for i := bit_start; i < bit_start+length; i++ {
		b, _ := extractBit(bytes, i)
		if b {
			res = 2*res + 1
		} else {
			res = res * 2
		}
	}
	return res, bit_start + length
}

func extractBit(b []byte, i uint32) (bool, uint32) {
	idx, offset := (i / 8), (i % 8)
	return (b[idx] & (1 << uint(7-offset))) != 0, i + 1
}

func testInput() []string {
	return []string{
		/*
			"D2FE28",
			"38006F45291200",
			"EE00D40C823060",
			"8A004A801A8002F478",
			"620080001611562C8802118E34",
			"C0015000016115A2E0802F182340",
			"A0016C880162017C3686B18A3D4780",
		*/
		"C200B40A82",
		"04005AC33890",
		"880086C3E88112",
		"CE00C43D881120",
		"D8005AC2A8F0",
		"F600BC2D8F",
		"9C005AC2F8F0",
		"9C0141080250320F1802104A08",
	}
}
