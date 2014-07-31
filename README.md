Gonalytics (tracker)
=============

Gonalytic is a open source web analytics application. It consists of:
- tracker (this repository)
- tracking script: https://github.com/piotrkowalczuk/gonalytics-tracking-script
- dashboard: https://github.com/piotrkowalczuk/gonalytics-dashboard

Instalation
------------
1. Set you GOPATH properly (http://golang.org/doc/code.html#GOPATH)
2. `go get github.com/piotrkowalczuk/gonalytics-tracker`
3. `go get` if some dependencies are missing
4. Create `conf/app.conf` based on `conf/app.conf.dist`

Commands
--------

#### Build

    go build app.go


Dependencies
------------
- MongoDB