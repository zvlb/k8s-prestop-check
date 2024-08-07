package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
)

func main() {
	go func() {
		coutchSignals()
		os.Exit(111)
	}()

	// simple http server
	server := gin.New()

	h := &handler{}

	routes := server.Group("/")
	routes.GET("/sleep", h.Wait15sec)

	if err := server.Run(":8080"); err != nil {
		log.Panic(err)
	}

}

type handler struct {
}

func (h *handler) Wait15sec(ctx *gin.Context) {
	fmt.Printf("%+v. %s\n", time.Now().Unix(), "Sleep request run")
	time.Sleep(15 * time.Second)
	fmt.Printf("%+v. %s\n", time.Now().Unix(), "Sleep request done")
	ctx.JSON(200, "Sleep is done")
}

func coutchSignals() {
	sigs := make(chan os.Signal, 1)

	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	done := make(chan bool, 1)

	go func() {
		sig := <-sigs
		fmt.Println()
		fmt.Printf("%+v. %s\n", time.Now().Unix(), sig)
		done <- true
	}()

	fmt.Printf("%+v. %s\n", time.Now().Unix(), "awaiting signal")
	<-done
	fmt.Printf("%+v. %s\n", time.Now().Unix(), "exiting")
}
