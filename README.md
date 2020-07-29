# Ichos

> Simple API and websocket in golang

## How to run

You can simply run the project using following command

```
go run main.go
```

However, it's recommended to run this project using `docker-compose` with following command

```
docker-compose -f docker-compose.yaml up --build
```

The app would available in `localhost:8000`

## Websocket

Open `localhost:8000` in your browser


## Available REST Endpoint

### POST `/message`

Payload example

```
{
    "message": "Hello World!"
}
```

Response example

```
{
    "Status": "Ok"
}
```

### GET `/message`

Response example

```
{
    "Status": "Ok",
    "Messages": [
        "hello world2",
        "hello world2",
        "hello world2"
  ]
}
```