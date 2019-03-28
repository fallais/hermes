FROM golang:latest as builder
RUN mkdir /go/src/gobirthday
WORKDIR /go/src/gobirthday
ADD . ./
RUN go get -d -v && go build -o gobirthday

FROM frolvlad/alpine-glibc:latest
LABEL maintainer="francois.allais@hotmail.com"

ENV CRON_EXP="0 50 15 * * *"
ENV HANDLE_LEAP_YEARS=true
ENV RUN_ON_STARTUP=true

RUN mkdir /app
WORKDIR /app
COPY --from=builder /go/src/gobirthday /app/

CMD /app/gobirthday --contacts_file /app/contacts.json --providers_file /app/providers.json --cron_exp="$CRON_EXP" --handle_leap_years=$HANDLE_LEAP_YEARS --run_on_startup=$RUN_ON_STARTUP
