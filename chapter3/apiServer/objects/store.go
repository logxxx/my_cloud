package objects

import (
	"fmt"
	"io"
	"lib/utils"
	"net/http"
	"net/url"
	"../locate"
)

func storeObject(r io.Reader, hash string, size int64) (int, error) {
	if locate.Exist(url.PathEscape(hash)) {
		return http.StatusOK, nil
	}

	stream, e := putStream(url.PathEscape(hash), size)
	if e != nil {
		return http.StatusServiceUnavailable, e
	}

	reader := io.TeeReader(r, stream)
	d := utils.CalculateHash(reader)
	if d != hash {
		stream.Commit(false)
		return http.StatusBadRequest, fmt.Errorf("object hash mismatch. cal=%s req=%s", d, hash)
	}
	stream.Commit(true)
	return http.StatusOK, nil
}