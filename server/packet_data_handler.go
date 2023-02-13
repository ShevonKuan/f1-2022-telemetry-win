package server

import (
	"time"

	"github.com/shevonkuan/f1-telemetry-go/pkg/env/event"
	"github.com/shevonkuan/f1-telemetry-go/pkg/packets"
	"github.com/shevonkuan/f1-telemetry-go/pkg/telemetry"
)

// Receive data from F1 2021 and pakcet to PacketData
func PacketDataHandler(p *packets.PacketData, c *telemetry.Client, startTime *time.Time) {
	// handler
	timeCost := 0 * time.Second
	p.Time = &timeCost
	c.OnMotionPacket(func(packet *packets.PacketMotionData) {
		car := packet.Self()
		timeCost = time.Since(*startTime)
		p.PacketMotionData = &car
		accelerationMetric.WithLabelValues("x").Set(float64(car.GForceLateral))
		accelerationMetric.WithLabelValues("y").Set(float64(car.GForceLongitudinal))
		accelerationMetric.WithLabelValues("z").Set(float64(car.GForceVertical))
		attitudeMetric.WithLabelValues("Roll").Set(float64(car.Roll))
		attitudeMetric.WithLabelValues("Pitch").Set(float64(car.Pitch))
		attitudeMetric.WithLabelValues("Yaw").Set(float64(car.Yaw))
	})
	c.OnSessionPacket(func(packet *packets.PacketSessionData) {
		timeCost = time.Since(*startTime)
		p.PacketSessionData = packet
	})
	c.OnLapPacket(func(packet *packets.PacketLapData) {
		car := packet.Self()
		timeCost = time.Since(*startTime)
		p.PacketLapData = &car
		// metrics
		lastLapTimeMetric.Set(float64(car.LastLapTimeInMS))
		currentLapTimeMetric.Set(float64(car.CurrentLapTimeInMS))
		currentLapNumMetric.Set(float64(car.CurrentLapNum))
	})
	c.OnEventPacket(func(packet *packets.PacketEventData) {
		car := packet.EventCodeString()
		timeCost = time.Since(*startTime)
		p.PacketEventData = &car
		// metrics
		switch packet.EventCodeString() {

		case event.SpeedTrapTriggered:
			trap := packet.EventDetails.(*packets.SpeedTrap)
			if trap.VehicleIdx == packet.Header.PlayerCarIndex {
				speedTrapMetric.Set(float64(trap.Speed))
			}
			break
		case event.FastestLap:
			fp := packet.EventDetails.(*packets.FastestLap)
			if fp.VehicleIdx == packet.Header.PlayerCarIndex {
				fastestLapMetric.Set(float64(fp.LapTime))
			}
			break
		}
	})
	c.OnParticipantsPacket(func(packet *packets.PacketParticipantsData) {
		car := packet.Self()
		timeCost = time.Since(*startTime)
		p.PacketParticipantsData = &car
	})
	c.OnCarSetupPacket(func(packet *packets.PacketCarSetupData) {
		car := packet.Self()
		timeCost = time.Since(*startTime)
		p.PacketCarSetupData = &car
	})
	c.OnCarStatusPacket(func(packet *packets.PacketCarStatusData) {
		car := packet.Self()
		timeCost = time.Since(*startTime)
		p.PacketCarStatusData = &car
		// metrics
		fuelMetric.WithLabelValues("FuelRemainingLaps").Set(float64(car.FuelRemainingLaps))
		fuelMetric.WithLabelValues("FuelInTank").Set(float64(car.FuelInTank))

		ERSMetric.WithLabelValues("ERSStoreEnergy").Set(float64(car.ERSStoreEnergy))
		ERSMetric.WithLabelValues("ERSHarvestedThisLapMGUK").Set(float64(car.ERSHarvestedThisLapMGUK))
		ERSMetric.WithLabelValues("ERSHarvestedThisLapMGUH").Set(float64(car.ERSHarvestedThisLapMGUH))
		ERSMetric.WithLabelValues("ERSDeployedThisLap").Set(float64(car.ERSDeployedThisLap))

		tyresAgeLapsMetric.Set(float64(car.TyresAgeLaps))

	})
	c.OnFinalClassificationPacket(func(packet *packets.PacketFinalClassificationData) {
		car := packet.Self()
		timeCost = time.Since(*startTime)
		p.PacketFinalClassificationData = &car
	})
	c.OnLobbyInfoPacket(func(packet *packets.PacketLobbyInfoData) {
		car := packet.Self()
		timeCost = time.Since(*startTime)
		p.PacketLobbyInfoData = &car
	})
	c.OnCarDamagePacket(func(packet *packets.PacketCarDamageData) {
		car := packet.Self()
		timeCost = time.Since(*startTime)
		p.PacketCarDamageData = &car
		// metrics
		DamageMetric.WithLabelValues("FrontLeftWing").Set(float64(car.FrontLeftWingDamage))
		DamageMetric.WithLabelValues("FrontRightWing").Set(float64(car.FrontRightWingDamage))
		DamageMetric.WithLabelValues("RearWing").Set(float64(car.RearWingDamage))
		DamageMetric.WithLabelValues("Floor").Set(float64(car.FloorDamage))
		DamageMetric.WithLabelValues("Diffuser").Set(float64(car.DiffuserDamage))
		DamageMetric.WithLabelValues("GearBox").Set(float64(car.GearBoxDamage))
		DamageMetric.WithLabelValues("Engine").Set(float64(car.EngineDamage))
		DamageMetric.WithLabelValues("EngineMGUKWear").Set(float64(car.EngineMGUKWear))
		for i, tyre := range wheelOrderArr {
			tyreWearMetric.WithLabelValues(tyre).Set(float64(car.TyresWear[i]))
		}
	})
	c.OnSessionHistoryPacket(func(packet *packets.PacketSessionHistoryData) {
		timeCost = time.Since(*startTime)
		p.PacketSessionHistoryData = packet
	})

	c.OnCarTelemetryPacket(func(packet *packets.PacketCarTelemetryData) {
		car := packet.Self()
		timeCost = time.Since(*startTime)
		p.PacketCarTelemetryData = &car
		// metrics
		speedMetric.Set(float64(car.Speed))
		engineRPMMetric.Set(float64(car.EngineRPM))
		throttleMetric.Set(float64(car.Throttle))
		brakeMetric.Set(float64(car.Brake))
		gearMetric.Set(float64(car.Gear))

		for i, brake := range wheelOrderArr {
			brakesTempMetric.WithLabelValues(brake).Set(float64(car.BrakesTemperature[i]))
		}
	})
	// start
	c.Run()
}
