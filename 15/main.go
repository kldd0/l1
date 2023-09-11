package main

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"
)

var justString string

func createHugeString(length int) string {
	b := make([]byte, length)
	_, err := rand.Read(b)
	if err != nil {
		panic(err)
	}
	res := base64.StdEncoding.EncodeToString(b)[:length]
	return res
}

func someFunc() {
    /*
	потенциально сохраняется вся большая строка,
    так как slice - всего лишь указатель, который указывает на кусок памяти

    type _string struct {
        elements *byte // underlying bytes
        len      int   // number of bytes
    }
    */

	v := createHugeString(1 << 10)
	tmp := make([]byte, 100)
	copy(tmp, v)
	justString = string(tmp)
}

func main() {
	someFunc()
	fmt.Println(justString)
}
