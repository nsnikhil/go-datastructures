package gmap

import (
	"bytes"
	"encoding/binary"
	"encoding/json"
	"github.com/nsnikhil/go-datastructures/utils"
	"hash"
	"math"
)

func toBytes(e interface{}) ([]byte, error) {
	buf := new(bytes.Buffer)
	if err := json.NewEncoder(buf).Encode(e); err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

func hashCode(h *hash.Hash, e interface{}) (uint32, error) {
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

func indexOf(h *hash.Hash, e interface{}, capacity float64) (int, error) {
	f, err := hashCode(h, e)
	if err != nil {
		return utils.InvalidIndex, err
	}

	return int(math.Mod(float64(f>>16), capacity)), nil
}
