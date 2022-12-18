 
FROM golang:1.16

RUN go get github.com/beego/bee/v2@latest
RUN go get github.com/go-sql-driver/mysql@v1.7.0 
RUN go get github.com/gorilla/mux@v1.8.0 
RUN go get github.com/joho/godotenv@v1.4.0

FROM golang:1.9.2 
ADD . /go/src/api-restfull-go-cake-store
WORKDIR /go/src/api-restfull-go-cake-store
RUN go get api-restfull-go-cake-store
RUN go install
ENTRYPOINT ["/go/bin/api-restfull-go-cake-store"]

ENV username=root
ENV password=tika
ENV db_name=cakes

ENV GO111MODULE=off
COPY go.mod .
COPY go.sum .
COPY . .

EXPOSE 8010
CMD ["bee", "run"]
