package handlers

import (
	"fmt"
	"log"
	"net/http"
	"os/exec"
)

type Sync struct {
	l *log.Logger
}

func NewSync(l *log.Logger) *Sync {
	return &Sync{l}
}

func (s *Sync) GetSync(rw http.ResponseWriter, r *http.Request) {
	cmd := exec.Command("youtube-dl", "-s", "--config-location", "/home/queb/.config/youtube-dl/autoconfig")

	err := cmd.Run()

	if err != nil {
		s.l.Printf("What went wrong: %s", err)
		http.Error(rw, "Something went wrong", http.StatusBadRequest)
		return
	}

	fmt.Fprintf(rw, "Hey!, how you doin'?")
}
