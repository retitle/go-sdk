package core

type Response interface {
	IsRef() bool
}

type Request interface{}
