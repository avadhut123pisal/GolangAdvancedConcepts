package main

import (
	"bufio"
	"fmt"
	"time"
	"context"
	"os"
)

//Background
//WithTimeout
//WithCancel
//WithDeadline 

func main() {
	ctx:=context.Background()
	
	ctx, cancel:=context.WithCancel(ctx)
	
	// check for input
	go func() {
		scanner:=bufio.NewScanner(os.Stdin)
		scanner.Scan()
		cancel()
	}()
	
	// time.AfterFunc(2*time.Second, cancel)
	
	
	sleepAndTalk(ctx, 5*time.Second, "Hello")
	
}

func sleepAndTalk(ctx context.Context, timeValue time.Duration, message string) {
	time.Sleep(timeValue)
	fmt.Println(message)
} 

