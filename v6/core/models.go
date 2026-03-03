package core

import "io"

type Response interface {
	IsRef() bool
}

type BinaryMetadata struct {
	ContentType        string
	ContentDisposition string
}

type BinaryResponse interface {
	SetData(dataSource io.Reader, metadata BinaryMetadata) error
}

type Request interface{}
