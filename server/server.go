package server

import (
	"log"
	"os"
	"os/signal"

	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/shevonkuan/f1-telemetry-go/pkg/packets"
	"github.com/shevonkuan/f1-telemetry-go/pkg/telemetry"
)

type Server struct {
	BindAddress string
	DataPoint   *packets.PacketData
	Client      *telemetry.Client
}

var wheelOrderArr = []string{"rl", "rr", "fl", "fr"}

func (p *Server) Start() {

	r := gin.Default()

	r.Use(func(context *gin.Context) {
		context.Set("packetData", p.DataPoint)
	})
	r.GET("/ws", ws)
	r.GET("/json", json)

	// prometheus handler
	r.GET("/metrics", func(c *gin.Context) {
		promhttp.Handler().ServeHTTP(c.Writer, c.Request)
	})

	// wait exit signal
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	go func() {
		<-c
		log.Printf("Packet RecvCount: %d\n", p.Client.Stats.RecvCount())
		log.Printf("Packet ErrCount: %d\n", p.Client.Stats.ErrCount())
		os.Exit(1)
	}()

	log.Println("F1 telemetry client running")
	r.Run(p.BindAddress)

}
