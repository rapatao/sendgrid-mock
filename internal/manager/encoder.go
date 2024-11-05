package manager

import (
	"encoding/base64"
)

func encode(value string) string {
	return base64.StdEncoding.EncodeToString([]byte(value))
}

func decode(value string) (string, error) {
	decodeString, err := base64.StdEncoding.DecodeString(value)
	if err != nil {
		return "", err
	}

	return string(decodeString), nil
}
