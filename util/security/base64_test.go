package security

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestBase64TripleURLEncoding(t *testing.T) {
	// mock
	tests := []struct {
		name   string
		source string
		want   string
	}{
		{
			name:   "TestBase64TripleURLEncoding URL: baidu.com",
			source: `http://www.baidu.com/adad/1232312/#1233?id=123`,
			want:   "WVVoU01HTkViM1pNTTJRelpIazFhVmxYYkd0a1V6VnFZakl3ZGxsWFVtaGFRemg0VFdwTmVVMTZSWGxNZVUxNFRXcE5lbEF5Ykd0UVZFVjVUWGM5UFE9PQ==",
		},
	}
	// do
	for _, test := range tests {
		encoding := Base64TripleURLEncoding(test.source)
		//t.Logf("Base64TripleURLEncoding\nFrom: %v\nTo %v", test.source, encoding)
		// verify
		assert.Equal(t, test.want, encoding)
	}
}

func TestBase64TripleURLDecoding(t *testing.T) {
	// mock
	tests := []struct {
		name   string
		source string
		want   string
	}{
		{
			name:   "TestBase64TripleURLDecoding URL: baidu.com",
			source: `WVVoU01HTkViM1pNTTJRelpIazFhVmxYYkd0a1V6VnFZakl3ZGxsWFVtaGFRemg0VFdwTmVVMTZSWGxNZVUxNFRXcE5lbEF5Ykd0UVZFVjVUWGM5UFE9PQ==`,
			want:   `http://www.baidu.com/adad/1232312/#1233?id=123`,
		},
	}
	// do
	for _, test := range tests {
		decoding, err := Base64TripleURLDecoding(test.source)
		if err != nil {
			t.Fatalf("TestBase64TripleURLDecoding err at %v, info: %v", test.name, err)
		}
		//t.Logf("Base64TripleURLEncoding\nFrom: %v\nTo %v", test.source, encoding)
		// verify
		assert.Equal(t, test.want, decoding)
	}
}

func TestBase64URLEncoding(t *testing.T) {
	// mock
	url := `http://www.baidu.com/adad/1232312/#1233?id=123`
	want := `aHR0cDovL3d3dy5iYWlkdS5jb20vYWRhZC8xMjMyMzEyLyMxMjMzP2lkPTEyMw==`
	// do
	encoding := Base64URLEncoding(url)
	// verify
	//t.Logf("encoding from %v to %v", url, encoding)
	assert.Equal(t, want, encoding)
}

func TestBase64URLDecoding(t *testing.T) {
	// mock
	source := `aHR0cDovL3d3dy5iYWlkdS5jb20vYWRhZC8xMjMyMzEyLyMxMjMzP2lkPTEyMw==`
	want := `http://www.baidu.com/adad/1232312/#1233?id=123`
	// do
	out, err := Base64URLDecoding(source)
	if err != nil {
		t.Fatalf("TestBase64URLDecoding error %v", err)
	}
	// verify
	assert.Equal(t, want, out)
}

func TestBase64StdEncoding(t *testing.T) {
	// mock
	source := `ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789+/=`
	want := `QUJDREVGR0hJSktMTU5PUFFSU1RVVldYWVphYmNkZWZnaGlqa2xtbm9wcXJzdHV2d3h5ejAxMjM0NTY3ODkrLz0=`
	// do
	out := Base64StdEncoding(source)
	// verify
	assert.Equal(t, want, out)
}

func TestBase64StdDecoding(t *testing.T) {
	// mock
	source := `QUJDREVGR0hJSktMTU5PUFFSU1RVVldYWVphYmNkZWZnaGlqa2xtbm9wcXJzdHV2d3h5ejAxMjM0NTY3ODkrLz0=`
	want := `ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789+/=`
	// do
	out, err := Base64StdDecoding(source)
	if err != nil {
		t.Fatalf("TestBase64StdDecoding err: %v", err)
	}
	// verify
	assert.Equal(t, want, out)
}

