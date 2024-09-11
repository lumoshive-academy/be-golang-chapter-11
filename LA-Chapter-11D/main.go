package main

import (
	"fmt"
	"time"
)

func printWord(word string, count int) {
	for i := 0; i < count; i++ {
		fmt.Println(word)
	}
}

func main() {
	// Menjalankan goroutine untuk mencetak "Hello" sebanyak 100 kali
	go printWord("Hello", 100)

	// Menjalankan goroutine untuk mencetak "World" sebanyak 100 kali
	printWord("World", 100)

	go func() {
		fmt.Println("Hello from anonymous goroutine!")
	}()
	time.Sleep(1 * time.Second) // Menunggu sejenak agar goroutine bisa selesai
}
