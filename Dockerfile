FROM golang:latest

WORKDIR /go/src/restwithgo

# RUN go get -u github.com/mitranim/gow
RUN go get github.com/codegangsta/gin
RUN go get -u github.com/golang/dep/cmd/dep
COPY ./Gopkg.toml ./
COPY ./Gopkg.lock ./

RUN dep ensure -vendor-only

COPY ./ ./
CMD ["gin"]




