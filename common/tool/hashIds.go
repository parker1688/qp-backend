package tool

import (
	"github.com/speps/go-hashids/v2"
)

func HashEncodeInt64(data []int64) string {
	hd := hashids.NewData()
	hd.Salt = "salt"
	hd.MinLength = 3
	h, _ := hashids.NewWithData(hd)
	e, _ := h.EncodeInt64(data)
	return e
}

func HashDecodeInt64(s string) []int64 {
	hd := hashids.NewData()
	hd.Salt = "salt"
	hd.MinLength = 3
	h, _ := hashids.NewWithData(hd)
	d, err := h.DecodeInt64WithError(s)
	if err != nil {
		return []int64{}
	}
	return d
}
