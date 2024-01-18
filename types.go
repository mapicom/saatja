package saatja

type SaatjaHeader struct {
	Destination [4]byte
	Source      [4]byte
	Length      uint16
	Checksum    byte
}

type SaatjaPacket struct {
	SaatjaHeader
	Content []byte
}
