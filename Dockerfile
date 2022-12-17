 
FROM golang:1.19

RUN go install github.com/beego/bee/v2@latest
RUN go install github.com/go-sql-driver/mysql@v1.7.0 
RUN go install github.com/gorilla/mux@v1.8.0 
RUN go install github.com/joho/godotenv@v1.4.0

ENV username=root
ENV password=tika
ENV db_name=cakes

EXPOSE 8010
CMD ["bee", "run"]