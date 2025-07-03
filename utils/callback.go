package utils

import "encoding/json"

type CallbackData struct {
	Position  string `json:"Position"`
	MessageId int64  `json:"MessageId"`
}

func EncodeCallbackData(data CallbackData) (string, error) {
	jsonData, err := json.Marshal(data)
	if err != nil {
		return "", err
	}

	return string(jsonData), nil
}

func DecodeCallbackData(data string) (result CallbackData, err error) {
	err = json.Unmarshal([]byte(data), &result)
	if err != nil {
		return result, err
	}

	return result, nil
}
