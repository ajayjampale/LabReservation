package main

import (
	"context"
	"log"
	"net/http"
	"sync/atomic"
	"syscall"
	"time"

	"os"
	"os/signal"

	"github.com/gorilla/mux"
)

//LRMServer denotes Lab Resource Management HTTP Server.
type LRMServer struct {
	http.Server
	shutdownReq chan bool
	reqCount    uint32
}

func NewLRMServer() *LRMServer {
	//create server
	s := &LRMServer{
		Server: http.Server{
			Addr:         ":8080",
			ReadTimeout:  10 * time.Second,
			WriteTimeout: 10 * time.Second,
		},
		shutdownReq: make(chan bool),
	}
	router := mux.NewRouter()

	//register handlers
	router.HandleFunc("/", s.RootHandler)
	router.HandleFunc("/shutdown", s.ShutdownHandler)

	//set http server handler
	s.Handler = router

	return s
}

func (s *LRMServer) WaitForGracefulShutdown() {
	irqSig := make(chan os.Signal, 1)
	signal.Notify(irqSig, syscall.SIGINT, syscall.SIGTERM)

	log.Printf("Waiting for shut signals ...")

	//Wait interrupt or shutdown request through /shutdown
	select {
	case sig := <-irqSig:
		log.Printf("Shutdown request (signal: %v)", sig)
	case sig := <-s.shutdownReq:
		log.Printf("Shutdown request (/shutdown %v)", sig)
	}

	log.Printf("Stoping http server ...")

	//Create shutdown context with 10 second timeout
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	//shutdown the server
	err := s.Shutdown(ctx)
	if err != nil {
		log.Printf("Shutdown request error: %v", err)
	}
}

func (s *LRMServer) RootHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello Gorilla MUX!\n"))
}

func (s *LRMServer) ShutdownHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Shutdown server"))

	//Do nothing if shutdown request already issued
	//if s.reqCount == 0 then set to 1, return true otherwise false
	if !atomic.CompareAndSwapUint32(&s.reqCount, 0, 1) {
		log.Printf("Shutdown through API call in progress...")
		return
	}

	go func() {
		s.shutdownReq <- true
	}()
}

func main() {
	//Start the lrmServer
	lrmServer := NewLRMServer()

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