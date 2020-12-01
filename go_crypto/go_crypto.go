// go_crypto.go
package main

import (
	"crypto/md5"
	"encoding/base64"
	"fmt"
	"io"
)

func main() {
	h := md5.New()
	io.WriteString(h, "123456")
	fmt.Printf("md5(123456)=%x\n", h.Sum(nil))

	b := []byte("123456")
	s := base64.StdEncoding.EncodeToString(b)
	fmt.Printf("base64.encode(123456)=%s\n", s)

	t, _ := base64.StdEncoding.DecodeString(s)
	fmt.Printf("base64.decode(s)=%s\n", t)
}
