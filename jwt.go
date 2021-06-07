package jwt

import "time"

type Payload struct {
	username string
}

func CreateToken(username string, duration time.Duration) (string, error) {
   return  ("123",nil)
}

func VerifyToken(token string) (*Payload, error) {
	return ({ uusername:"Hady" },nil)

}
