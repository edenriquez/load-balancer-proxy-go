package utils

import (
	"path"
	"path/filepath"
	"strings"

	"github.com/joho/godotenv"
)

// LoadEnvVars find and load environment variables
func LoadEnvVars() {
	base, _ := filepath.Abs("../")
	envPath := path.Join(base, ".env")
	godotenv.Load(envPath)
}

// SanitizeSQLParam returns a sanitized param
func SanitizeSQLParam(param interface{}) interface{} {
	sanitizedParam := param
	replaceValues := map[string]string{
		"'":  "''",
		"\n": "",
		"\\": "\\\\",
		"\"": "\"\"",
		"%":  "_%",
		"*":  "_*",
		"=":  "_=",
		"?":  "_?",
		";":  "_;",
		"$":  "_$",
		"@":  "_@",
	}
	if str, ok := param.(string); ok {
		for k, v := range replaceValues {
			str = strings.Replace(str, k, v, -1)
		}
		sanitizedParam = str
	}
	return sanitizedParam
}