func TestReverseBase64Code(t *testing.T) {
	// mock
	tests := []struct {
		name   string
		source string
		want   string
	}{
		{
			name:   "TestReverseBase64Code key: d, value: 643",
			source: `zQjN`,
			want:   `NjQz`,
		},
		{
			name:   "TestReverseBase64Code key: az, value: 643",
			source: `2MjL3MTNGJTJpJXYmF2UwITJwITMuUjN4MjLw4yN3YkMlUWbvJHaDBjMlkybrNWZHBjMlU2apxGMyUyQyUCTNRFSLhCMyUiNz4yNzUjRyUCdptkYldVZsBHcBBjMlkiNfRTMfBTMwITJYBjMlM1TwITJjFWTwITJsVGdulEMyUiQzUCaz9Gdul2Yh1EKwITJw4SN`,
			want:   `NS4wJTIwKE1hY2ludG9zaCUzQiUyMEludGVsJTIwTWFjJTIwT1MlMjBYJTIwMTBfMTRfNiklMjBBcHBsZVdlYktpdCUyRjUzNy4zNiUyMChLSFRNTCUyQyUyMGxpa2UlMjBHZWNrbyklMjBDaHJvbWUlMkY3Ny4wLjM4NjUuMTIwJTIwU2FmYXJpJTJGNTM3LjM2`,
		},
		{
			name:   "TestReverseBase64Code URL: 0",
			source: `AM`,
			want:   `MA==`,
		},
		{
			name:   "TestReverseBase64Code URL: 0",
			source: `=AM`,
			want:   `MA==`,
		},
		{
			name:   "TestReverseBase64Code URL: A",
			source: `QQ`,
			want:   `QQ==`,
		},
		{
			name:   "TestReverseBase64Code URL: BC",
			source: `MKQ`,
			want:   `QKM=`,
		},
		{
			name:   "TestReverseBase64Code URL: landscape-primary",
			source: `knch1WayBXLlBXYjNHZuFGb`,
			want:   `bGFuZHNjYXBlLXByaW1hcnk=`,
		},
		{
			name:   "TestReverseBase64Code ",
			source: `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`,
			want:   `Q3VzdG9tJTIwdmFyJTIwJTNBX19SRUFDVF9ERVZUT09MU19HTE9CQUxfSE9PS19fJTJGJTJGJTIwQ3VzdG9tJTIwdmFyJTIwJTNBX19WVUVfREVWVE9PTFNfR0xPQkFMX0hPT0tfXyUyRiUyRiUyMEN1c3RvbSUyMGZ1bmN0aW9uJTIwJTNBZnVuY3Rpb24oZSUyQ3QpJTdCcmV0dXJuJTIwbmV3JTIway5mbi5pbml0KGUlMkN0KSU3RCUyRiUyRiUyMEN1c3RvbSUyMGZ1bmN0aW9uJTIwJTNBZnVuY3Rpb24oZSUyQ3QpJTdCcmV0dXJuJTIwbmV3JTIway5mbi5pbml0KGUlMkN0KSU3RCUyRiUyRiUyMEN1c3RvbSUyMHZhciUyMCUzQV9fUkVBQ1RfREVWVE9PTFNfQ09NUE9ORU5UX0ZJTFRFUlNfXyUyRiUyRiUyMEN1c3RvbSUyMHZhciUyMCUzQV9fUkVBQ1RfREVWVE9PTFNfQVBQRU5EX0NPTVBPTkVOVF9TVEFDS19fJTJGJTJGJTIwQ3VzdG9tJTIwZnVuY3Rpb24lMjAlM0FmdW5jdGlvbiUyMGFwcGVuZERhdGEoKSUyMCU3QiUwQSUwOSUwOSUwOSUyNCgnJTIzbm93UmVmZXJlcicpJTVCMCU1RC52YWx1ZSUyMCUzRCUyMHdpbmRvdy5sb2NhdGlvbi5ocmVmJTBBJTA5JTA5JTdEJTJGJTJGJTIwQ3VzdG9tJTIwdmFyJTIwJTNBXzB4MjFhYyUyRiUyRiUyMEN1c3RvbSUyMGZ1bmN0aW9uJTIwJTNBZnVuY3Rpb24lMjAoXzB4M2IzOTE1JTJDJTIwXzB4NWM5Yjc3KSUyMCU3QiUyMF8weDNiMzkxNSUyMCUzRCUyMF8weDNiMzkxNSUyMC0lMjAweDAlM0IlMjB2YXIlMjBfMHgyNzI1NjElMjAlM0QlMjBfMHgyMWFjJTVCXzB4M2IzOTE1JTVEJTNCJTIwaWYlMjAoXzB4NGRkNyU1QiduaFFVT28nJTVEJTIwJTNEJTNEJTNEJTIwdW5kZWZpbmVkKSUyMCU3QiUyMChmdW5jdGlvbiUyMCgpJTIwJTdCJTIwdmFyJTIwXzB4NWRhNjI3JTIwJTNEJTIwZnVuY3Rpb24lMjAoKSUyMCU3QiUyMHZhciUyMF8weDQyZDJiYiUzQiUyMHRyeSUyMCU3QiUyMF8weDQyZDJiYiUyMCUzRCUyMEZ1bmN0aW9uKCdyZXR1cm4lNUN4MjAoZnVuY3Rpb24oKSU1Q3gyMCclMjAlMkIlMjAnJTdCJTdELmNvbnN0cnVjdG9yKCU1Q3gyMnJldHVybiU1Q3gyMHRoaXMlNUN4MjIpKCU1Q3gyMCknJTIwJTJCJTIwJyklM0InKSgpJTNCJTIwJTdEJTIwY2F0Y2glMjAoXzB4NTUxMWRkKSUyMCU3QiUyMF8weDQyZDJiYiUyMCUzRCUyMHdpbmRvdyUzQiUyMCU3RCUyMHJldHVybiUyMF8weDQyZDJiYiUzQiUyMCU3RCUzQiUyMHZhciUyMF8weDM1ZmUwZiUyMCUzRCUyMF8weDVkYTYyNygpJTNCJTIwdmFyJTIwXzB4MWRkMjEwJTIwJTNEJTIwJ0FCQ0RFRkdISUpLTE1OT1BRUlNUVVZXWFlaYWJjZGVmZ2hpamtsbW5vcHFyc3R1dnd4eXowMTIzNDU2Nzg5JTJCJTJGJTNEJyUzQiUyMF8weDM1ZmUwZiU1QidhdG9iJyU1RCUyMCU3QyU3QyUyMChfMHgzNWZlMGYlNUInYXRvYiclNUQlMjAlM0QlMjBmdW5jdGlvbiUyMChfMHg1MDA4NjYpJTIwJTdCJTIwdmFyJTIwXzB4NWJjZTE1JTIwJTNEJTIwU3RyaW5nKF8weDUwMDg2NiklNUIncmVwbGFjZSclNUQoJTJGJTNEJTJCJTI0JTJGJTJDJTIwJycpJTNCJTIwZm9yJTIwKHZhciUyMF8weDE1OTc4ZCUyMCUzRCUyMDB4MCUyQyUyMF8weDQzZjE1MyUyQyUyMF8weDQ5NTkzMSUyQyUyMF8weDNlMjVhOCUyMCUzRCUyMDB4MCUyQyUyMF8weGFlMmVhJTIwJTNEJTIwJyclM0IlMjBfMHg0OTU5MzElMjAlM0QlMjBfMHg1YmNlMTUlNUInY2hhckF0JyU1RChfMHgzZTI1YTglMkIlMkIpJTNCJTIwfl8weDQ5NTkzMSUyMCUyNiUyNiUyMChfMHg0M2YxNTMlMjAlM0QlMjBfMHgxNTk3OGQlMjAlMjUlMjAweDQlMjAlM0YlMjBfMHg0M2YxNTMlMjAqJTIwMHg0MCUyMCUyQiUyMF8weDQ5NTkzMSUyMCUzQSUyMF8weDQ5NTkzMSUyQyUyMF8weDE1OTc4ZCUyQiUyQiUyMCUyNSUyMDB4NCklMjAlM0YlMjBfMHhhZTJlYSUyMCUyQiUzRCUyMFN0cmluZyU1Qidmcm9tQ2hhckNvZGUnJTVEKDB4ZmYlMjAlMjYlMjBfMHg0M2YxNTMlMjAlM0UlM0UlMjAoLTB4MiUyMColMjBfMHgxNTk3OGQlMjAlMjYlMjAweDYpKSUyMCUzQSUyMDB4MCklMjAlN0IlMjBfMHg0OTU5MzElMjAlM0QlMjBfMHgxZGQyMTAlNUInaW5kZXhPZiclNUQoXzB4NDk1OTMxKSUzQiUyMCU3RCUyMHJldHVybiUyMF8weGFlMmVhJTNCJTIwJTdEKSUzQiUyMCU3RCgpKSUzQiUyMF8weDRkZDclNUInbkxCTVV2JyU1RCUyMCUzRCUyMGZ1bmN0aW9uJTIwKF8weDMzZGFhZCklMjAlN0IlMjB2YXIlMjBfMHhkNWZjZmElMjAlM0QlMjBhdG9iKF8weDMzZGFhZCklM0IlMjB2YXIlMjBfMHg1NDBkZTglMjAlM0QlMjAlNUIlNUQlM0IlMjBmb3IlMjAodmFyJTIwXzB4MjU0MGFkJTIwJTNEJTIwMHgwJTJDJTIwXzB4MjYxNTViJTIwJTNEJTIwXzB4ZDVmY2ZhJTVCJ2xlbmd0aCclNUQlM0IlMjBfMHgyNTQwYWQlMjAlM0MlMjBfMHgyNjE1NWIlM0IlMjBfMHgyNTQwYWQlMkIlMkIpJTIwJTdCJTIwXzB4NTQwZGU4JTIwJTJCJTNEJTIwJyUyNSclMjAlMkIlMjAoJzAwJyUyMCUyQiUyMF8weGQ1ZmNmYSU1QidjaGFyQ29kZUF0JyU1RChfMHgyNTQwYWQpJTVCJ3RvU3RyaW5nJyU1RCgweDEwKSklNUInc2xpY2UnJTVEKC0weDIpJTNCJTIwJTdEJTIwcmV0dXJuJTIwZGVjb2RlVVJJQ29tcG9uZW50KF8weDU0MGRlOCklM0IlMjAlN0QlM0IlMjBfMHg0ZGQ3JTVCJ1d2SU5VdCclNUQlMjAlM0QlMjAlN0IlN0QlM0IlMjBfMHg0ZGQ3JTVCJ25oUVVPbyclNUQlMjAlM0QlMjAhISU1QiU1RCUzQiUyMCU3RCUyMHZhciUyMF8weDI5ZjI4NiUyMCUzRCUyMF8weDRkZDclNUInV3ZJTlV0JyU1RCU1Ql8weDNiMzkxNSU1RCUzQiUyMGlmJTIwKF8weDI5ZjI4NiUyMCUzRCUzRCUzRCUyMHVuZGVmaW5lZCklMjAlN0IlMjBfMHgyNzI1NjElMjAlM0QlMjBfMHg0ZGQ3JTVCJ25MQk1VdiclNUQoXzB4MjcyNTYxKSUzQiUyMF8weDRkZDclNUInV3ZJTlV0JyU1RCU1Ql8weDNiMzkxNSU1RCUyMCUzRCUyMF8weDI3MjU2MSUzQiUyMCU3RCUyMGVsc2UlMjAlN0IlMjBfMHgyNzI1NjElMjAlM0QlMjBfMHgyOWYyODYlM0IlMjAlN0QlMjByZXR1cm4lMjBfMHgyNzI1NjElM0IlMjAlN0QlNUIuLi4lNUQ=`,
		},
	}
	// do
	for _, test := range tests {
		//testByte := []byte(test.source)
		reverse := reverseBase64Code(test.source)
		t.Logf("reverseBase64Code\nFrom: %v\nTo %v", test.source, reverse)
		// verify
		assert.Equal(t, test.want, reverse)
		reverse = fmt.Sprintf("%v", reverse)
		out, err := Base64StdDecoding(reverse)
		if err != nil {
			t.Fatalf("reverseBase64Code and reverseBase64Code error %v", err)
		}
		t.Logf("reverseBase64Code and Base64StdDecoding %v", out)
	}
}

func TestReverseBase64URLDecode(t *testing.T) {
	// mock
	tests := []struct {
		name   string
		source string
		want   string
	}{
		{
			name:   "TestReverseBase64Code URL: landscape-primary",
			source: `knch1WayBXLlBXYjNHZuFGb`,
			want:   `landscape-primary`,
		},
	}

	// do
	for _, test := range tests {
		decoding, err := ReverseBase64URLDecode(test.source)
		if err != nil {
			t.Fatalf("ReverseBase64URLDecode err at %v, info: %v", test.name, err)
		}
		//t.Logf("Base64TripleURLEncoding\nFrom: %v\nTo %v", test.source, encoding)
		// verify
		assert.Equal(t, test.want, decoding)
	}
}
