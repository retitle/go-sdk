package tests_utils

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"strings"
)

func ParseStructToIoReadCloser[T any](v *T) io.ReadCloser {
	b, _ := json.Marshal(v)
	stringReader := strings.NewReader(string(b))
	stringReadCloser := io.NopCloser(stringReader)
	return stringReadCloser
}

func ParseDirectStructToIoReadClose[T any](v T) io.ReadCloser {
	b, _ := json.Marshal(v)
	stringReader := strings.NewReader(string(b))
	stringReadCloser := io.NopCloser(stringReader)
	return stringReadCloser
}

func ParseReaderToStruct[T any](v io.Reader, t *T) error {
	bytes, err := ioutil.ReadAll(v)
	if err != nil {
		return err
	}
	err = json.Unmarshal(bytes, t)
	return err
}
