package jwt

import (
	"fmt"
	"github.com/golang-jwt/jwt/v4"
	"github.com/speps/go-hashids/v2"
	"time"
)

type JWT struct {
	SigningKey []byte          // 签名密钥
	ExpiresIn  time.Duration   // token 过期时间
	encoder    *hashids.HashID // 数据区加密器
	data       *Data           // 数据区
}

func New(key []byte, salt string) *JWT {

	return &JWT{
		SigningKey: key,
		ExpiresIn:  time.Hour * 24 * 3,
		encoder:    NewEncoder(salt),
	}
}

func (j *JWT) WithData(data *Data) *JWT {
	j.data = data
	return j
}

func (j *JWT) Encode() string {

	now := time.Now()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &LoginClaims{
		RegisteredClaims: jwt.RegisteredClaims{
			Subject:   "login",
			IssuedAt:  jwt.NewNumericDate(now),
			ExpiresAt: jwt.NewNumericDate(now.Add(j.ExpiresIn)),
		},
		Data: j.data,
	})

	s, _ := token.SignedString(j.SigningKey)
	return s
}

func (j *JWT) Decode(s string) error {

	token, err := jwt.ParseWithClaims(s, &LoginClaims{}, func(token *jwt.Token) (interface{}, error) {
		return j.SigningKey, nil
	})
	if err != nil {
		return err
	}

	if claims, ok := token.Claims.(*LoginClaims); ok && token.Valid {
		j.data = claims.Data
		return nil
	}

	return fmt.Errorf("token 无效")
}

func (j *JWT) Data() *Data {
	return j.data
}

func (j *JWT) GetEncoder() *hashids.HashID {
	return j.encoder
}
