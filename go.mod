module github.com/shevonkuan/f1-telemetry-go

go 1.16

replace github.com/shevonkuan/f1-telemetry-go => ./

require (
	github.com/gin-gonic/gin v1.8.2
	github.com/gorilla/websocket v1.5.0
	github.com/prometheus/client_golang v1.14.0
)
