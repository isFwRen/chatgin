package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"testing"
)

func TestChat(t *testing.T) {
	fmt.Println("********")
}

func TestMd5(t *testing.T) {
	h := md5.New()
	h.Write([]byte("The fog is getting thicker!"))
	temStr := h.Sum(nil)
	fmt.Printf("%v\n", hex.EncodeToString(temStr))
}
