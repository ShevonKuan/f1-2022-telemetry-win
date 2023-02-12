package main

import (
	"fmt"
	"log"

	"github.com/shevonkuan/f1-telemetry-go/pkg/packets"
	"github.com/shevonkuan/f1-telemetry-go/pkg/telemetry"
)

func main() {
	client, err := telemetry.NewClientByCustomIpAddressAndPort("0.0.0.0", 20777)
	if err != nil {
		log.Fatal(err)
	}

	client.OnCarTelemetryPacket(func(packet *packets.PacketCarTelemetryData) {
		car := packet.Self()
		fmt.Println(car)
		fmt.Printf("Speed: %d\n", car.Speed)
		fmt.Printf("Engine RPM: %d\n", car.EngineRPM)
	})

	client.Run()
}
