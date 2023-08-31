package env

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMustGetEnvInProduction(t *testing.T) {
	os.Setenv("ENV", "production")
	key := "LOGISTICS_HUB_DB_USERNAME"
	os.Setenv(key, "1.1.1")
	f := assert.PanicTestFunc(func() { _ = MustGetEnvInProduction(key) })
	assert.NotPanics(t, f, "production 变环境 env有值，应该不报错")

	os.Setenv(key, "")
	f = assert.PanicTestFunc(func() { _ = MustGetEnvInProduction(key) })
	assert.Panics(t, f, "production 变环境 env没有值，应该panic")

	os.Setenv("ENV", "staging")
	os.Setenv(key, "1.1.1")
	f = assert.PanicTestFunc(func() { _ = MustGetEnvInProduction(key) })
	assert.NotPanics(t, f, "非 production 变环境 env有值，应该不报错")

	os.Setenv(key, "")
	f = assert.PanicTestFunc(func() { _ = MustGetEnvInProduction(key) })
	assert.NotPanics(t, f, "非 production 变环境 env没有值，应该不报错")
}
