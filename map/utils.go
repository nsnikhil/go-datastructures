package gmap

import (
	"bytes"
	"encoding/binary"
	"encoding/json"
	"hash"
	"math"
)

func toBytes[T comparable](e T) ([]byte, error) {
	buf := new(bytes.Buffer)
	if err := json.NewEncoder(buf).Encode(e); err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

func hashCode[T comparable](h *hash.Hash, e T) (uint32, error) {
	b, err := toBytes(e)
	if err != nil {
		return math.MaxInt32, err
	}

	_, err = (*h).Write(b)
	if err != nil {
		return math.MaxInt32, err
	}

	res := (*h).Sum(nil)
	(*h).Reset()

	return binary.BigEndian.Uint32(res), nil
}

var m = make(map[interface{}]map[int64]int)

func indexOf[T comparable](h *hash.Hash, e T, capacity int64) (int64, error) {
	f, err := hashCode(h, e)
	if err != nil {
		return invalidIndex, err
	}

	return int64(math.Mod(float64(f>>16), float64(capacity))), nil
}
