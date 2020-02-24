package security

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestURLPathEncode(t *testing.T) {
	// mock
	tests := []struct {
		name   string
		source string
		want   string
	}{
		{
			name:   "URLPathEncode: 你好",
			source: `你好 我在`,
			want:   `%E4%BD%A0%E5%A5%BD%20%E6%88%91%E5%9C%A8`,
		},
	}
	// do
	for _, test := range tests {
		result := URLPathEncode(test.source)
		// verify
		t.Logf("source: %v, result: %v", test.source, result)
		assert.Equal(t, test.want, result)
	}
}

func TestURLPathDecode(t *testing.T) {
	// mock
	tests := []struct {
		name   string
		source string
		want   string
	}{
		{
			name:   "URLPathDecode: 你好",
			source: `%E4%BD%A0%E5%A5%BD%20%E6%88%91%E5%9C%A8`,
			want:   `你好 我在`,
		},
	}
	// do
	for _, test := range tests {
		result, err := URLPathDecode(test.source)
		if err != nil {
			t.Fatalf("TestURLPathDecode error: %v", err)
		}
		// verify
		t.Logf("source: %v, result: %v", test.source, result)
		assert.Equal(t, test.want, result)
	}
}

func TestURLQueryEncode(t *testing.T) {
	// mock
	tests := []struct {
		name   string
		source string
		want   string
	}{
		{
			name:   "URLQueryEncode: 你好",
			source: `你好 我在`,
			want:   `%E4%BD%A0%E5%A5%BD+%E6%88%91%E5%9C%A8`,
		},
	}
	// do
	for _, test := range tests {
		result := URLQueryEncode(test.source)
		// verify
		t.Logf("source: %v, result: %v", test.source, result)
		assert.Equal(t, test.want, result)
	}
}

func TestURLQueryDecode(t *testing.T) {
	// mock
	tests := []struct {
		name   string
		source string
		want   string
	}{
		{
			name:   "URLQueryDecode: 你好",
			source: `%E4%BD%A0%E5%A5%BD%20%E6%88%91%E5%9C%A8`,
			want:   `你好 我在`,
		},
	}
	// do
	for _, test := range tests {
		result, err := URLQueryDecode(test.source)
		if err != nil {
			t.Fatalf("TestURLPathDecode error: %v", err)
		}
		// verify
		t.Logf("source: %v, result: %v", test.source, result)
		assert.Equal(t, test.want, result)

	}
}

func TestUnicodeDecode(t *testing.T) {
	// mock
	tests := []struct {
		name   string
		source string
		want   string
	}{
		{
			name:   "UnicodeDecode: 你好",
			source: `\u4f60\u597d\u0020\u6211\u5728`,
			want:   `你好 我在`,
		},
	}
	// do
	for _, test := range tests {
		result, err := UnicodeDecode(test.source)
		if err != nil {
			t.Fatalf("TestUnicodeDecode error: %v", err)
		}
		// verify
		t.Logf("source: %v, result: %v", test.source, result)
		assert.Equal(t, test.want, result)
	}
}
