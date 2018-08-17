FROM golang:1.10.3-alpine
LABEL maintainer=9to6
LABEL email=ktk0011@gmail.com

RUN apk add --update git make

ADD . /go/src/github.com/9to6/aws-autoscaling-tester
RUN cd /go/src/github.com/9to6/aws-autoscaling-tester && make install
WORKDIR /go/bin

# RUN apk add --update \
#     python \
#     python-dev \
#     py-pip \
#     build-base
# RUN pip install awscli \
#   && rm -rf /var/cache/apk/*

EXPOSE 8080
ENTRYPOINT ["aws-autoscaling-tester"]
