package StructToJson

import (
	"bootpkg/common/global"
	"bootpkg/pkg/core/modules/dos"
	"fmt"
	"testing"
)

func TestToJSON(t *testing.T) {
	j := &dos.DictsDetail{}
	fmt.Println(global.JSON.MarshalToString(j))
}