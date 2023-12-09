package encoding

import "github.com/ireina7/fgo/structs/result"

type Encoder[T any] interface {
	Encode(T) []byte
}

type Encoding interface {
	Encode() []byte
}

type Decoder[T any] interface {
	Decode([]byte) result.Result[T]
}
