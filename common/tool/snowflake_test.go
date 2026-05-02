package tool

import (
	"testing"
)

func TestSnowflakeId(t *testing.T) {
	t.Log(SnowflakeId())
}

func TestSnowflakeIdByKey(t *testing.T) {
	type args struct {
		key string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SnowflakeIdByKey(tt.args.key); got != tt.want {
				t.Errorf("SnowflakeIdByKey() = %v, want %v", got, tt.want)
			}
		})
	}
}
