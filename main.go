package main

import (
	"fmt"
	"log"
	"os/exec"
	"time"

	"github.com/shevonkuan/f1-telemetry-go/pkg/packets"
	"github.com/shevonkuan/f1-telemetry-go/pkg/telemetry"
	"github.com/shevonkuan/f1-telemetry-go/server"
)

// Start Grafana
func grafanaHandler() {
	log.Println("Start Grafana...")
	cmd := exec.Command("./grafana-9.3.6/bin/grafana-server.exe", "--homepath", "./grafana-9.3.6", "--config", "")
	_, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
}

// Start Prometheus
func prometheusHandler() {
	log.Println("Start Prometheus...")
	cmd := exec.Command("./prometheus-2.37.5.windows-amd64/prometheus.exe", "--config.file=./prometheus/prometheus.yml")
	_, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
}

// Start Server
func serverStarter(s *server.Server) {
	s.Start()
}
func main() {
	ch := make(chan struct{})
	startTime := time.Now()
	packetData := new(packets.PacketData)
	client, err := telemetry.NewClientByCustomIpAddressAndPort("0.0.0.0", 20777)
	if err != nil {
		log.Fatal(err)
	}
	// server
	svr := server.Server{
		BindAddress: "127.0.0.1:8888",
		DataPoint:   packetData,
		Client:      client,
	}
	// handler
	go server.PacketDataHandler(packetData, client, &startTime)
	go serverStarter(&svr)

	go prometheusHandler()
	go grafanaHandler()

	<-ch
	// for {
	// 	if packetData.Time != nil {
	// 		jsonBytes, err := json.Marshal(packetData)
	// 		if err != nil {
	// 			log.Fatal(err)
	// 		}
	// 		fmt.Println(string(jsonBytes))
	// 	}
	// 	time.Sleep(100 * time.Millisecond)
	// }
}
