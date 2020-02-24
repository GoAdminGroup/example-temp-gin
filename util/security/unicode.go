package security

import (
	"bytes"
	"encoding/binary"
	"encoding/hex"
	"net/url"
	"strings"
)

func URLPathEncode(input string) string {
	return url.PathEscape(input)
}

func URLPathDecode(input string) (string, error) {
	return url.PathUnescape(input)
}

func URLQueryEncode(input string) string {
	return url.QueryEscape(input)
}

func URLQueryDecode(input string) (string, error) {
	return url.QueryUnescape(input)
}

func UnicodeDecode(input string) (out string, err error) {
	bs, err := hex.DecodeString(strings.Replace(input, `\u`, ``, -1))
	if err != nil {
		return "", err
	}
	for i, bl, br, r := 0, len(bs), bytes.NewReader(bs), uint16(0); i < bl; i += 2 {
		err := binary.Read(br, binary.BigEndian, &r)
		if err != nil {
			return "", err
		}
		out += string(r)
	}
	return out, nil
}
