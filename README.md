# GoBirthday

![Birthday](https://github.com/fallais/gobirthday/blob/master/assets/gobirthday.png)

**GoBirthday** is a tool written in **Go** that reminds you all birthdays that you need to wish ! It uses the [gonotify](https://github.com/fallais/gonotify) library for the notifications.

## Why

You do not have a Facebook account ? You do not have a smartphone ? So you do not wish any more birthday, except for your parents, do you ? This software is for you !

### What about the leap years

Hum, you have a friend who was born the 29 of February, that is sad, because every four years, no birthday for your friend. Do not worry, if you want to, you will be noticed the 1st of March !

## Configuration file

The configuration file must be as follow, it must be provided with `--config` flag.

```yaml
general:
  cron_exp: "0 30 14 * * *"
  run_on_startup: true
  handle_leap_years: true
  notification_template:
    header: Hey !
    base: This is the birthday of {{contact}} !
    age: "{{age}} years old ! :)"
    footer: See you !

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

providers:
  - type: "sms"
    vendor: "free"
    settings:
      user: "1234568797"
      pass: "xxxxxxx"
```

### Notification template

**notification_template** is used to configure the message you want to receive. Two variables must be provided : `{{contact}}` and `{{age}}`. The template is divided in for parts. They will be concatenated in one large message.

```yaml
notification_template:
    header: Hey !
    base: This is the birthday of {{contact}} !
    age: "{{age}} years old ! :)"
    footer: See you !
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