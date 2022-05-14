# proc-tcp-scanner
Reads in /proc/net/tcp and reports connection data

## Description
This program reads the /proc/net/tcp file at 10 second intervals and prints new connections to stdout
with the following format:

`2021-04-28 15:28:05: New connection: 192.0.2.56:5973 -> 10.0.0.5:80`

## Build
### Go binary
Requires go 1.18+
`go build`

### Docker
If go isn't installed, build and run in a Docker container
`docker run -f Dockerfile .`
