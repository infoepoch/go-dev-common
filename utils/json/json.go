package json

import "encoding/json"

// 转换 字符串 到 JSON
func JsonFromStr(str string, v interface{}) error {
	err := json.Unmarshal([]byte(str), v)
	return err
}

func JsonFromByte(byte []byte, v interface{}) error {
	err := json.Unmarshal(byte, v)
	return err
}
