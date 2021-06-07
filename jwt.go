package jwt

import "time"

type Payload struct {
	username string
}

func CreateToken(payload *Payload) (string, error) {
	return "123", nil
}

func VerifyToken(token string) (*Payload, error) {
	return &Payload{username: "Hady"}, nil

}
