package main

import (
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/RivanJarjes/goLinker/internal/config"
	httprouter "github.com/RivanJarjes/goLinker/internal/http"
)

func main() {
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
	cfg := config.Load()
	handler := httprouter.NewRouter()
	go func() {
		if err := http.ListenAndServe(":"+cfg.Port, handler); err != nil {
			log.Fatalf("listen: %s\n", err)
		}
	}()
	<-sigs
	log.Println("[server] Shutting down")
}
