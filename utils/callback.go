package utils

import "encoding/json"

type CallbackData struct {
	Position    string `json:"P"`
	Tactic      string `json:"T"`
	MessageId   int64  `json:"M"`
	NextCommand string `json:"NC"`
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
