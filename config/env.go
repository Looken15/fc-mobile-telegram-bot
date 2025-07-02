package config

const EnvVarName = "ENV"

const (
	EnvProd = Env("prod")
	EnvDev  = Env("dev")
)

type Env string
