package general

import (
	"matar/controllers"
	"matar/middlewares/authMiddleware"

	"github.com/gin-gonic/gin"
)

func AutomobileAdRoute(routerGroup *gin.RouterGroup) {
	automobileAds := routerGroup.Group("/automobile-ads")
	automobileAds.POST("/", authMiddleware.VerifyUser(), controllers.CreateAutomobileAd())
	automobileAds.GET("/", controllers.SearchAutomobileAd())
	automobileAds.GET("/:id", controllers.GetAutomobileAdById())
	automobileAds.DELETE("/:id", authMiddleware.VerifyUser(), controllers.DeleteAutomobileAdById())
	automobileAds.PUT("/:id", authMiddleware.VerifyUser(), controllers.UpdateAutomobileAdById())
}
