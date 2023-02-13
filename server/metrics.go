package server

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

var (
	// Motion Metric
	accelerationMetric = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Name: "f1_telemetry_acceleration",
		Help: "Acceleration",
	}, []string{"axis"})
	attitudeMetric = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Name: "f1_telemetry_attitude",
		Help: "Attitude",
	}, []string{"axis"})
	// Speed Metric
	speedTrapMetric = promauto.NewGauge(prometheus.GaugeOpts{
		Name: "f1_telemetry_speed_trap",
		Help: "Speed Trap",
	})
	speedMetric = promauto.NewGauge(prometheus.GaugeOpts{
		Name: "f1_telemetry_car_speed",
		Help: "Car Speed",
	})
	// operate metric
	throttleMetric = promauto.NewGauge(prometheus.GaugeOpts{
		Name: "f1_telemetry_throttle",
		Help: "Throttle",
	})
	brakeMetric = promauto.NewGauge(prometheus.GaugeOpts{
		Name: "f1_telemetry_brake",
		Help: "Brake",
	})
	gearMetric = promauto.NewGauge(prometheus.GaugeOpts{
		Name: "f1_telemetry_gear",
		Help: "Gear",
	})
	// Engine Metrics
	engineRPMMetric = promauto.NewGauge(prometheus.GaugeOpts{
		Name: "f1_telemetry_car_engine_rpm",
		Help: "Car Engine RPM",
	})

	// Lap Time Metrics
	fastestLapMetric = promauto.NewGauge(prometheus.GaugeOpts{
		Name: "f1_telemetry_fastest_lap",
		Help: "Fastest Lap Time",
	})
	lastLapTimeMetric = promauto.NewGauge(prometheus.GaugeOpts{
		Name: "f1_telemetry_last_lap",
		Help: "Last Lap Time",
	})
	currentLapTimeMetric = promauto.NewGauge(prometheus.GaugeOpts{
		Name: "f1_telemetry_current_lap",
		Help: "Current Lap Time",
	})
	currentLapNumMetric = promauto.NewGauge(prometheus.GaugeOpts{
		Name: "f1_telemetry_current_lap_num",
		Help: "Current Lap Number",
	})

	// Car Status Metrics
	fuelMetric = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Name: "f1_telemetry_fuel",
		Help: "Fuel",
	}, []string{"fuel"})
	ERSMetric = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Name: "f1_telemetry_ers",
		Help: "ERS",
	}, []string{"ERS"})
	// Car Damage Metrics
	DamageMetric = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Name: "f1_telemetry_damage",
		Help: "Car Damage",
	}, []string{"damage"})
	// Tyre Metric
	tyresAgeLapsMetric = promauto.NewGauge(prometheus.GaugeOpts{
		Name: "f1_telemetry_tyres_age_lap",
		Help: "Tyres Age Lap",
	})
	tyreWearMetric = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Name: "f1_telemetry_tyre_wear",
		Help: "Tyres Wear",
	}, []string{"tyre"})

	// Brake Temp Metrics
	brakesTempMetric = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Name: "f1_telemetry_brake_temp",
		Help: "Brake Temperatures",
	}, []string{"brake"})
	// Session History Metrics
	//sessionHistoryMetric = promauto.NewGauge
)
