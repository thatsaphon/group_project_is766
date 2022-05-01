package router

import (
	"fmt"
	"group-project/middleware"
	"group-project/model"
	"group-project/mongo"
	"net/http"

	"github.com/gin-gonic/gin"
)

func RouterBoat(r *gin.Engine) {
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
	r.GET("/alluserjob", func(c *gin.Context) {
		var response model.JobsResponseModel
		Jobs, err := mongo.GetJobs()
		if err != nil {
			fmt.Println(err)
			response.Message = "Internal Server Error"
			c.JSON(http.StatusOK, response)
		}
		response.Position = Jobs
		//response.Message = "Success"

		c.JSON(http.StatusOK, response)
	})
	r.Use(middleware.AuthMiddleware)

	//แสดง ตำแหน่งงานทั้งหมดของแต่ละ User โดยหาตาม email
	r.GET("/userjob/", func(c *gin.Context) {
		email := c.MustGet("email").(string)
		var response model.UserResponseJobModel
		useremails, err := mongo.GetUserJob(email)
		if err != nil {
			fmt.Println(err)
			response.Message = "Internal Server Error"
			c.JSON(http.StatusOK, response)
		}
		response.Userjobs = useremails
		//response.Message = "Success"

		c.JSON(http.StatusOK, response)
	})

	//insert job ที่เลือกลง DB
	r.POST("/userjob", func(c *gin.Context) {
		var request model.CreateJobRequest
		err := c.ShouldBindJSON(&request)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": err.Error(),
			})
			return
		}

		// ตรวจสอบ URL ถูกต้องหรือไม่ ก่อน บันทึกลงฐานข้อมูล Recheck
		resp, err := http.Get(request.Urllink)
		if err == nil {
			fmt.Println(resp.Status)
			defer resp.Body.Close()
			err = mongo.CreateJob(request)
			if err != nil {
				c.JSON(http.StatusBadRequest, gin.H{
					"message": err.Error(),
				})
				return
			}

			c.JSON(http.StatusOK, gin.H{
				"message": "Job is created.",
			})

		} else {
			c.JSON(http.StatusOK, gin.H{
				"message": "Job is expire",
			})
		}

	})

	//update status ของ Job โดยส่งใช้ Email link เป็นเงื่อนไข และเปลี่ยนเฉพาะ status ไป
	//status = Register , Delete , Like
	r.PUT("/userjob/", func(c *gin.Context) {
		upemail := c.MustGet("email").(string)

		var request model.CreateJobRequest
		err := c.ShouldBindJSON(&request)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": err.Error(),
			})
			return
		}

		// ตรวจสอบ URL ถูกต้องหรือไม่ ก่อน บันทึกลงฐานข้อมูล Recheck
		resp, err := http.Get(request.Urllink)
		if err == nil {
			fmt.Println(resp.Status)
			defer resp.Body.Close()
			err = mongo.UpdateJob(request.Status, request.Urllink, upemail, request)
			if err != nil {
				c.JSON(http.StatusBadRequest, gin.H{
					"message": err.Error(),
				})
				return
			}

			c.JSON(http.StatusOK, gin.H{
				"message": "Job Status is update.",
			})

		} else {
			c.JSON(http.StatusOK, gin.H{
				"message": "Job is expire",
			})
		}

	})

	//ลบ Job ที่ ไม่ต้องการออกไปโดยใช้ ID
	r.DELETE("/userjob/:id", func(c *gin.Context) {
		id := c.Param("id")

		err := mongo.DeleteJob(id)
		println(id)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": err.Error(),
			})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"message": "Job is Deleted.",
		})
	})

	//แสดงตำแหน่งงานทั้งหมดที่ Scraping มาเก็บไว้แล้ว ใน DB Table : Jobtest2
	r.GET("/jobbyPosition", func(c *gin.Context) {
		var response model.JobsResponseModel
		Jobs, err := mongo.GetPosition2()
		if err != nil {
			fmt.Println(err)
			response.Message = "Internal Server Error"
			c.JSON(http.StatusOK, response)
		}
		response.Position = Jobs
		//response.Message = "Success"

		c.JSON(http.StatusOK, response)
	})

}
