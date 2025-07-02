package utils

import (
	"encoding/json"
	"fmt"
)

func PrintJSON(data interface{}) error {
	jsonData, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		return fmt.Errorf("ошибка преобразования в JSON: %w", err)
	}

	fmt.Println(string(jsonData))
	return nil
}
