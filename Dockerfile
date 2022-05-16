FROM golang:latest as builder
WORKDIR /go/src/gobirthday
ADD . /go/src/gobirthday
RUN go get -d -v ./...
RUN go build -o /go/bin/gobirthday

FROM gcr.io/distroless/base-debian10
COPY --from=builder /go/bin/gobirthday /
CMD [ "/gobirthday", "--config", "/config.yaml" ]