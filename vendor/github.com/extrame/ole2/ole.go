package ole2

import (
	"encoding/binary"
	"io"
)

var ENDOFCHAIN = uint32(0xFFFFFFFE) //-2
var FREESECT = uint32(0xFFFFFFFF)   // -1

type Ole struct {
	header   *Header
	Lsector  uint32
	Lssector uint32
	SecID    []uint32
	SSecID   []uint32
	Files    []File
	reader   io.ReadSeeker
}

func Open(reader io.ReadSeeker, charset string) (ole *Ole, err error) {
	var header *Header
	var hbts = make([]byte, 512)
	reader.Read(hbts)
	if header, err = parseHeader(hbts); err == nil {
		ole = new(Ole)
		ole.reader = reader
		ole.header = header
		ole.Lsector = 512 //TODO
		ole.Lssector = 64 //TODO
		err = ole.readMSAT()
		return ole, err
	}

	return nil, err
}

func (o *Ole) ListDir() (dir []*File, err error) {
	sector := o.stream_read(o.header.Dirstart, 0)
	dir = make([]*File, 0)
	for {
		d := new(File)
		err = binary.Read(sector, binary.LittleEndian, d)
		if err == nil && d.Type != EMPTY {
			dir = append(dir, d)
		} else {
			break
		}
	}
	if err == io.EOF && dir != nil {
		return dir, nil
	}

	return
}

func (o *Ole) OpenFile(file *File, root *File) io.ReadSeeker {
	if file.Size < o.header.Sectorcutoff {
		return o.short_stream_read(file.Sstart, file.Size, root.Sstart)
	} else {
		return o.stream_read(file.Sstart, file.Size)
	}
}

// Read MSAT
func (o *Ole) readMSAT() error {
	// int sectorNum;

	count := uint32(109)
	if o.header.Cfat < 109 {
		count = o.header.Cfat
	}

	for i := uint32(0); i < count; i++ {
		if sector, err := o.sector_read(o.header.Msat[i]); err == nil {
			sids := sector.AllValues(o.Lsector)
			o.SecID = append(o.SecID, sids...)
		} else {
			return err
		}
	}

	for sid := o.header.Difstart; sid != ENDOFCHAIN; {
		if sector, err := o.sector_read(sid); err == nil {
			sids := sector.MsatValues(o.Lsector)

			for _, sid := range sids {
				if sector, err := o.sector_read(sid); err == nil {
					sids := sector.AllValues(o.Lsector)

					o.SecID = append(o.SecID, sids...)
				} else {
					return err
				}
			}

			sid = sector.NextSid(o.Lsector)
		} else {
			return err
		}
	}

	for i := uint32(0); i < o.header.Csfat; i++ {
		sid := o.header.Sfatstart

		if sid != ENDOFCHAIN {
			if sector, err := o.sector_read(sid); err == nil {
				sids := sector.MsatValues(o.Lsector)

				o.SSecID = append(o.SSecID, sids...)

				sid = sector.NextSid(o.Lsector)
			} else {
				return err
			}
		}
	}
	return nil

}

func (o *Ole) stream_read(sid uint32, size uint32) *StreamReader {
	return &StreamReader{o.SecID, sid, o.reader, sid, 0, o.Lsector, int64(size), 0, sector_pos}
}

func (o *Ole) short_stream_read(sid uint32, size uint32, startSecId uint32) *StreamReader {
	ssatReader := &StreamReader{o.SecID, startSecId, o.reader, sid, 0, o.Lsector, int64(uint32(len(o.SSecID)) * o.Lssector), 0, sector_pos}
	return &StreamReader{o.SSecID, sid, ssatReader, sid, 0, o.Lssector, int64(size), 0, short_sector_pos}
}

func (o *Ole) sector_read(sid uint32) (Sector, error) {
	return o.sector_read_internal(sid, o.Lsector)
}

func (o *Ole) short_sector_read(sid uint32) (Sector, error) {
	return o.sector_read_internal(sid, o.Lssector)
}

func (o *Ole) sector_read_internal(sid, size uint32) (Sector, error) {
	pos := sector_pos(sid, size)
	if _, err := o.reader.Seek(int64(pos), 0); err == nil {
		var bts = make([]byte, size)
		o.reader.Read(bts)
		return Sector(bts), nil
	} else {
		return nil, err
	}
}

func sector_pos(sid uint32, size uint32) uint32 {
	return 512 + sid*size
}

func short_sector_pos(sid uint32, size uint32) uint32 {
	return sid * size
}
