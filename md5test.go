package main

import (
	"crypto/md5"
	"encoding/base64"
	"fmt"
	"github.com/tilinna/z85"
	"io"
)

func main() {
	h := md5.New()
	io.WriteString(h, "The fog is getting thicker!")
	io.WriteString(h, "And Leon's getting laaarger!")
	md5Encoded := h.Sum(nil)
	// 乱码
	fmt.Println(string(md5Encoded))
	// 32个16进制值
	fmt.Printf("%x\n", md5Encoded)
	// 生成的字节数
	fmt.Println(len(md5Encoded))

	base64Encoded := base64.StdEncoding.EncodeToString(md5Encoded)
	fmt.Println(base64Encoded)

	z85Encoded, _:= z85Encoding(md5Encoded)
	fmt.Println(z85Encoded)

}

func z85Encoding(src []byte) (string, error) {
	dst := make([]byte, z85.EncodedLen(len(src)))
	_, err := z85.Encode(dst, src)
	if err != nil {
		return "", err
	}

	return string(dst), nil
}