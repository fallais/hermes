# GoBirthday

![Birthday](https://github.com/fallais/gobirthday/blob/master/birthday.png)

**GoBirthday** is a tool written in **Go** that reminds you all birthdays that you need to wish !

## Why ?

You do not have a Facebook account ? You do not have a smartphone ? Then you do not wish any birthday except your parents, do you ?  
This software is for you.

## Features

### Contacts

A contact is defined by a firstname, a lastname and a birthdate (DD/MM/YYYY or DD/MM).  
Contacts list must be as follow :

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
		"birthdate": "31/01/1951"
	},
	{
		"firstname": "John",
		"lastname": "Doe",
		"birthdate": "08/04/1951"
	}
]
```

### Providers

Providers are used to send notifications (SMS, email, ..).  
Providers list must be as follow :

```json
[
  {
    "type": "sms",
    "vendor": "free",
    "settings": {
      "user": "xxxxxxxx",
      "pass": "xxxxxxx"
    }
  }
]
```

## Usage

### As a software

Software can be used as follow : `gobirthday --contacts_file /app/contacts.json --providers_file /app/providers.json`

### As a Docker container

It can (or must) be deployed in a Docker container as follow :

`docker run -d --name gobirthday -e CRON="30 11 * * *" -v contacts.json:/app/contacts.json -v providers.json:/app/providers.json fallais/gobirthday`
