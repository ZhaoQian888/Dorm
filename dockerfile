FROM golang:1.13-alpine as build

ENV GO111MODULE=on
ENV GOPROXY=https://goproxy.io

WORKDIR /app
COPY go.mod .
COPY go.sum .
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o dorm

########################################

FROM alpine:3.7 as prod

RUN apk update && apk add --no-cache curl

ENV TZ=Asia/Shanghai
COPY --from=build /app/dorm /usr/bin/dorm
RUN chmod +x /usr/bin/dorm \
&& ln -snf /usr/share/zoneinfo/$TZ /etc/localtime && echo $TZ > /etc/timezone

ENTRYPOINT ["dorm"]