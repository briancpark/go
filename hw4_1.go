package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	type BotColor struct {
		ID    int
		name  string
		color string
	}

	// Create a struct
	group := BotColor{
		ID:    1,
		name:  "EvanBot",
		color: "Purple",
	}

	// Use json.Marshal to compress the struct
	// into a byte array
	marshalBot, err := json.Marshal(group)
	if err != nil {
		fmt.Println("error:", err)
	}

	// use json.Unmarshal to decompress the struct
	var unmarshalBot BotColor
	err = json.Unmarshal(marshalBot, &unmarshalBot)
	if err != nil {
		fmt.Println("error:", err)
	}
	fmt.Printf("%+v", unmarshalBot)
}