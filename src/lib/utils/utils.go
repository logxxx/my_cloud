package utils

import (
	"bytes"
	"crypto/sha256"
	"encoding/base64"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
)

func GetHashFromHeader(h http.Header) string {
	digest := h.Get("digest")
	if len(digest) < 9 {
		return ""
	}
	if digest[:8] != "SHA-256=" {
		return ""
	}
	return digest[8:]
}

func GetSizeFromHeader(h http.Header) int64 {
	size, _ := strconv.ParseInt(h.Get("content-length"), 0, 64)
	return size
}

func CalculateHash(r io.Reader) string {
	resp, e := ioutil.ReadAll(r)
	if e != nil {
		log.Println("CalculateHash ReadAll err:", e)
		return ""
	}
	log.Println("CalculateHash ReadAll resp:", string(resp))
	h := sha256.New()
	io.Copy(h, bytes.NewReader(resp))
	return base64.StdEncoding.EncodeToString(h.Sum(nil))
}