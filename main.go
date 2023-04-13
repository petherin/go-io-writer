package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"os"
)

func main() {
	fWrite()
	jsonNewEncoder()

	fmt.Print("Exiting")
}

func fWrite() {
	fmt.Println("f.Write()")

	f, err := os.OpenFile("/bin/hi.txt", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0600)
	if err != nil {
		panic(err)
	}

	defer f.Close()

	// Implements io.Writer interface
	// type Writer interface {
	//	 Write(p []byte) (n int, err error)
	// }
	n, err := f.Write([]byte("writing some data"))
	if err != nil {
		panic(err)
	}

	fmt.Printf("wrote %d bytes\n", n)
}

type user struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func jsonNewEncoder() {
	fmt.Println("json.NewEncoder")

	buf := new(bytes.Buffer)
	u := user{
		Name: "bob",
		Age:  20,
	}

	err := json.NewEncoder(buf).Encode(u)
	if err != nil {
		panic(err)
	}

	fmt.Print(buf.String())
}
