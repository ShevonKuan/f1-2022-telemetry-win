# F1 Game Telemetry Client in Go [![Made With Go](https://img.shields.io/badge/Made%20with-Go-1f425f.svg?color=007EC6)](http://golang.org)

Telemetry client for F1 Game, written in Go. Supports the F1 UDP 2022 and previous formats.

![f1-telemetry-client](./dashboard.jpg)

## Features

- Event System
- Rich Env Constants
- UDP Stats (recv, err and packet per second rate),
- Vector3 support

## Install

```bash
# latest version (F1 2022)
go get -u github.com/shevonkuan/f1-telemetry-go@master

# for F1 2022
go get -u github.com/shevonkuan/f1-telemetry-go@v1.1.0

# for F1 2020
go get -u github.com/shevonkuan/f1-telemetry-go@v1.0.0
```

## Quick Start

```go
func main() {
  client, err := telemetry.NewClient()
  if err != nil {
	log.Fatal(err)
  }

  client.OnEventPacket(func(packet *packets.PacketEventData) {
  	fmt.Printf("Code: %s\n", packet.EventCodeString())
  })

  client.Run()
}
```
Or you can download the latest release file and run it.

## Docs

If you need more information on the F1 UDP specifications, see the [docs](/docs).

## Acknowledgements

This project is based on code from [f1-telemetry-go](https://github.com/anilmisirlioglu/f1-telemetry-go) by anilmisirlioglu.

Modifications and adaptations for this project by ShevonKuan.