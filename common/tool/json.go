package tool

import (
	jsoniter "github.com/json-iterator/go"
	"strings"
)

var Cjson = jsoniter.ConfigCompatibleWithStandardLibrary

func JsonEncode(v interface{}) ([]byte, error) {
	return Cjson.Marshal(v)
}

func JsonDecodeV2(v string) (map[string]string, error) {
	reader := strings.NewReader(v)
	decoder := Cjson.NewDecoder(reader)
	params := make(map[string]string)
	err := decoder.Decode(&params)
	return params, err
}

func JsonDecode(v string) (map[string]interface{}, error) {
	reader := strings.NewReader(v)
	decoder := Cjson.NewDecoder(reader)
	params := make(map[string]interface{})
	err := decoder.Decode(&params)
	return params, err
}

func JsonMarshal(v interface{}) ([]byte, error) {
	return Cjson.Marshal(v)
}

func JsonMarshalString(v interface{}) (string, error) {
	return Cjson.MarshalToString(v)
}

func JsonUnmarshal(data []byte, v interface{}) error {
	return Cjson.Unmarshal(data, v)
}

func JsonUnmarshalFromString(data string, v interface{}) error {
	return Cjson.UnmarshalFromString(data, v)
}

func MarshalToString(v interface{}) (string, error) {
	return Cjson.MarshalToString(v)
}

func GetValueByKeyFromJson(data []byte, key string) string {
	return Cjson.Get(data, key).ToString()
}
