package internal

import (
	"encoding/base64"
	"fmt"
)

// Encoder is a struct able to encode / decode base64 strings
type Encoder struct {
	WithPadding      bool
	URLEncodingFomat bool
	ShouldDecode     bool
}

// GetEncoder returns the base64 encoder to use based on the WithPadding and URLEncodingFomat values
func (e Encoder) GetEncoder() *base64.Encoding {
	if e.URLEncodingFomat {
		if e.WithPadding {
			return base64.URLEncoding
		}
		return base64.RawURLEncoding
	}

	if e.WithPadding {
		return base64.StdEncoding
	}

	return base64.RawStdEncoding
}

// Transform encode/decode a string to its string/base64 representation based on the ShouldDecode value
func (e Encoder) Transform(input string) (string, error) {
	if e.ShouldDecode {
		data, err := e.GetEncoder().DecodeString(input)
		if err != nil {
			return "", fmt.Errorf("can't decode input: %v", err)
		}

		return string(data), nil
	}

	return e.GetEncoder().EncodeToString([]byte(input)), nil
}
