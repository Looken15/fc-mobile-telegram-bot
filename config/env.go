package config

const EnvVarName = "APP_ENV"

const (
	EnvProd = Env("prod")
	EnvDev  = Env("dev")
)

type Env string
