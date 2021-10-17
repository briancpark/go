package main

import (
	"fmt"
	"github.com/google/uuid"
)

func main() {
	// Create a UUID from bytes
	a := []byte("A stack of pancakes")
	aUUID, err := uuid.FromBytes(a[:16])
	if err != nil {
		fmt.Println("error:", err)
	}
	fmt.Println(aUUID)

	// Want to see me do it again?
	b := []byte("A stack of pancakes")
	bUUID, err := uuid.FromBytes(b[:16])
	if err != nil {
		fmt.Println("error:", err)
	}
	fmt.Println(bUUID)
}
