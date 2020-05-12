module github.com/calc-log/calc-log-demo

go 1.14

// +heroku goVersion go1.14.2

// +heroku get github.com/gorilla/websocket

require (
	github.com/gomodule/redigo v2.0.0+incompatible
	github.com/gorilla/websocket v1.4.0
	github.com/heroku/x v0.0.1
	github.com/pkg/errors v0.8.1
	github.com/sirupsen/logrus v1.4.2
	golang.org/x/sys v0.0.0-20190624142023-c5567b49c5d0 // indirect
)
