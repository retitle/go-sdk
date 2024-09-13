package tests_utils

import (
	"encoding/json"
	"io"
	"strings"
)

func ParseStructToIoReadCloser[T any](v *T) io.ReadCloser {
	b, _ := json.Marshal(v)
	stringReader := strings.NewReader(string(b))
	stringReadCloser := io.NopCloser(stringReader)
	return stringReadCloser
}

func ParseReaderToStruct[T any](v io.Reader, t *T) error {
	bytes, err := io.ReadAll(v)
	if err != nil {
		return err
	}
	err = json.Unmarshal(bytes, t)
	return err
}
