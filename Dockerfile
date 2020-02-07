FROM golang:rc-buster AS builder

WORKDIR /home/go/src/github.com/ken343/vnf-suite 
COPY  . .
RUN pwd
RUN ls
RUN ls proxy/

RUN go build ./cmd/mcproxy

FROM debian:latest
COPY --from=builder /home/go/src/github.com/ken343/vnf-suite/ /home/vnf/
RUN  mv /home/vnf/ /bin/
EXPOSE 80 443 22
EXPOSE 4444
RUN cd /home/vnf/