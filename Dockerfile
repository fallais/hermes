FROM golang:latest as builder
RUN mkdir /go/src/gobirthday
WORKDIR /go/src/gobirthday
ADD . ./
RUN go get -d -v && go build -o gobirthday

FROM frolvlad/alpine-glibc:latest
LABEL maintainer="francois.allais@hotmail.com"

ENV CRON="50 15 * * *"

RUN mkdir /app
WORKDIR /app
COPY --from=builder /go/src/gobirthday /app/

CMD /app/gobirthday --contacts_file /app/contacts.json --providers_file /app/providers.json --cron_exp=$CRON
