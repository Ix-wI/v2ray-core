package protocol

import (
	"io"
)

type RequestEncoder interface {
	EncodeRequestHeader(*RequestHeader, io.Writer)
	EncodeRequestBody(io.Writer) io.Writer
}

type RequestDecoder interface {
	DecodeRequestHeader(io.Reader) *RequestHeader
	DecodeRequestBody(io.Reader) io.Reader
}

type ResponseEncoder interface {
	EncodeResponseHeader(*ResponseHeader, io.Writer)
	EncodeResponseBody(io.Writer) io.Writer
}

type ResponseDecoder interface {
	DecodeResponseHeader(io.Reader) *ResponseHeader
	DecodeResponseBody(io.Reader) io.Reader
}
