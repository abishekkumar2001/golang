package main

import (
	"context"
	"fmt"
	"log"
)

type key int

const keyvalue key = 12345

func main() {
	ctx := context.WithValue(context.Background(), keyvalue, 23456)
	keyvalue1 := ctx.Value(keyvalue)
	i, ok := keyvalue1.(int)
	if !ok {
		log.Fatal("not an integer")
	}
	fmt.Println("value: ", i)
}
