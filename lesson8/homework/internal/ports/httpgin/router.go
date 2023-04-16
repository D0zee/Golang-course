package httpgin

import (
	"github.com/gin-gonic/gin"

	"homework8/internal/app"
)

func adRouter(r *gin.RouterGroup, a app.App) {
	r.POST("/", createAd(a))
	r.PUT("/:id/status", changeAdStatus(a))
	r.PUT("/:id", updateAd(a))

	//r.GET()

}

func userRouter(r *gin.RouterGroup, a app.App) {
	r.POST("/", CreateUser(a))
	r.PUT("/:id/nickname", ChangeUser(a, changeNickname))
	r.PUT("/:id/email", ChangeUser(a, changeEmail))
}