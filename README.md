[![Build Status](https://travis-ci.org/anisse/proxylistener.svg?branch=master)](https://travis-ci.org/anisse/proxylistener)
# proxylistener

A Go [net.Listener](https://golang.org/pkg/net/#Listener) that understands the [haproxy PROXY protocol](www.haproxy.org/download/1.8/doc/proxy-protocol.txt). Uses [go-proxyproto](https://github.com/pires/go-proxyproto).


## Usage
```go
func serveHere() error {
	l, err := proxylistener.Listen("tcp", "localhost:8080") // you can also use a unix socket if you wish
	if err != nil {
		return err
	}
	return http.Serve(l, FileServer("."))
}
```
