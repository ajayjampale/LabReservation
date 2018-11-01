package server

import (
	"log"
	"context"
	"net/http"
	"time"
	"os"
	"os/signal"
	"syscall"
	"sync/atomic"
	"app-server/infra/rest"
)

//LRMServer denotes Lab Resource Management HTTP Server.
type LRMServer struct {
	http.Server
	ShutDownRequest      chan bool
	ShutDownRequestCount uint32
}

func NewLRMServer() *LRMServer {
	//create server
	s := &LRMServer{
		Server: http.Server{
			Addr:         ":8081",
			ReadTimeout:  10 * time.Second,
			WriteTimeout: 10 * time.Second,
		},
		ShutDownRequest: make(chan bool),
	}
	router := rest.NewLRMRouter()

	//register handlers
	router.HandleFunc("/", s.RootHandler)
	router.HandleFunc("/shutdown", s.ShutdownHandler)

	//set http server handler
	s.Handler = router

	return s
}

func (s *LRMServer) RootHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello Gorilla MUX!\n"))
}

func (s *LRMServer) ShutdownHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Shutdown server"))

	//Do nothing if shutdown request already issued
	//if s.reqCount == 0 then set to 1, return true otherwise false
	if !atomic.CompareAndSwapUint32(&s.ShutDownRequestCount,0, 1) {
		log.Printf("Shutdown through API call in progress...")
		return
	}

	go func() {
		s.ShutDownRequest <- true
	}()
}

func (s *LRMServer) WaitForGracefulShutdown() {
	irqSig := make(chan os.Signal, 1)
	signal.Notify(irqSig, syscall.SIGINT, syscall.SIGTERM)

	log.Printf("Waiting for shutdown signals ...")

	//Wait interrupt or shutdown request through /shutdown
	select {
	case sig := <-irqSig:
		log.Printf("Shutdown request (signal: %v)", sig)
	case sig := <-s.ShutDownRequest:
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