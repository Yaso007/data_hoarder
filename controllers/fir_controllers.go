package controllers

import (
	"data_hoarder_go/models"
	"data_hoarder_go/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetFIRs(
	c *gin.Context,
) {

	firs, err :=
		services.GetAllFIRs()

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
		http.StatusOK,
		firs,
	)
}

func CreateFIR(
	c *gin.Context,
) {

	var fir models.FIR

	err := c.BindJSON(&fir)

	if err != nil {

		c.JSON(
			http.StatusBadRequest,
			gin.H{
				"error": "invalid request",
			},
		)

		return
	}

	err = services.CreateFIR(fir)

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
			"message": "FIR created",
		},
	)
}

func UpdateFIR(
	c *gin.Context,
) {

	id := c.Param("id")

	var fir models.FIR

	err := c.BindJSON(&fir)

	if err != nil {

		c.JSON(
			http.StatusBadRequest,
			gin.H{
				"error": "invalid request",
			},
		)

		return
	}

	err = services.UpdateFIR(
		id,
		fir,
	)

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
		http.StatusOK,
		gin.H{
			"message": "FIR updated",
		},
	)
}

func DeleteFIR(
	c *gin.Context,
) {

	id := c.Param("id")

	err := services.DeleteFIR(id)

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
		http.StatusOK,
		gin.H{
			"message": "FIR deleted",
		},
	)
}