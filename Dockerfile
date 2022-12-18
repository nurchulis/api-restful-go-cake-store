FROM golang:1.16

WORKDIR /go/src/api-restfull-go-cake-store
COPY . .
RUN go get -d -v ./...
RUN go install -v ./...

COPY go.mod .
COPY go.sum .
COPY . .

RUN go build -o api-restfull-go-cake-store

EXPOSE 7000

EXPOSE 3306 33060

ENTRYPOINT ./api-restfull-go-cake-store