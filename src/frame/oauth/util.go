package oauth

import (
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"fmt"
)

const Base64HS256Header = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9"
const JWTSeparators = "."

func JwtHS256(data interface{}, secret string) string {
	dataBytes, _ := json.Marshal(data)

	payload := base64.StdEncoding.EncodeToString(dataBytes)
	headerAndPayload := fmt.Sprint(Base64HS256Header, JWTSeparators, payload)

	signatureString := fmt.Sprint(headerAndPayload, secret)
	sha256Signature := sha256.New()
	sha256Signature.Write([]byte(signatureString))

	signatureBytes := sha256Signature.Sum(nil)
	signature := base64.StdEncoding.EncodeToString(signatureBytes)
	return fmt.Sprint(headerAndPayload, JWTSeparators, signature)
}
