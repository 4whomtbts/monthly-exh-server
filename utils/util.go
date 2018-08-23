package utils

import (
	"encoding/json"
	"io"
)

func CopyStringMap(originalMap map[string]string) map[string]string {
	copyMap := make(map[string]string)
	for k, v := range originalMap {
		copyMap[k] = v
	}
	return copyMap
}

func MapToJson(objmap map[string]string) string {
	b, _ := json.Marshal(objmap)
	return string(b)
}

func MapFromJson(data io.Reader) map[string]string {
	decoder := json.NewDecoder(data)
	var objmap map[string]string
	if err := decoder.Decode(&objmap); err != nil {
		return make(map[string]string)
	}else{
		return objmap
	}
}

func ArrayToJson(data io.Reader) []string {
	decoder := json.NewDecoder(data)
	var objmap []string
	if err := decoder.Decode(&objmap); err != nil {
		return make([]string, 0)
	}else{
		return objmap
	}
}

func ArrayFromInterface(data interface{}) []string {
	stringArray := []string{}
	dataArray, ok := data.([]interface{})
	if !ok {
		return stringArray
	}

	for _, v := range dataArray {
		if str, ok := v.(string); ok {
			stringArray = append(stringArray,str)
		}
	}
	return stringArray
}

func StringInterfaceToJson(objmap map[string]interface{}) string {
	b, _ := json.Marshal(objmap)
	return string(b)
}

func StringInterfaceFromJson(data io.Reader) map[string]interface{} {
	decoder := json.NewDecoder(data)
	var objmap map[string]interface{}
	if err := decoder.Decode(&objmap); err != nil {
		return make(map[string]interface{})
	}else{
		return objmap
	}
}

func StringToJson(s string) string {
	b , _ := json.Marshal(s)
	return string(b)
}


func StringFromJson(data io.Reader) string {
	decoder := json.NewDecoder(data)

	var s string
	if err := decoder.Decode(&s); err != nil {
		return ""
	} else {
		return s
	}
}
