package main

import(
	"fmt"
)

func main() {
	msgchan := make(chan string)
	// chan in Go prevent race-condition between Gorountines it works as a pipe you put some-value on one end and it comes out the other end, there is no risk of the two Gorountines accessing what-ever value is passed in the same time (chan(as a pipe): provide safe way to path values between Gorountines).
	go func(){
			msgchan <- "Hello"
			// "Hello" is created in a seprate Goroutine(func) and sent to a seprated Goroutine(main) via a channel{chan} then it's printed. 
		}()
	msg := <- msgchan
	//seprated Goroutine that running Goroutine(main) 
	fmt.Println(msg)	
}