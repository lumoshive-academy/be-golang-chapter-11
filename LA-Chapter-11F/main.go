package main

import (
	"fmt"
)

// function yang mengirim data ke channel
func sendData(ch chan int, value int) {
	ch <- value // Mengirim nilai ke channel
}

// function send hanya mengirim data ke channel
func send(ch chan<- int, value int) {
	fmt.Println("Mengirim data ke channel:", value)
	ch <- value // Mengirim nilai ke channel
}

// function receive hanya menerima data dari channel
func receive(ch <-chan int, done chan<- bool) {
	value := <-ch // Menerima nilai dari channel
	fmt.Println("Data diterima di function:", value)
	done <- true // Mengirim sinyal selesai ke channel done
}

func main() {
	ch := make(chan int)

	// Menjalankan goroutine untuk mengirim data
	go sendData(ch, 42)

	// Menerima data dari channel
	result := <-ch
	fmt.Println("Data diterima:", result)

	// implementasai direction channel
	dataCh := make(chan int)
	doneCh := make(chan bool)

	// Menjalankan goroutine untuk mengirim data
	go send(dataCh, 42)

	// Menjalankan goroutine untuk menerima data
	go receive(dataCh, doneCh)

	// Menunggu sinyal selesai dari goroutine receive
	<-doneCh
	fmt.Println("Semua pekerjaan selesai.")

}
