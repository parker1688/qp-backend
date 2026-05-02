package tool

import "testing"

func TestName(t *testing.T) {
	t.Log(HashEncodeInt64([]int64{Int(37787029961965701)}))
	t.Log(HashDecodeInt64("1QLb99xv52R9"))
}

func TestValidCPF(t *testing.T) {
	t.Log(ValidCPF("48979218133"))
}

func TestHMACSignatureSha256(t *testing.T) {
	t.Log(HMACSignatureSha256("sha256", "111"))
}
