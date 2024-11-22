package exchangesignature

import (
	"crypto/hmac"
	"crypto/sha512"
	"fmt"
	"math/rand"
	"net/url"
	"sort"
	"strings"
	"time"
)

func NewOrbix(secret, payload string) (string, error) {
	var sig string
	values, err := url.ParseQuery(payload)
	if err != nil {
		return sig, err
	}
	var keys []string
	for key := range values {
		keys = append(keys, key)
	}
	sort.Strings(keys)
	var sortedPayload []string
	for _, key := range keys {
		sortedPayload = append(sortedPayload, fmt.Sprintf("%s=%s", key, values.Get(key)))
	}
	payload = strings.Join(sortedPayload, "&")
	h := hmac.New(sha512.New, []byte(secret))
	h.Write([]byte(payload))
	sig = fmt.Sprintf("%x", h.Sum(nil))
	return sig, nil
}

func GenerateNonceInt() int64 {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	return r.Int63n(90000) + 10000
}
