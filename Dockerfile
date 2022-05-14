FROM golang:bullseye

COPY *.go go.mod /src/
WORKDIR /src
RUN go build -o /pts ./...
WORKDIR /
CMD ["./pts"]
