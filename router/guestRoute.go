package router

import (
	"fmt"
	"group-project/firebase"
	"group-project/model"
	"group-project/mongo"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GuestRoute(r *gin.Engine) {
	//search job from position
	r.GET("/position/:position", func(c *gin.Context) {
		position := c.Param("position")
		var response model.JobResponseModel
		jobs, err := mongo.GetPosition(position)
		if err != nil {
			fmt.Println(err)
			response.Message = "Internal Server Error"
			c.JSON(http.StatusOK, response)
		}
		response.Jobs = jobs
		response.Message = "Success"

		c.JSON(http.StatusOK, response)
	})

	//แสดง ตำแหน่งงานทั้งหมด ที่ user ทุกคน สมัครงาน หรือ ถูกใจไว้
	r.GET("/alljob", func(c *gin.Context) {
		var response model.JobsResponseModel
		Jobs, err := mongo.GetJobs()
		if err != nil {
			fmt.Println(err)
			response.Message = "Internal Server Error"
			c.JSON(http.StatusOK, response)
		}
		response.Jobs = Jobs
		response.Message = "Success"

		c.JSON(http.StatusOK, response)
	})

	//POST ข้อมูลส่วนตัว
	r.POST("/register", func(c *gin.Context) {
		var request model.CreateRegisterRequest
		err := c.ShouldBindJSON(&request)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": err.Error(),
			})
			return
		}
		doesEmailExist, err := firebase.GetUserByEmail(request.Email)
		if doesEmailExist != nil {
			doesEmailExist, _ := mongo.GetRegister(request.Email)
			if doesEmailExist != nil {
				c.JSON(400, gin.H{
					"message": "this email is already used",
				})
				return
			}
			err = mongo.CreateRegister(request)
			if err != nil {
				c.JSON(http.StatusBadRequest, gin.H{
					"message": err.Error(),
				})
				return
			}
			if err == nil {
				c.JSON(http.StatusOK, gin.H{
					"message": "Register is created.",
				})
			}
			return
		}

		err = firebase.CreateUser(
			request.Email,
			request.Password,
			request.Firstname+" "+request.Lastname,
		)
		if err != nil && err.Error() != "EOF" {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": err.Error(),
			})
			return
		}

		err = mongo.CreateRegister(request)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": err.Error(),
			})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"message": "Register is created.",
		})
	})

	//search job from company
	r.GET("/company/:company", func(c *gin.Context) {
		company := c.Param("company")
		var response model.JobResponseModel
		jobs, err := mongo.GetCompany(company)
		if err != nil {
			fmt.Println(err)
			response.Message = "Internal Server Error"
			c.JSON(http.StatusOK, response)
		}
		response.Jobs = jobs
		response.Message = "Success"

		c.JSON(http.StatusOK, response)
	})

	//search job from location
	r.GET("/location/:location", func(c *gin.Context) {
		location := c.Param("location")
		var response model.JobResponseModel
		jobs, err := mongo.GetLocation(location)
		if err != nil {
			fmt.Println(err)
			response.Message = "Internal Server Error"
			c.JSON(http.StatusOK, response)
		}
		response.Jobs = jobs
		response.Message = "Success"

		c.JSON(http.StatusOK, response)
	})

}
