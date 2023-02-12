package main

import (
	"log"
	"time"

	"github.com/shevonkuan/f1-telemetry-go/pkg/packets"
	"github.com/shevonkuan/f1-telemetry-go/pkg/telemetry"
	"github.com/shevonkuan/f1-telemetry-go/server"
)

func packetDataHandler(p *packets.PacketData, c *telemetry.Client, startTime *time.Time) {
	// handler
	timeCost := 0 * time.Second
	c.OnMotionPacket(func(packet *packets.PacketMotionData) {
		car := packet.Self()
		timeCost = time.Since(*startTime)
		p.Time = &timeCost
		p.PacketMotionData = &car
	})
	c.OnSessionPacket(func(packet *packets.PacketSessionData) {
		timeCost = time.Since(*startTime)
		p.Time = &timeCost
		p.PacketSessionData = packet
	})
	c.OnLapPacket(func(packet *packets.PacketLapData) {
		car := packet.Self()
		timeCost = time.Since(*startTime)
		p.Time = &timeCost
		p.PacketLapData = &car
	})
	c.OnEventPacket(func(packet *packets.PacketEventData) {
		car := packet.EventCodeString()
		timeCost = time.Since(*startTime)
		p.Time = &timeCost
		p.PacketEventData = &car
	})
	c.OnParticipantsPacket(func(packet *packets.PacketParticipantsData) {
		car := packet.Self()
		timeCost = time.Since(*startTime)
		p.Time = &timeCost
		p.PacketParticipantsData = &car
	})
	c.OnCarSetupPacket(func(packet *packets.PacketCarSetupData) {
		car := packet.Self()
		timeCost = time.Since(*startTime)
		p.Time = &timeCost
		p.PacketCarSetupData = &car
	})
	c.OnCarStatusPacket(func(packet *packets.PacketCarStatusData) {
		car := packet.Self()
		timeCost = time.Since(*startTime)
		p.Time = &timeCost
		p.PacketCarStatusData = &car
	})
	c.OnFinalClassificationPacket(func(packet *packets.PacketFinalClassificationData) {
		car := packet.Self()
		timeCost = time.Since(*startTime)
		p.Time = &timeCost
		p.PacketFinalClassificationData = &car
	})
	c.OnLobbyInfoPacket(func(packet *packets.PacketLobbyInfoData) {
		car := packet.Self()
		timeCost = time.Since(*startTime)
		p.Time = &timeCost
		p.PacketLobbyInfoData = &car
	})
	c.OnCarDamagePacket(func(packet *packets.PacketCarDamageData) {
		car := packet.Self()
		timeCost = time.Since(*startTime)
		p.Time = &timeCost
		p.PacketCarDamageData = &car
	})
	c.OnSessionHistoryPacket(func(packet *packets.PacketSessionHistoryData) {
		timeCost = time.Since(*startTime)
		p.Time = &timeCost
		p.PacketSessionHistoryData = packet
	})

	c.OnCarTelemetryPacket(func(packet *packets.PacketCarTelemetryData) {
		car := packet.Self()
		timeCost = time.Since(*startTime)
		p.Time = &timeCost
		p.PacketCarTelemetryData = &car
	})
	// start
	c.Run()
}

func main() {
	startTime := time.Now()
	packetData := new(packets.PacketData)
	client, err := telemetry.NewClientByCustomIpAddressAndPort("0.0.0.0", 20777)
	if err != nil {
		log.Fatal(err)
	}
	// handler
	go packetDataHandler(packetData, client, &startTime)

	// server
	svr := server.Server{
		BindAddress: "127.0.0.1:8888",
		DataPoint:   packetData,
	}
	svr.Start()

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
