package server

import (
	"github.com/gin-gonic/gin"
	"github.com/shevonkuan/f1-telemetry-go/pkg/packets"
)

type Server struct {
	BindAddress string
	DataPoint   *packets.PacketData
}

func (p *Server) Start() {

	r := gin.Default()

	r.Use(func(context *gin.Context) {
		context.Set("packetData", p.DataPoint)
	})
	r.GET("/ws", ws)
	r.GET("/json", json)
	r.Run(p.BindAddress)
}
