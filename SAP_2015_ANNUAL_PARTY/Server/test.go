package main

import (
	"fmt"
//	"reflect"
	"runtime"
//	"log"
	"net/http"
//	"strings"
//	"os"
//	"crypto/aes"
//	"crypto/cipher"
	log "github.com/cihub/seelog"
)

var commonIV = []byte{0x00,0x01,0x02,0x03,0x04,0x05,0x06,0x07}
func say(s string) {
	for i:=0;i<10;i++{
		runtime.Gosched()
		fmt.Println(s)
	}
}

func sayhelloName(w http.ResponseWriter, r *http.Request)  {
	r.ParseForm()
	fmt.Println(r.Form)
	fmt.Println("path", r.URL.Path)
	fmt.Println("scheme", r.URL.Scheme)
	fmt.Println(r.Form["url_long"])
	/*
	for k, v := range r.Form {
		fmt.Println("key: ", k)
		fmt.Println("value: ", strings.Join(v, ""))
	}
	*/
	fmt.Fprintf(w, "hello bbknightzm")
}

func main() {
	defer log.Flush()
	log.Info("hello")
	/*
	plaintext := []byte("my name is bbknightzm")
	if len(os.Args) > 1 {
		plaintext = []byte(os.Args[1])
	}
	key_text := "astaxie12798akljzmknm.ahkjkljl;k"
	if len(os.Args) > 2 {
		key_text = os.Args[2]
	}
	fmt.Println(len(key_text))

	c, err := aes.NewCipher([]byte(key_text))
	if err != nil {
		fmt.Printf("Error: NewCipher(%d bytes) = %s", len(key_text), err)
		os.Exit(-1)
	}

	cfb := cipher.NewCFBEncrypter(c, commonIV)
	ciphertext := make([]byte, len(plaintext))
	cfb.XORKeyStream(ciphertext, plaintext)
	fmt.Printf("%s=>%x\n", plaintext, ciphertext)

	cfbdec := cipher.NewCFBDecrypter(c, commonIV)
	plaintextCopy := make([]byte, len(plaintext))
	cfbdec.XORKeyStream(plaintextCopy, ciphertext)
	fmt.Printf("%s=>%x\n", ciphertext, plaintextCopy)
	
	defer fmt.Println()
	fmt.Printf("test a\n")
	for i := 0; i<5; i++ {
		defer fmt.Printf("%d ",i)
	} 
	fmt.Printf("test b\n")

	var x int = 3
	v := reflect.ValueOf(x)
	fmt.Println("type: ", v.Type())
	fmt.Println("kind: ", v.Kind() == reflect.Float64)
	fmt.Println("value:", v.Int())

	go say("world")
	say("hello")

	http.HandleFunc("/". sayhelloName)
	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		log.Fatal("ListenAndServer: ", err)
	}
	*/
}