package main

import (
	"crypto/aes"
    "crypto/cipher"
    "crypto/rand"
    "fmt"
    "io"
    //"io/ioutil"
)

func GenerateId() string {
	b := make([]byte, 16)
	rand.Read(b)
	return fmt.Sprintf("%x", b)
}


func Encryptor(secretkey, content  string) []byte {
    fmt.Println(" --encryption program run v0.01")

    key := []byte(secretkey)
    text := []byte(content)

    c, err := aes.NewCipher(key)
    if err != nil {
        fmt.Println(err)
    }
    gcm, err := cipher.NewGCM(c)
    if err != nil {
        fmt.Println(err)
    }

    nonce := make([]byte, gcm.NonceSize())
    if _, err = io.ReadFull(rand.Reader, nonce); err != nil {
        fmt.Println(err)
    }

    x := gcm.Seal(nonce, nonce, text, nil)

/*	err = ioutil.WriteFile("file.data", x, 0777)
	if err != nil {
		fmt.Println(err)
	}*/

	return x

}