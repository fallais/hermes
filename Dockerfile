FROM frolvlad/alpine-glibc:latest
LABEL maintainer="francois.allais@hotmail.com"

ENV CRON="50 15 * * *"

RUN mkdir /app
ADD gobirthday /app

CMD /app/gobirthday --contacts_file /app/contacts.json --providers_file /app/providers.json --cron_exp=$CRON
