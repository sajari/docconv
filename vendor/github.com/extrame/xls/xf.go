package xls

type Xf5 struct {
	Font      uint16
	Format    uint16
	Type      uint16
	Align     uint16
	Color     uint16
	Fill      uint16
	Border    uint16
	Linestyle uint16
}

func (x *Xf5) formatNo() uint16 {
	return x.Format
}

type Xf8 struct {
	Font        uint16
	Format      uint16
	Type        uint16
	Align       byte
	Rotation    byte
	Ident       byte
	Usedattr    byte
	Linestyle   uint32
	Linecolor   uint32
	Groundcolor uint16
}

func (x *Xf8) formatNo() uint16 {
	return x.Format
}

type st_xf_data interface {
	formatNo() uint16
}
