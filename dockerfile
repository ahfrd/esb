FROM golang:1.19 as builder

WORKDIR /go/src/github.com/ahfrd/esb-assesment/
ADD . .

RUN go mod tidy
RUN go mod vendor
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .

FROM alpine:latest

ENV TZ=Asia/Jakarta
RUN ln -snf /usr/share/zoneinfo/$TZ /etc/localtime && echo $TZ > /etc/timezone
RUN apk add --no-cache tzdata

WORKDIR /app/
RUN mkdir logs
RUN mkdir config

COPY --from=builder /go/src/github.com/ahfrd/esb-assesment/main .

RUN chmod -R 777 /app
RUN chmod -R 755 config

EXPOSE 9018

CMD ["./main"]