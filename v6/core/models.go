package core

import "io"

type Response interface {
	IsRef() bool
}

type BinaryMetaData struct {
	ContentType string
}

type BinaryResponse interface {
	SetData(dataSource io.Reader, metaData BinaryMetaData) error
}

type Request interface{}
