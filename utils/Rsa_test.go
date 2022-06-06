package utils

import (
    "fmt"
    "testing"
)

const (
    PublicPath  = "pem/publicKey.pem"
    PrivatePath = "pem/privateKey.pem"
)

func TestRsaGenKey(t *testing.T) {
    err := RsaGenKey(2048, PrivatePath, PublicPath)
    if err != nil {
        return
    }
}

func TestRsaEncrypt(t *testing.T) {
    msg := "1000"
    // 公钥加密
    cipherText, err := RSAEncrypt([]byte(msg), PublicPath)
    if err != nil {
        return
    }
    fmt.Println(string(cipherText))
    
    plainText, err := RSADecrypt(cipherText, PrivatePath)
    if err != nil {
        return
    }
    fmt.Println(string(plainText))
}
