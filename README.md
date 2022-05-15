[![Go Report Card](https://goreportcard.com/badge/github.com/fallais/hermes)](https://goreportcard.com/report/github.com/fallais/hermes)
[![made-with-Go](https://img.shields.io/badge/Made%20with-Go-1f425f.svg)](http://golang.org)

# Hermes

![Hermes](https://github.com/fallais/gobirthday/blob/master/assets/logo.png)

**Hermes** is a tool written in **Go** that reminds you things that you have to do in your day-to-day life !

> Disclaimer : the project has been **renamed** ! It was previously **GoBirthday** but the name was not fully accurate because the aim of the tool is to help you being notify of everything, not only birthday.

If you answer *yes* to one of theses questions, you need this tool:

- You do not have a Facebook account ? So you do not wish any more birthday, except for your parents, do you ?  
- You do not have a smartphone ? So you can't use the Google Agenda to reminds you things you have to do, can you ?

## Concept

The two things that you can be notified about are :
 - Birthday
 - Task

The **birthday** will occur once per year, and a **task** is something simple that you have to remind, for example, taking out the trash every week.

### What about the leap years

Hum, you have a friend who was born the 29 of February, that is sad, because you wish it only once every four years. Do not worry, if you want to, you will be noticed the **28th of February** or **1st of March** (depends on how superstitious you are) !

## Configuration file

The configuration file must be as follow, it must be provided with `--config` flag.

```yaml
general:
  run_on_startup: true
  leap_years:
    is_enabled: true
    mode: before

contacts:
  - firstname: "Daniel"
    lastname: "Doe"
    birthdate: "27/05"
  - firstname: "Henry"
    lastname: "Doe"
    birthdate: "31/01/1951"
  - firstname: "John"
    lastname: "Doe"
    birthdate: "08/04/1951"

things:
  - name: "Sortir les poubelles"
    when: "30 19 * * WED"

providers:
  - type: "sms"
    vendor: "free"
    settings:
      user: "1234568797"
      pass: "xxxxxxx"
```

### Contacts

A **contact** is defined by :

- a *firstname* : mandatory
- a *lastname*
- a *nickname*
- a *description*
- a *birthdate (DD/MM/YYYY or DD/MM)* : mandatory

If you do not provide the year in the birthdate, you will not know the age of your contact, but only that it is its birthday.

### Notifiers

A **notifier** is used to send notifications, it could be one of the following :

- SMS
  - Free
  - Orange *(not yet)*
  - SFR *(not yet)*
- Email
- Webhook
  - Slack *(not yet)*
  - Mattermost *(not yet)*
  - IFTTT *(not yet)*
- etc..

### CRON

A **CRON expression** must be provided if you want to control the time when you receive the notification. If you need help with CRON expression : [CronTabGuru](https://crontab.guru/)

> **Attention** : a **second** must must added before the CRON expression because of the library used (`github.com/robfig/cron`). For example : `0 50 15 * * *`

## Usage

### As a software

It can be used as follow : `gobirthday --config config.yaml`

### As a Docker container

It can also be deployed in a Docker container, it is only 20MB.

`docker run -d --name gobirthday --volume config.yaml:/config.yaml fallais/gobirthday`

### With docker-compose

If you use the SMTP provider, you may want to use `docker-compose` :

```yaml
version: "3"

services:
  gobirthday:
    image: fallais/gobirthday
    container_name: gobirthday
    restart: always
    volumes:
      - config.yaml:/config.yaml
    networks:
      main:
        aliases:
          - gobirthday
  
  smtp:
    image: namshi/smtp
    container_name: smtp
    restart: always
    networks:
      main:
        aliases:
          - smtp
```

## Contributing

1. Fork it!
2. Create your feature branch: `git checkout -b my-new-feature`
3. Commit your changes: `git commit -am 'Add some feature'`
4. Push to the branch: `git push origin my-new-feature`
5. Submit a pull request

## Credits

Implemented by François Allais

## License

See `LICENSE`.