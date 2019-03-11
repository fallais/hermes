# GoBirthday

![Birthday](https://github.com/fallais/gobirthday/blob/master/birthday.png)

**GoBirthday** is a tool written in **Go** that reminds you all birthdays that you need to wish !

## Why

You do not have a Facebook account ? You do not have a smartphone ? So you do not wish any more birthday, except for your parents, do you ? This software is for you !

### What about the leap years

Hum, you have a friend who was born the 29 of February, that is sad, because every four years, no birthday for your friend. Do not worry, if you want to, you will be noticed the 1st of March : `HANDLE_LEAP_YEARS=true`.

## Features

### Contacts

A **contact** is defined by a *firstname*, a *lastname*, a *nickname*, a *description* and a *birthdate (DD/MM/YYYY or DD/MM)*. Contacts list must be as follow :

```json
[
  {
    "firstname": "Daniel",
    "lastname": "Doe",
    "birthdate": "08/03"
  },
  {
    "firstname": "Henry",
    "lastname": "Doe",
    "birthdate": "31/01/1951",
    "nickname": "Johnny",
    "description": "Best friend"
  },
  {
    "firstname": "John",
    "lastname": "Doe",
    "birthdate": "08/04/1951",
    "nickname": "Johnny"
  }
]
```

### Providers

A **provider** is used to send notifications, it could be one of the following :

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

The list of providers must be as follow :

```json
[
  {
    "type": "sms",
    "vendor": "free",
    "settings": {
      "user": "xxxxxxxx",
      "pass": "xxxxxxx"
    }
  },
  {
    "type": "email",
    "vendor": "email",
    "settings": {
      "host": "smtp",
      "port": 25,
      "recipient": "xxx.xxx@hotmail.com",
      "subject": "Birthay !"
    }
  }
]
```

### CRON

A **CRON expression** can be provided if you want to control the time when you receive the notification. If you need help with CRON expression : [CronTabGuru](https://crontab.guru/)

## Usage

### As a software

It can be used as follow : `gobirthday --cron_exp="30 11 * * *" --contacts_file /app/contacts.json --providers_file /app/providers.json`

### As a Docker container

It can also be deployed in a Docker container, it is only 20MB.

`docker run -d --name gobirthday -e CRON_EXP="30 11 * * *" -v contacts.json:/app/contacts.json -v providers.json:/app/providers.json fallais/gobirthday`

### With docker-compose

If you use the SMTP provider, you may want to use `docker-compose` :

```yaml
version: "3"

services:
  plex:
    image: fallais/gobirthday
    container_name: gobirthday
    restart: always
    environment:
      - HANDLE_LEAP_YEARS=true
      - RUN_ON_STARTUP=true
      - CRON_EXP=50 15 * * *
    volumes:
      - contacts.json:/app/contacts.json
      - providers.json:/app/providers.json
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