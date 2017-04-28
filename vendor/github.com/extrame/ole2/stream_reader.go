package ole2

import (
	"io"
	"log"
)

var DEBUG = false

type StreamReader struct {
	sat              []uint32
	start            uint32
	reader           io.ReadSeeker
	offset_of_sector uint32
	offset_in_sector uint32
	size_sector      uint32
	size             int64
	offset           int64
	sector_pos       func(uint32, uint32) uint32
}

func (r *StreamReader) Read(p []byte) (n int, err error) {
	if r.offset_of_sector == ENDOFCHAIN {
		return 0, io.EOF
	}
	pos := r.sector_pos(r.offset_of_sector, r.size_sector) + r.offset_in_sector
	r.reader.Seek(int64(pos), 0)
	readed := uint32(0)
	for remainLen := uint32(len(p)) - readed; remainLen > r.size_sector-r.offset_in_sector; remainLen = uint32(len(p)) - readed {
		if n, err := r.reader.Read(p[readed : readed+r.size_sector-r.offset_in_sector]); err != nil {
			return int(readed) + n, err
		} else {
			readed += uint32(n)
			r.offset_in_sector = 0
			if r.offset_of_sector >= uint32(len(r.sat)) {
				log.Fatal(`
				THIS SHOULD NOT HAPPEN, IF YOUR PROGRAM BREAK, 
				COMMENT THIS LINE TO CONTINUE AND MAIL ME XLS FILE 
				TO TEST, THANKS`)
				return int(readed), io.EOF
			} else {
				r.offset_of_sector = r.sat[r.offset_of_sector]
			}
			if r.offset_of_sector == ENDOFCHAIN {
				return int(readed), io.EOF
			}
			pos := r.sector_pos(r.offset_of_sector, r.size_sector) + r.offset_in_sector
			r.reader.Seek(int64(pos), 0)
		}
	}
	if n, err := r.reader.Read(p[readed:len(p)]); err == nil {
		r.offset_in_sector += uint32(n)
		if DEBUG {
			log.Printf("pos:%x,bit:% X", r.offset_of_sector, p)
		}
		return len(p), nil
	} else {
		return int(readed) + n, err
	}

}

func (r *StreamReader) Seek(offset int64, whence int) (offset_result int64, err error) {

	if whence == 0 {
		r.offset_of_sector = r.start
		r.offset_in_sector = 0
		r.offset = offset
	} else {
		r.offset += offset
	}

	if r.offset_of_sector == ENDOFCHAIN {
		return r.offset, io.EOF
	}

	for offset >= int64(r.size_sector-r.offset_in_sector) {
		r.offset_of_sector = r.sat[r.offset_of_sector]
		offset -= int64(r.size_sector - r.offset_in_sector)
		r.offset_in_sector = 0
		if r.offset_of_sector == ENDOFCHAIN {
			err = io.EOF
			goto return_res
		}
	}

	if r.size <= r.offset {
		err = io.EOF
		r.offset = r.size
	} else {
		r.offset_in_sector += uint32(offset)
	}
return_res:
	offset_result = r.offset
	return
}
