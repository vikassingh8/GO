package main

import (
	"fmt"
	"time"
	// "time"
)

//  func sendMessage(ch chan string) {
// 	ch <- "Hello, Channel!"
//  }
//  func main() {
// 	ch := make(chan string)
// 	go sendMessage(ch)
// 	fmt.Println("Received:", <-ch)
// 	time.Sleep(time.Second*2)
// }

// func sendMessage(ch chan bool) {
// 	defer func() {
// 		ch <- true
// 	}()
// 	fmt.Println("Hello, Channel!")
// }
// func main() {
// 	ch := make(chan bool)
// 	go sendMessage(ch)
// 	<-ch
// }

func sendMessage(ch chan string,done chan bool){
	defer func(){
		
		done <- true
	}()
	for email:= range ch {
		fmt.Println("Sending email to:",email)
		time.Sleep(time.Second*2)
	}

	
	fmt.Println("All emails sent!")
	
}

func main(){
	emailChain := make(chan string,15)
	done := make(chan bool)
	go sendMessage(emailChain,done)
	// done := make(chan bool)
	for i:=0;i<10;i++{
		emailChain <- fmt.Sprintf("Email %d",i)
	}
	close(emailChain)
	<-done
}