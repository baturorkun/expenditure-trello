## Expenditure Form With Trello Integration

This application being developed in GoLang and gives you a chance that storing expenditure form to Trello as a card. Users don't need Trello account and don't know using Trello. The system provides a fancy HTML form to save cards to Trello.

All you have to set your trello api key and token in app.ini file. This config file is in the conf directory. At first, rename the file to app.ini.

You can generate api key and token from "https://trello.com/app-key".

Example app.ini file:

```
[app]

Title = "ExpenditureForm"

AllowIps = All
#AllowIps = 192.168.1.1, 192.168.2.1,  192.168.3.1

[trello]
AppKey = "64b1040607bad87a791b116e5280e66c"
Token = "b4e194a44be6278be797846acae003be9a6961d03da67802a3d12f3040bb17ef"
BoardID = "KE4wqorD"
ListNumber = 1

[server]
RunMode = debug
Port = :8000
ReadTimeout = 60
WriteTimeout = 60
```

###### Another important parameters:

BoardID : Every board has a ID on Trello like "KE4wqorD". You can see it on the URL when then board url is open.

> https://trello.com/b/KE4wqorD/batus.

ListNumber : Boards has lists like doing, todo, done. ListNumber is order number of the lists and starts with 0.

Port : You can change server port in server trunk. 

### Run

> go tun main.go

### Install
```
go build .
go install
```

### Running
```
./expenditure
```

### Running on Docker
```
docker build . -t expenditure

docker run -p 8000:8000 expenditure

```

### Usage

> Open " http://[your ip]:8000 " on your  browser.


