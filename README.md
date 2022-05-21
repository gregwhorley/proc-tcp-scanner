# proc-tcp-scanner
Reads in /proc/net/tcp and reports connection data

## Description
This program reads the /proc/net/tcp file at 10 second intervals and prints new connections to stdout
with the following format:

`2021-04-28 15:28:05: New connection: 192.0.2.56:5973 -> 10.0.0.5:80`

## Build
### Go binary
Requires go 1.18+. Compiled artifact will show up in the `build/` directory.

Run `make` to build the binary.

### Docker
If go isn't installed, you can build the binary inside a Docker image. The image will automatically
be tagged as `pts:latest`.

Run `make docker-build` to create a Docker image with the compiled binary contained within.

## Test
Run `make test` to run unit and integration tests.

## Run
Depending on how the app was [built](README.md#build)
### Go binary
`./build/pts`

### Docker
`docker run pts:latest`
