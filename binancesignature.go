package exchangesignature

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
)

func NewBinance(secret, payload string) (string, error) {
	h := hmac.New(sha256.New, []byte(secret))
	h.Write([]byte(payload))
	signature := h.Sum(nil)
	return hex.EncodeToString(signature), nil
}
