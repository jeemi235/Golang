package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {

	// create channels for safe exit
	exitSignal := make(chan interface{})

	// create server to run on port the 9000
	server := &http.Server{
		Addr:    ":9000",
		Handler: nil, // use 'DefaultServeMux'
	}

	// close server after 3 seconds
	time.AfterFunc(3*time.Second, func() {
		fmt.Println("Close(): completed!", server.Close()) // close `server`
		close(exitSignal)                                  // close `exitSignal` channel
	})

	// launch server
	err := server.ListenAndServe()
	fmt.Println("ListenAndServe():", err)

	// listen to `exitSignal` channel
	<-exitSignal
	fmt.Println("Main(): exit complete!")
}

/*
In the above program, we have created a bare minimal server and spawned it in the main method. We have also created a exitSignal channel to block the main goroutine unless the server is closed.
Even after server exits, the main goroutine is blocked until unless some goroutine writes a value to or closes the exitSignal channel.

Using time.AfterFunc function, we have launched another goroutine that executes a function after 3 seconds. In this function, we are calling server.Close() method which closes the server.
This function call may return an error but before that happens, the server may close already and exit with an error.

This is why exitSignal channel is so important in this scenario. Once server.Close() is returned, we are closing the exitSignal channel using built-in close function.

Once the exitSignal channel is closed, this goroutine will be killed and control will pass back to main goroutine. Since exitSignal channel is now closed, main goroutine is no longer
 blocked and it will start executing the code below the read from channel (<- chan) expression.

*/
