package utils

import "encoding/json"

type CallbackData struct {
	Position    string `json:"Position"`
	MessageId   int64  `json:"MessageId"`
	NextCommand string `json:"NextCommand"`
}

func EncodeCallbackData(data CallbackData) string {
	jsonData, _ := json.Marshal(data)
	return string(jsonData)
}

func DecodeCallbackData(data string) (result CallbackData, err error) {
	err = json.Unmarshal([]byte(data), &result)
	if err != nil {
		return result, err
	}

	return result, nil
}
