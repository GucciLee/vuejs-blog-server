package learnku_json

import (
	"log"
	"testing"
)

type Monster struct {
	Name     string `json:"monster_name"`
	Age      int    `json:"monster_age"`
	Birthday string
	Sal      float64
	Skill    string
}

// JSON 转 Map
func TestJson2Map(t *testing.T) {
	str := `{"address":"洪崖洞","age":"30","name":"红孩儿"}`
	data := map[string]string{}
	b, err := Json2Map([]byte(str), &data)
	if !b {
		log.Fatalf("1_1、TestJson2Map 失败；= %v", err)
	}
	t.Logf("1_1、TestJson2Map == %v", data)
	t.Logf("1_1、data[\"address\"] == %v", data["address"])
}

// JSON 转 Struct
func TestJson2Struct(t *testing.T) {
	str := `{"monster_name":"牛魔王22","monster_age":500,"Birthday":"2019-07-11","Sal":8000,"Skill":"牛魔拳11"}`
	// 定义一个 Monster 实例
	var monster Monster
	b, err := Json2Struct([]byte(str), &monster)
	if !b {
		log.Fatalf("1_2、TestJson2Struct 失败；= %v", err)
	}
	t.Logf("1_2、TestJson2Struct == %v", monster)
	t.Logf("1_2、`monster.Name = %v`", monster.Name)
}

// Map 转 JSON
func TestMap2Json(t *testing.T) {
	var a map[string]interface{}
	a = make(map[string]interface{})
	a["name"] = "红孩儿"
	a["age"] = 30
	a["address"] = "洪崖洞"

	data, err := Map2Json(a)
	if err == nil {
		t.Logf("2_1、TestMap2Json == %v", string(data))
	} else {
		log.Fatalln("2_1、TestMap2Json 失败")
	}
}

// Struct 转 JSON
func TestStruct2Json(t *testing.T) {
	monster := Monster{
		Name:     "牛魔王",
		Age:      500,
		Birthday: "2019-07-11",
		Sal:      8000.0,
		Skill:    "牛魔拳",
	}
	data, err := Struct2Json(monster)
	if err == nil {
		t.Logf("2_2、TestStruct2Json == %v", string(data))
	} else {
		log.Fatalln("2_2、TestStruct2Json 失败")
	}
}
