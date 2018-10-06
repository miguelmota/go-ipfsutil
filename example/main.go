package main

import (
	"fmt"
	"log"

	"github.com/miguelmota/go-ipfsutil"
)

func main() {
	hash, err := ipfsutil.AddBytes([]byte("hello world"))
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(hash) // Qmf412jQZiuVUtdgnB36FXFX7xg5V6KEbSJ4dpQuhkLyfD
}
