FROM golang:1.10.3-alpine
LABEL maintainer=surieven
LABEL email=ktk0011+dev@gmail.com

RUN apk add --update git make

ADD . /go/src/metric-generator
RUN cd /go/src/metric-generator && make install
WORKDIR /go/bin

RUN apk add --update \
    python \
    python-dev \
    py-pip \
    build-base \
  && pip install awscli \
  && rm -rf /var/cache/apk/*

EXPOSE 8080
ENTRYPOINT ["metric-generator"]
