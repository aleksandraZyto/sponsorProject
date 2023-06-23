package services

import "encoding/base64"

func Encode(toEncode string) string {
	return base64.StdEncoding.EncodeToString([]byte(toEncode))
}
