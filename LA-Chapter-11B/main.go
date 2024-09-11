package main

import (
	"fmt"

	"github.com/google/uuid"
)

func main() {
	newUUID := uuid.New()
	fmt.Println("New UUID:", newUUID.String())
}
