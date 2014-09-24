package main

import (
	"runtime"

	"github.com/astaxie/beego"
	_ "github.com/piotrkowalczuk/gonalytics-tracker/routers"
	"github.com/piotrkowalczuk/gonalytics-tracker/services"
)

type flags struct {
	runTracker bool
}

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())

	runTracker()

}

func runTracker() {
	services.InitLogger()
	mongoDB := services.InitMongoDB("mongodb://mongodb/gonalytics")
	cassandra := services.InitCassandra("gonalytics", []string{"127.0.0.1"})
	services.InitRepositoryManager(mongoDB, cassandra)

	defer cassandra.Close()
	beego.Run()
}
