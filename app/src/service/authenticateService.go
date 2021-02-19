package service

import (
	"log"
	"net/http"
	"path/filepath"
	"strings"

	"github.com/magiconair/properties"
	jwtverifier "github.com/okta/okta-jwt-verifier-golang"
)

// LogHandlerFunc defines the function prototype for logging errors.
type logHandlerFunc func(fmt string, args ...interface{})

var logPrintf logHandlerFunc = log.Printf

/*
IsAuthenticated func to authenticae a request
*/
func IsAuthenticated(r *http.Request) bool {
	authHeader := r.Header.Get("Authorization")

	if authHeader == "" {
		return false
	}
	tokenParts := strings.Split(authHeader, "Bearer ")
	bearerToken := tokenParts[1]

	jwtVerifier := getJWTVerifier()

	_, err := jwtVerifier.VerifyAccessToken(bearerToken)

	if err != nil {
		return false
	}

	return true
}

func getJWTVerifier() *jwtverifier.JwtVerifier {
	path, err := filepath.Abs("app/src/config/config.properties")
	if err != nil {
		logPrintf("Error creating file path %s ", path)
		return &jwtverifier.JwtVerifier{}
	}
	p := properties.MustLoadFile(path, properties.UTF8)
	toValidate := map[string]string{}
	toValidate["aud"] = p.GetString("audience", "")
	toValidate["cid"] = p.GetString("clientId", "")

	verifier := jwtverifier.JwtVerifier{
		Issuer:           p.GetString("issuer", ""),
		ClaimsToValidate: toValidate,
	}
	return verifier.New()
}
