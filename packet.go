package saatja

import (
	"encoding/binary"
	"errors"
	"math"

	"github.com/sigurn/crc8"
)

func CreatePacket(source, destination [4]byte, content []byte) ([]byte, error) {
	if float64(len(content)) > math.Pow(2, 16) {
		return []byte{}, errors.New("Bad content's length")
	}
	packet := []byte{}
	// Destintion
	packet = append(packet, destination[0], destination[1], destination[2], destination[3])
	// Source
	packet = append(packet, source[0], source[1], source[2], source[3])
	// Length
	packet = append(packet, byte(len(content)&0xFF), byte((len(content)>>8)&0xFF))
	// Checksum
	table := crc8.MakeTable(crc8.CRC8_MAXIM)
	packet = append(packet, crc8.Checksum(content, table))
	// Content
	packet = append(packet, content...)
	return packet, nil
}

func ParsePacket(content []byte) (*SaatjaPacket, error) {
	if len(content) < 11 {
		return nil, errors.New("Bad packet's length")
	}
	packet := SaatjaPacket{}
	table := crc8.MakeTable(crc8.CRC8_MAXIM)
	if packet.Checksum = crc8.Checksum(content[11:], table); packet.Checksum != content[10] {
		return nil, errors.New("Bad checksum")
	}
	packet.Destination = [4]byte(content[0:4])
	packet.Source = [4]byte(content[4:8])
	packet.Length = binary.LittleEndian.Uint16(content[8:10])
	packet.Content = content[11:]
	return &packet, nil
}
