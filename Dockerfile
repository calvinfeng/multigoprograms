FROM golang:1.13.1-alpine as build
WORKDIR /go/src/multigoprograms
COPY . .

ENV CGO_ENABLED=0
ENV GOOS=linux
ENV GOARCH=amd64
RUN go build -a -tags netgo -ldflags '-w' -o multigoprograms .

FROM ubuntu:18.04 as deploy
RUN mkdir p /var/log/multigoprograms
WORKDIR /go/bin
COPY --from=build /go/src/multigoprograms/multigoprograms .
COPY --from=build /go/src/multigoprograms/supervisor/supervisord.conf /etc/supervisor/conf.d/supervisord.conf

RUN apt-get update && apt-get install -y supervisor

CMD /usr/bin/supervisord




