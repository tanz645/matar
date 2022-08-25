package general

import (
	automobileAdGeneralcontroller "matar/controllers/automobileAdController"
	"matar/middlewares/authMiddleware"

	"github.com/gin-gonic/gin"
)

func AutomobileAdRoute(routerGroup *gin.RouterGroup) {
	automobileAds := routerGroup.Group("/automobile-ads")
	automobileAds.GET("/:id/by-vendor", authMiddleware.VerifyUser(), automobileAdGeneralcontroller.GetAutomobileAdByUserId())
	automobileAds.GET("/by-vendor", authMiddleware.VerifyUser(), automobileAdGeneralcontroller.GetAutomobileAdsByUserId())
	automobileAds.POST("/", authMiddleware.VerifyUser(), automobileAdGeneralcontroller.CreateAutomobileAd())
	automobileAds.GET("/", automobileAdGeneralcontroller.SearchAutomobileAd())
	automobileAds.GET("/:id", automobileAdGeneralcontroller.GetAutomobileAdById())
	automobileAds.DELETE("/:id", authMiddleware.VerifyUser(), automobileAdGeneralcontroller.DeleteAutomobileAdById())
	automobileAds.PUT("/:id", authMiddleware.VerifyUser(), automobileAdGeneralcontroller.UpdateAutomobileAdById())

}
