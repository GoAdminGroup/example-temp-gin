package security

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestReverseString(t *testing.T) {
	// mock
	tests := []struct {
		name   string
		source string
		want   string
	}{
		{
			name:   "URL: baidu.com",
			source: `http://www.baidu.com/adad/1232312/#1233?id=123`,
			want:   `321=di?3321#/2132321/dada/moc.udiab.www//:ptth`,
		},
	}
	// do
	for _, test := range tests {
		result := ReverseString(test.source)
		// verify
		t.Logf("source: %v, result: %v", test.source, result)
		assert.Equal(t, test.want, result)
	}
}
