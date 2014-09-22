package routers

import (
	"github.com/astaxie/beego"
	"github.com/piotrkowalczuk/gonalytics-tracker/controllers"
)

func init() {
	beego.Router("/visits", &controllers.VisitsController{})
	beego.Router("/visits/live", &controllers.VisitsLiveController{})
	beego.Router("/visits/count", &controllers.VisitsCountController{})
	beego.Router("/visits/average-duration", &controllers.VisitsAverageTimeController{})
	beego.Router("/visits/distribution", &controllers.VisitsDistributionController{})
}
