package main

import (
    "fmt"
    "net/http"
    "time"
    "context"
    "sync"
)

func main() {

    // create a `WaitGroup` for safe exit
    wg := new(sync.WaitGroup)
    wg.Add(2) // add `2` goroutines to finish

    // create server to run on port the 9000
    server := &http.Server {
        Addr: ":9000",
        Handler: nil, // use `DefaultServeMux`
    }

    // register a cleanup function
    server.RegisterOnShutdown( func() {
        fmt.Println( "RegisterOnShutdown(): completed!" ) // perform a cleanup
        wg.Done() // WaitGroup done
    } )

    // shutdown server after 3 seconds
    time.AfterFunc( 3 * time.Second, func() {
        err := server.Shutdown( context.Background() ) // shutdown `server`
        fmt.Println( "Shutdown(): completed!", err )
        wg.Done() // WaitGroup done
    } )
    
    // launch server
    fmt.Println( "ListenAndServe():", server.ListenAndServe() )

    // listen for `wg` to complete
    wg.Wait()
    fmt.Println("Main(): exit complete!")
}