### Go [here](docs/Q_and_A.md) for answers to the list of questions.

# proc-tcp-scanner
Reads in /proc/net/tcp and reports connection data such as:
* Print new connections from remote IPv4 addresses and ports
* Print possible port scans when detected
* Expose connection count as a Prometheus metric

## Description
This program reads the /proc/net/tcp file at 10 second intervals and prints new connections to stdout
with the following format:

`2021-04-28 15:28:05: New connection: 192.0.2.56:5973 -> 10.0.0.5:80`

If port scans are detected, a message will be printed out with the following format:

`2021-04-28 15:28:05: Port scan detected: 192.0.2.56 -> 10.0.0.5 on ports 80,81,82,83`

The app listens on port 1234 and serves up a Prometheus counter metric for new connections.

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
While this app will run in a container, some extra command line flags will be required if the intent
is to print out connection data from the host.
`docker run --privileged --mount type=bind,src=/proc,target=/proc pts:latest`
