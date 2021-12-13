FROM golang:alpine as apibuilder

WORKDIR /go/src/app/

RUN apk add --no-cache git && \
    apk add build-base

RUN go get github.com/rs/cors && \
    go get github.com/gorilla/mux && \
    go get github.com/gorilla/handlers && \
    go get gopkg.in/natefinch/lumberjack.v2 && \
    go get github.com/sirupsen/logrus && \
    go get github.com/logmatic/logmatic-go && \
    go get github.com/simplereach/timeutils && \
    go get go.mongodb.org/mongo-driver/mongo

COPY ./*.go ./

RUN GOOS=linux go build -a -o bytevault-api .

FROM golang:alpine

RUN mkdir -p /app && mkdir -p /logs && \
    apk add --no-cache tzdata

WORKDIR /app

COPY --from=apibuilder /go/src/app/bytevault-api .

ENTRYPOINT ["./bytevault-api"]

EXPOSE 8000