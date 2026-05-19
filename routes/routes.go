package routes

import (
	"data_hoarder_go/controllers"
	"data_hoarder_go/middleware"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(
	router *gin.Engine,
) {

	auth := router.Group("/auth")

	{
		auth.POST(
			"/register",
			controllers.Register,
		)

		auth.POST(
			"/login",
			controllers.Login,
		)
	}

	fir := router.Group("/fir")

	fir.Use(
		middleware.AuthMiddleware(),
	)

	{
		fir.GET(
			"/",
			controllers.GetFIRs,
		)

		fir.POST(
			"/",
			controllers.CreateFIR,
		)

		fir.PUT(
			"/:id",
			controllers.UpdateFIR,
		)

		fir.DELETE(
			"/:id",
			controllers.DeleteFIR,
		)
	}
}