package ole2

type Stream struct {
	Ole     *Ole
	Start   uint32
	Pos     uint32
	Cfat    int
	Size    int
	Fatpos  uint32
	Bufsize uint32
	Eof     byte
	Sfat    bool
}
