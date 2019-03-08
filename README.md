# GoBirthday

![Birthday](https://github.com/fallais/gobirthday/blob/master/birthday.png)

**GoBirthday** is a tool written in **Go** that reminds you all birthdays that you need to wish !

## Why ?

You do not have a Facebook account ? You do not have a smartphone ? Then you do not wish any birthday except your parents, do you ?  
This software is for you.

## Usage

### As a software

Software can be used as follow : `gobirthday --contacts_file /app/contacts.json --providers_file /app/providers.json`

### As a Docker container

It can (or must) be deployed in a Docker container as follow :

`docker run -d --name gobirthday -v contacts.json:/app/contacts.json -v providers.json:/app/providers.json fallais/gobirthday`
