# go-request

[![Go Reference](https://pkg.go.dev/badge/github.com/hb0730/go-request.svg)](https://pkg.go.dev/github.com/hb0730/go-request)

http request

# install

```yaml
go get github.com/hb0730/go-request
```

# Example

```yaml
import "github.com/hb0730/go-request"

func Example(){
    req:=request.CreateRequest(http.MethodPost,"https://exmple.com","")
    _=req.Do()
    resp:=req.GetResponse()
}
```
