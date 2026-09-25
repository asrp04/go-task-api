package main

import (
	"fmt"
	"time"
)

func sendEmail(user string, ch chan string) {
	fmt.Printf("📧 Email processing starts for %s ...\n", user)
	time.Sleep(2 * time.Second) 
	ch <- fmt.Sprintf("✅ work completed for %s", user)
}

func main() {
	startTime := time.Now()
	ch := make(chan string)
	go sendEmail("User A", ch)
	go sendEmail("User B", ch)
	go sendEmail("User C", ch)
	
	msg1 := <-ch
	msg2 := <-ch
	msg3 := <-ch
	
	fmt.Println(msg1)
	fmt.Println(msg2)
	fmt.Println(msg3)
	fmt.Printf(" Total time taken : %v\n", time.Since(startTime))
}
