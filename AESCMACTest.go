package main

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"fmt"
	"github.com/aead/cmac"
	"github.com/tilinna/z85"
)

const PUBLIC_KEY string = "12345678abcdefgh"

func main() {
	//key的长度必须是16、24或者32字节，分别用于选择AES-128, AES-192, or AES-256
	var aeskey = []byte("12345678vbcdefgh")
	pass := []byte("vdncloud123456")
	xpass, mac, err := AesEncrypt(pass, aeskey)
	fmt.Println(xpass)
	if err != nil {
		fmt.Println(err)
		return
	}

	//pass64 := base64.StdEncoding.EncodeToString(xpass)
	pass85 := make([]byte, z85.EncodedLen(len(xpass)))
	_, err = z85.Encode(pass85, xpass)
	if err != err {
		//t.Fatalf("got %q, want %q", err, test.err)
	}
	if err == nil {
		//if !bytes.Equal([]byte(test.z85), dst) {
		//	t.Fatalf("got %q, want %q", dst, test.z85)
		//}
		//if n != len(test.z85) {
		//	t.Fatalf("got %d, want %d", n, len(test.z85))
		//}
	}
	fmt.Printf("加密后:%v\n", string(pass85))

	//bytesPass, err := base64.StdEncoding.DecodeString(pass64)

	pass85 = []byte("}:fe:r-%ANc>fy}!zhUV")
	dst := make([]byte, z85.DecodedLen(len(pass85)))
	_, err = z85.Decode(dst, pass85)
	tpass, err := AESDecrypt(dst, aeskey, mac)

	if err != err {
		//t.Fatalf("got %q, want %q", err, test.err)
	}
	if err == nil {
		//if !bytes.Equal([]byte(test.z85), dst) {
		//	t.Fatalf("got %q, want %q", dst, test.z85)
		//}
		//if n != len(test.z85) {
		//	t.Fatalf("got %d, want %d", n, len(test.z85))
		//}
	}
	fmt.Printf("解密后:%s\n", tpass)
}

func PKCS5Padding(plainText []byte, blockSize int) []byte {
	padding := blockSize - len(plainText)%blockSize
	pddText := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(plainText, pddText...)
}

func PKCS5UnPadding(origdata []byte) []byte {
	length := len(origdata)
	unpadding := int(origdata[length-1])
	return origdata[:length-unpadding]
}

func AesEncrypt(orgData, key []byte) ([]byte, []byte, error) {
	block, err := aes.NewCipher(key)

	if err != nil {
		return nil, nil, err
	}

	blockSize := block.BlockSize()
	orgData = PKCS5Padding(orgData, blockSize)
	blockMode := cipher.NewCBCEncrypter(block, key[:blockSize])
	crypted := make([]byte, len(orgData))
	blockMode.CryptBlocks(crypted, orgData)

	mac, err := cmac.Sum(orgData, block, blockSize)
	if err != nil {
		return nil, nil, err
	}
	return crypted, mac, nil
}

func AESDecrypt(crypted, key, mac []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	//AES分组长度为128位，所以blockSize=16，单位字节
	blockSize := block.BlockSize()
	blockMode := cipher.NewCBCDecrypter(block, key[:blockSize]) //初始向量的长度必须等于块block的长度16字节
	origData := make([]byte, len(crypted))
	blockMode.CryptBlocks(origData, crypted)

	//verifyResult := cmac.Verify(mac, origData, block, blockSize)

	origData = PKCS5UnPadding(origData)

	//if !verifyResult {
	//	return nil, err
	//}
	return origData, nil
}
