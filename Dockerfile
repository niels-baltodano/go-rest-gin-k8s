FROM golang:1.13-buster as build

RUN curl https://glide.sh/get | sh

ENV PKG_NAME=go-rest-gin-k8s
ENV PKG_PATH=$GOPATH/src/$PKG_NAME
WORKDIR $PKG_PATH

COPY glide.yaml glide.lock $PKG_PATH/
RUN glide install

COPY . $PKG_PATH
#RUN ls -l
RUN go build main.go

FROM gcr.io/distroless/base
COPY --from=build /go/src/go-rest-gin-k8s/main /
CMD ["/main"]