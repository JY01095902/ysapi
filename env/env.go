package env

import "os"

func MustGetEnv(env string) string {
	val := os.Getenv(env)
	if val == "" {
		panic("Please set env " + env)
	}

	return val
}

func MustGetEnvInProduction(env string) string {
	if MustGetEnv("ENV") == "production" {
		return MustGetEnv(env)
	}

	return os.Getenv(env)
}
