package learnku_json

import (
	"encoding/json"
)

// @title json 转 map
// @params str 是一个 json 格式的 byte 切片
// @params _map 是一个 map  指针【必须传入指针】
// @return map
func Json2Map(str []byte, _map interface{}) (bool, error) {
	err := json.Unmarshal(str, _map)
	if err != nil {
		return false, err
	}
	return true, nil
}

// @title json 转 struct结构体
// @params str 是一个 json 格式的 byte 切片
// @params _struct 是一个 struct 指针【必须传入指针】
// @return bool 转换是否成功
func Json2Struct(str []byte, _struct interface{}) (bool, error) {
	err := json.Unmarshal(str, _struct)
	if err != nil {
		return false, err
	}
	return true, nil
}

// @title map 转 json
// @params data map
// @return 是一个 json 格式的 byte 切片
func Map2Json(m map[string]interface{}) ([]byte, error) {
	return json.Marshal(m)
}

// @title struct结构体 转 json
// @params data struct
// @return 是一个 json 格式的 byte 切片
func Struct2Json(s interface{}) ([]byte, error) {
	return json.Marshal(s)
}
