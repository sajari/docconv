package ole2

import ()

type PSS struct {
	name      [64]byte
	bsize     uint16
	typ       byte
	flag      byte
	left      uint32
	right     uint32
	child     uint32
	guid      [16]uint16
	userflags uint32
	time      [2]uint64
	sstart    uint32
	size      uint32
	_         uint32
}
