# github.com/stretchr/testify

[github.com/stretchr/testify](https://github.com/stretchr/testify)

```bash
GO111MODULE=on go mod edit -require='github.com/stretchr/testify@v1.4.0'
GO111MODULE=on go mod vendor
```

fast use

```go
import (
  "testing"
  "github.com/stretchr/testify/assert"
)

func TestSomething(t *testing.T) {

  // assert equality
  assert.Equal(t, 123, 123, "they should be equal")

  // assert inequality
  assert.NotEqual(t, 123, 456, "they should not be equal")

  // assert for nil (good for errors)
  assert.Nil(t, object)

  // assert for not nil (good when you expect something)
  if assert.NotNil(t, object) {

    // now we know that object isn't nil, we are safe to make
    // further assertions without causing any errors
    assert.Equal(t, "Something", object.Value)

  }
}
```

# github.com/bar-counter/monitor

[github.com/bar-counter/monitor](https://github.com/bar-counter/monitor)

```bash
GO111MODULE=on go mod edit -require='github.com/bar-counter/monitor@v1.1.0'
GO111MODULE=on go mod vendor
```

# gin

- source [https://github.com/gin-gonic/gin](https://github.com/gin-gonic/gin)

```bash
go list -m -versions github.com/gin-gonic/gin
GO111MODULE=on go mod edit -require='github.com/gin-gonic/gin@v1.5.0'
GO111MODULE=on go mod vendor
```

#  github.com/parnurzeal/gorequest

- source [https://github.com/parnurzeal/gorequest](https://github.com/parnurzeal/gorequest)


```bash
go list -m -versions github.com/parnurzeal/gorequest
GO111MODULE=on go mod edit -require='github.com/parnurzeal/gorequest@v0.2.16'
GO111MODULE=on go mod vendor
```

```go
import "github.com/parnurzeal/gorequest"

request := gorequest.New()
```

- POST json

```go
// JSON is for sure a default
resp, body, errs := request.Post("https://httpbin.org/post").
  Set("Notes","gorequst is coming!"). // set Header
  Send(`{"name":"backy", "species":"dog"}`).
  End()
```
```bash
curl 'https://httpbin.org/post' \
  -X POST \
  -H "Content-Type: application/json" \
  -H 'Notes: gorequst is coming!' \
  -d '{"name":"backy", "species":"dog"}'
```

- form
```go
// also TypeFormData TypeForm
// application/x-www-form-urlencoded
gorequest.New().Post("http://example.com/").
  Type(gorequest.TypeUrlencoded).
  Send(`{"name":"backy", "species":"dog"}`).
  End()
```
```bash
curl 'https://httpbin.org/post' \
  -X POST \
  -H "Content-Type: application/x-www-form-urlencoded" \
  -H 'Notes: gorequst is coming!' \
  -d 'name=backy&species=dog'
```
- form-data
```go
// multipart/form-data
f, _ := filepath.Abs("./file2.txt")
bytesOfFile, _ := ioutil.ReadFile(f)

gorequest.New().Post("http://example.com/").
  Type(gorequest.TypeMultipart).
  SendFile("./file1.txt").
  SendFile(bytesOfFile, "file2.txt", "my_file_fieldname").
  End()
```
