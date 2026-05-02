package test

import (
	"testing"
)

// 文件 _test结尾, 方法Test开头
func TestDemo(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping test in short mode.")
	}
}

func TestAccountEn(t *testing.T) {
}
