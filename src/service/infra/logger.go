package infra

import (
	"context"
	"fmt"
	"github.com/gorilla/mux"
	"io"
	"log"
	"net/http"
	"os"
)

func loggerman() {

	f, err := os.OpenFile("C:\\Users\\gs-0813\\Documents\\MyProjects\\bredec\\logger.log",
		os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Println(err)
	}
	defer f.Close()

	logger := log.New(f, "prefix", log.LstdFlags)
	logger.Println("text to append")
	logger.Println("more text to append")
}


func StartHttpServer() *http.Server {
	r := mux.NewRouter()
	srv := &http.Server{Addr: ":8080", Handler:r}
	elog.Info(1,"create server on 8080 with mux router")
	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, "hello world\n")
	})
	//http.Handle("/",r)

	go func() {
		// returns ErrServerClosed on graceful close
		if err := srv.ListenAndServe(); err != http.ErrServerClosed {
			// NOTE: there is a chance that next line won't have time to run,
			// as main() doesn't wait for this goroutine to stop. don't use
			// code with race conditions like these for production. see post
			// comments below on more discussion on how to handle this.
			msg := fmt.Sprintf("ListenAndServe() error : %s", err)
			elog.Error(1,msg)
		}
	}()

	// returning reference so caller can call Shutdown()
	return srv
}

func ShutdownHttpServer(server *http.Server){
	if err := server.Shutdown(context.TODO()); err != nil {
		msg := fmt.Sprintf("Stop server error(): %s", err)
		elog.Error(1,msg)
		panic(err) // failure/timeout shutting down the server gracefully
	}
}