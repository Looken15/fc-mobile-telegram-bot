package config

import (
	"encoding/json"
	"fmt"
	"os"
)

func Get() *Settings {
	env, err := GetEnv()
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	fmt.Println(env)

	var result Settings
	filePath := "./.config/" + string(env) + ".json"
	fileBytes, err := os.ReadFile(filePath)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	if err := json.Unmarshal(fileBytes, &result); err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	result.Token = os.Getenv("TOKEN")
	result.DbUrl = os.Getenv("DBCONN")

	return &result
}

func GetEnv() (Env, error) {
	env := os.Getenv(EnvVarName)

	if env == "" {
		return EnvDev, nil
	}
	switch Env(env) {
	case EnvDev:
		return EnvDev, nil
	case EnvProd:
		return EnvProd, nil
	default:
		return "", fmt.Errorf("unknown environment variable %s", os.Args[1])
	}
}
