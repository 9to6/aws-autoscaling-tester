FROM golang:1.10.3-alpine
LABEL maintainer=surieven
LABEL email=ktk0011+dev@gmail.com

RUN apk add --update git make

ADD . /go/src/metric-generator
RUN cd /go/src/metric-generator && make install
RUN pwd
WORKDIR /go/bin
RUN ls

EXPOSE 8080
ENTRYPOINT ["metric-generator"]
