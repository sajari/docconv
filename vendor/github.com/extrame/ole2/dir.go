package ole2

import (
	"unicode/utf16"
)

const (
	EMPTY       = iota
	USERSTORAGE = iota
	USERSTREAM  = iota
	LOCKBYTES   = iota
	PROPERTY    = iota
	ROOT        = iota
)

type File struct {
	NameBts   [32]uint16
	Bsize     uint16
	Type      byte
	Flag      byte
	Left      uint32
	Right     uint32
	Child     uint32
	Guid      [8]uint16
	Userflags uint32
	Time      [2]uint64
	Sstart    uint32
	Size      uint32
	Proptype  uint32
}

func (d *File) Name() string {
	runes := utf16.Decode(d.NameBts[:d.Bsize/2-1])
	return string(runes)
}
