package main

import (
	"fmt"
)

func main() {
	// Membuat unbuffered channel untuk tipe data int
	ch := make(chan int)

	// Goroutine untuk mengirim data
	go func() {
		fmt.Println("Mengirim data ke channel...")
		ch <- 42 // Mengirim data ke channel
		fmt.Println("Data telah dikirim.")
	}()

	// Menerima data dari channel
	fmt.Println("Menunggu menerima data...")
	value := <-ch // Menerima data dari channel
	fmt.Println("Data diterima:", value)
}
