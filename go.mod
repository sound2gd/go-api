module github.com/sound2gd/go-api

require (
	github.com/golang/protobuf v1.2.0
	github.com/gorilla/handlers v1.4.0
	github.com/gorilla/websocket v1.4.0
	github.com/joncalhoun/qson v0.0.0-20170526102502-8a9cab3a62b1
	github.com/micro/go-api v0.6.0
	github.com/micro/go-log v0.1.0
	github.com/micro/go-micro v0.27.0
	github.com/micro/go-rcache v0.2.1
	github.com/micro/micro v0.23.2
	github.com/micro/util v0.2.0
	github.com/pborman/uuid v1.2.0
	golang.org/x/crypto v0.0.0-20190325154230-a5d413f7728c
	gopkg.in/go-playground/validator.v9 v9.27.0
)

replace github.com/micro/go-api => ./
