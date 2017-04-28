package xls

import (
	"fmt"
)

// range type of multi rows
type Ranger interface {
	FirstRow() uint16
	LastRow() uint16
}

// range type of multi cells in multi rows
type CellRange struct {
	FirstRowB uint16
	LastRowB  uint16
	FristColB uint16
	LastColB  uint16
}

func (c *CellRange) FirstRow() uint16 {
	return c.FirstRowB
}

func (c *CellRange) LastRow() uint16 {
	return c.LastRowB
}

func (c *CellRange) FirstCol() uint16 {
	return c.FristColB
}

func (c *CellRange) LastCol() uint16 {
	return c.LastColB
}

//hyperlink type's content
type HyperLink struct {
	CellRange
	Description      string
	TextMark         string
	TargetFrame      string
	Url              string
	ShortedFilePath  string
	ExtendedFilePath string
	IsUrl            bool
}

//get the hyperlink string, use the public variable Url to get the original Url
func (h *HyperLink) String(wb *WorkBook) []string {
	res := make([]string, h.LastColB-h.FristColB+1)
	var str string
	if h.IsUrl {
		str = fmt.Sprintf("%s(%s)", h.Description, h.Url)
	} else {
		str = h.ExtendedFilePath
	}

	for i := uint16(0); i < h.LastColB-h.FristColB+1; i++ {
		res[i] = str
	}
	return res
}
