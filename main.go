package main

import (
	"fmt"
	"github.com/google/uuid"
)
import "github.com/atotto/clipboard"

func main() {
	uuid, err := uuid.NewRandom()

	if err != nil {
		panic(err)
	}

	clipboard.WriteAll(uuid.String())
	fmt.Println("UUID ready to paste")
}
