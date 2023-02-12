package utility

import "strings"

func GetTokenStringFromBearerToken(bearerToken string) string {
	splitBearerToken := strings.Split(bearerToken, " ")
	if len(splitBearerToken) != 2 {
		return ""
	}

	if splitBearerToken[0] != "Bearer" {
		return ""
	}

	return splitBearerToken[1]
}
