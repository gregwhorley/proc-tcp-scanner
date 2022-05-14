# proc-tcp-scanner
Reads in /proc/net/tcp and reports connection data

## Description
This program reads the /proc/net/tcp file at 10 second intervals and prints new connections to stdout
with the following format:

`2021-04-28 15:28:05: New connection: 192.0.2.56:5973 -> 10.0.0.5:80`

## Build
### Go binary
Requires go 1.18+
`go build -v ./...`

### Docker
If go isn't installed, you can build the binary inside a Docker image
`docker build -t pts:latest -f Dockerfile .`

## Test
`go test -v ./...`

## Run
Depending on how the app was [built](README.md#build)
### Go binary
`./proc-tcp-scanner`

### Docker
`docker run pts:latest`
