package exante

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"log"
	"time"
	"unicode/utf8"
)

type JWT struct {
	ClientID      string   `json:"iss"`
	ApplicationID string   `json:"sub"`
	IssuedAt      int64    `json:"iat"`
	Expiration    int64    `json:"exp"`
	Rights        []string `json:"aud"`
}

func NewToken(cfg Config) string {
	token := JWT{
		ClientID:      cfg.ClientID,
		ApplicationID: cfg.AppID,
		Rights:        []string{"symbols", "ohlc", "feed", "change", "crossrates", "orders", "summary", "accounts", "transactions"},
		IssuedAt:      time.Now().Unix(),
		Expiration:    time.Now().AddDate(1, 0, 0).Unix(),
	}
	dat, err := json.Marshal(token)
	if err != nil {
		log.Fatalln(err)
	}
	payload := base64.RawStdEncoding.EncodeToString(toUTF8(string(dat)))
	var buffer bytes.Buffer
	buffer.WriteString("eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.")
	buffer.WriteString(payload)
	sha := hmac256(cfg.Key, buffer.Bytes())
	buffer.WriteRune('.')
	buffer.WriteString(sha)
	return string(buffer.Bytes())
}

func toUTF8(source string) []byte {
	bs := make([]byte, len(source)*utf8.UTFMax)
	count := 0
	for _, r := range source {
		count += utf8.EncodeRune(bs[count:], r)
	}
	bs = bs[:count]
	return bs
}

func hmac256(secret string, source []byte) string {
	h := hmac.New(sha256.New, []byte(secret))
	h.Write(source)
	return base64.RawStdEncoding.EncodeToString(h.Sum(nil))
}
