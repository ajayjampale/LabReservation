package main

import (
	"log"
	"app-server/infra/rest/server"
)


func main() {
	//Start the lrmServer
	lrmServer := server.NewLRMServer()

	serverDone := make(chan bool)
	go func() {
		err := lrmServer.ListenAndServe()
		if err != nil {
			log.Printf("lrmServer ListenAndServe errored : %v", err)
		}
		serverDone <- true
	}()

	lrmServer.WaitForGracefulShutdown()

	<-serverDone
	log.Printf("Exiting ...")
}