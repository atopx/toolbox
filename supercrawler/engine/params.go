package engine

import (
	"encoding/json"
	"supercrawler/common/utils"
)

func EncodeParams(label string, data []string) map[string]interface{} {
	b, _ := json.Marshal(data)
	return map[string]interface{}{"label": label, "data": utils.String(b)}
}

func DecodeParams(value map[string]interface{}) (label string, data []string) {
	label = value["label"].(string)
	_ = json.Unmarshal(utils.Bytes(value["data"].(string)), &data)
	return label, data
}
