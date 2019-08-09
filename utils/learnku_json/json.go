package learnku_json

import (
	"encoding/json"
	"github.com/astaxie/beego"
)

// json 转 map
func Json2Map(str []byte) map[string]interface{} {
	var tmp map[string]interface{}
	err := json.Unmarshal(str, &tmp)
	if err != nil {
		beego.Error("Json2Map 失败：" + string(str))
	}
	return tmp
}