package jwt

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/dgrijalva/jwt-go"
)

const CERTS_URL = "https://www.googleapis.com/oauth2/v1/certs"

type CustomClaims struct {
	Email      string `json: email`
	PictureUrl string `json: picture`
	Name       string `json: name`
	jwt.StandardClaims
}

func ParseIdToken(idToken string) (*CustomClaims, error) {
	token, err := jwt.ParseWithClaims(idToken, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		kid := token.Header["kid"]
		return jwt.ParseRSAPublicKeyFromPEM([]byte(getJson(CERTS_URL)[kid.(string)].(string)))
	})

	if err != nil {
		return nil, err
	}

	return token.Claims.(*CustomClaims), nil
}

func getJson(url string) map[string]interface{} {
	res, _ := http.Get(url)
	body, _ := ioutil.ReadAll(res.Body)
	var mapData map[string]interface{}
	json.Unmarshal(body, &mapData)
	return mapData
}
