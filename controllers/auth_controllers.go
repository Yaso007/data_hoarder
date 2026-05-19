package controllers

import (
	"data_hoarder_go/models"
	"data_hoarder_go/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Register(
	c *gin.Context,
) {

	var user models.User

	err := c.BindJSON(&user)

	if err != nil {

		c.JSON(
			http.StatusBadRequest,
			gin.H{
				"error": "invalid request",
			},
		)

		return
	}

	err = services.RegisterUser(user)

	if err != nil {

		c.JSON(
			http.StatusInternalServerError,
			gin.H{
				"error": err.Error(),
			},
		)

		return
	}

	c.JSON(
		http.StatusCreated,
		gin.H{
			"message": "user registered",
		},
	)
}

func Login(
	c *gin.Context,
) {

	var user models.User

	err := c.BindJSON(&user)

	if err != nil {

		c.JSON(
			http.StatusBadRequest,
			gin.H{
				"error": "invalid request",
			},
		)

		return
	}

	token, err :=
		services.LoginUser(user)

	if err != nil {

		c.JSON(
			http.StatusUnauthorized,
			gin.H{
				"error": err.Error(),
			},
		)

		return
	}

	c.JSON(
		http.StatusOK,
		gin.H{
			"token": token,
		},
	)
}