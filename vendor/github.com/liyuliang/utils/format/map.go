package format

import "encoding/json"

type data map[string]string

func (mapData data) Values() (data []string) {
	for _, value := range mapData {
		data = append(data, value)
	}
	return data
}

func (mapData data) Keys() (data []string) {

	for key, _ := range mapData {
		data = append(data, key)
	}
	return data
}

func (mapData data) ToString() string {
	byteData, _ := json.Marshal(mapData)
	return string(byteData)
}

func MapData() data {
	return make(map[string]string)
}

func ToMapData(params map[string]string) data {
	mapData := MapData()
	for key, value := range params {
		mapData[key] = value
	}
	return mapData
}
