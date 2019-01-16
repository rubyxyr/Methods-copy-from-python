package main

import (
	"encoding/binary"
	"strings"
	"crypto/rand"
)

func divmod(numerator, denominator int64) (quotient, remainder int64) {
	quotient = numerator / denominator // integer division, decimals are truncated
	remainder = numerator % denominator
	return
}

func CreateSuid() string{
	random_bs := make([]byte, 8)
	rand.Read(random_bs)

	num := binary.LittleEndian.Uint64(random_bs) & (1<<63 - 1)
	alphabet_string := "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"
	alphabet := strings.Split(alphabet_string, "")
	var key[]string
	num_left := int64(num)
	for num_left > 0{
		var rem int64
		num_left, rem = divmod(num_left, 62)
		key = append(key, alphabet[rem])
	}

	return strings.Join(key, "")
}