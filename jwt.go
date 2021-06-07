package jwt

// import (
//     "crypto/hmac"
//     "crypto/sha256"
// )

type Payload struct {
	username string
}

type header struct {
}

func newHeaderEncoded() string {

	return "header"
}

func newPayloadEncoded() string {
	return "payload"
}

func newSignature(h string, p string, s string) string {
	return "sign"
}

func CreateToken(payload *Payload) (string, error) {
	encoded_header := newHeaderEncoded()
	encoded_payload := newPayloadEncoded()
	secret := "secret"
	signature := newSignature(encoded_header, encoded_payload, secret)
	token := "." + signature
	return token, nil
}

func VerifyToken(token string) (*Payload, error) {
	return &Payload{username: "Hady"}, nil

}
