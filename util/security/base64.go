package security

import "encoding/base64"

func Base64StdEncoding(input string) string {
	return base64.StdEncoding.EncodeToString([]byte(input))
}

func Base64StdDecoding(input string) (string, error) {
	bytes, err := base64.StdEncoding.DecodeString(input)
	if err != nil {
		return "", err
	}
	return string(bytes), nil
}

func Base64URLEncoding(input string) string {
	return base64.URLEncoding.EncodeToString([]byte(input))
}

func Base64URLDecoding(input string) (string, error) {
	bytes, err := base64.URLEncoding.DecodeString(input)
	if err != nil {
		return "", err
	}
	return string(bytes), nil
}

func Base64TripleURLEncoding(input string) string {
	src := []byte(input)
	enc := base64.URLEncoding

	bufOne := make([]byte, enc.EncodedLen(len(src)))
	enc.Encode(bufOne, src)

	bufTwo := make([]byte, enc.EncodedLen(len(bufOne)))
	enc.Encode(bufTwo, bufOne)

	bufThree := make([]byte, enc.EncodedLen(len(bufTwo)))
	enc.Encode(bufThree, bufTwo)
	return string(bufThree)
}

func Base64TripleURLDecoding(input string) (string, error) {
	src := []byte(input)
	enc := base64.URLEncoding

	bufOne := make([]byte, enc.DecodedLen(len(src)))
	n, err := base64.URLEncoding.Decode(bufOne, src)
	if err != nil {
		return "", err
	}
	nextBuf := bufOne[:n]

	bufTwo := make([]byte, enc.DecodedLen(len(nextBuf)))
	n, err = base64.URLEncoding.Decode(bufTwo, nextBuf)
	if err != nil {
		return "", err
	}
	nextBuf = bufTwo[:n]

	bufTree := make([]byte, enc.DecodedLen(len(nextBuf)))
	n, err = base64.URLEncoding.Decode(bufTree, nextBuf)
	if err != nil {
		return "", err
	}

	return string(bufTree[:n]), nil
}

func ReverseBase64URLDecode(input string) (string, error) {
	base64Code := reverseBase64Code(input)
	src := []byte(base64Code)
	enc := base64.URLEncoding

	buf := make([]byte, enc.DecodedLen(len(src)))
	n, err := base64.URLEncoding.Decode(buf, src)
	if err != nil {
		return "", err
	}
	return string(buf[:n]), nil
}

func reverseBase64Code(input string) string {
	midSource := ReverseString(input)
	code := len([]byte(midSource)) * 6
	if code > 24 {
		mod := code % 24
		if mod > 0 && mod < 18 {
			return midSource + "=="
		} else if mod > 17 {
			return midSource + "="
		} else {
			return midSource
		}
	} else if code == 24 {
		return midSource
	} else {
		if code > 0 && code < 18 {
			return midSource + "=="
		} else if code > 17 {
			return midSource + "="
		} else {
			return midSource
		}
	}
}
