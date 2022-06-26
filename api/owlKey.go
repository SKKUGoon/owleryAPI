package api

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"net/http"
	"net/url"
	"time"
)

func getApiKey(p url.Values) bool {
	key := p.Get("apiKey")
	return key == ""
}

func NewRegister(w http.ResponseWriter, r *http.Request) {
	params := r.URL.Query()
	procNewRegister(params)
}

func procNewRegister(p url.Values) {
	user := p.Get("username")
	t := time.Now().Format("2006-01-02T15:04:05")
	key := sha256.New()
	key.Write([]byte(user + t))

	hashKey := hex.EncodeToString(key.Sum(nil))
	fmt.Println(hashKey)
}
