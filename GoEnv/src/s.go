package main

import "io"

type Reader interface {
	Read(p []byte)(n int,err error)
}
type ReaderAt interface {
	ReaderAt(p []byte,off int64)(n int,err error)
}