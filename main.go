package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	mux "github.com/gorilla/mux"
	"github.com/hekatx/yt_mus/handlers"
)

func main() {
	l := log.New(os.Stdout, "yt-api", log.LstdFlags)
	r := mux.NewRouter()

	sh := handlers.NewSync(l)

	r.HandleFunc("/sync", sh.GetSync)

	srv := &http.Server{
		Handler:      r,
		ErrorLog:     l,
		Addr:         "localhost:8080",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	go func() {
		err := srv.ListenAndServe()
		if err != nil {
			l.Fatal(err)
		}
	}()

	sigChan := make(chan os.Signal)
	signal.Notify(sigChan, os.Interrupt)
	signal.Notify(sigChan, os.Kill)

	sig := <-sigChan
	l.Println("Received termiante, graceful shutdown", sig)

	tc, _ := context.WithTimeout(context.Background(), 30*time.Second)
	srv.Shutdown(tc)

}
