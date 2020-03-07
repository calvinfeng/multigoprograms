FROM golang:1.13.1-alpine as build
WORKDIR /go/src/multigoprograms
COPY . .

ENV CGO_ENABLED=0
ENV GOOS=linux
ENV GOARCH=amd64
RUN go build -a -tags netgo -ldflags '-w' -o multigoprograms .

FROM ubuntu:18.04 as deploy
RUN mkdir p /var/log/multigoprograms
COPY --from=build /go/src/multigoprograms/supervisor/supervisord.conf /etc/supervisor/conf.d/supervisord.conf

WORKDIR /go/bin
COPY --from=build /go/src/multigoprograms/conf ./conf
COPY --from=build /go/src/multigoprograms/launch_client.sh .
COPY --from=build /go/src/multigoprograms/multigoprograms .

RUN apt-get update
RUN apt-get install -y supervisor
RUN apt-get install -y curl net-tools
RUN chmod +x /go/bin/launch_client.sh

ENV SERVER_PORT 8081

CMD /usr/bin/supervisord




