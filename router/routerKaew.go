package router

import (
	"bytes"
	"context"
	"fmt"
	"group-project/firebase"
	"group-project/middleware"
	"group-project/model"
	"group-project/mongo"
	"io"
	"log"
	"mime/multipart"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/gridfs"
)

func RouterKaew(r *gin.Engine) {
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
			request.Phone,
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

	//GET ข้อมูลส่วนตัว
	r.Use(middleware.AuthMiddleware)
	r.GET("/register/", func(c *gin.Context) {
		email := c.MustGet("email").(string)
		var response model.RegisterResponseModel
		registers, err := mongo.GetRegister(email)
		if err != nil {
			fmt.Println(err)
			response.Message = "Internal Server Error"
			c.JSON(http.StatusOK, response)
		}
		response.Registers = registers
		response.Message = "Success"

		c.JSON(http.StatusOK, response)
	})

	//PUT ข้อมูลส่วนตัว
	r.PUT("/register", func(c *gin.Context) {
		email := c.MustGet("email").(string)
		var request model.CreateRegisterRequest
		err := c.ShouldBindJSON(&request)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": err.Error(),
			})
			return
		}

		err = mongo.UpdateRegister(email, request)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": err.Error(),
			})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"message": "Registers is update.",
		})
	})

	//DEL ข้อมูลส่วนตัว
	r.DELETE("/register", func(c *gin.Context) {
		email := c.MustGet("email").(string)

		err := mongo.DeleteRegister(email)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": err.Error(),
			})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"message": "Registers is Delete",
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

	// POST file
	r.POST("/file", func(c *gin.Context) {
		email := c.PostForm("email")
		fmt.Println(email)
		filename := time.Now().Format("2006-01-02T150405") + ".pdf"
		file, _ := c.FormFile("file")
		err := c.SaveUploadedFile(file, "temp/"+filename)
		if err != nil {
			fmt.Println(err)
			c.JSON(500, gin.H{
				"message": "save file fail",
			})
		}
		readFile, _ := os.ReadFile("temp/" + filename)

		conn := mongo.InitiateMongoClient()
		defer conn.Disconnect(context.Background())
		bucket, err := gridfs.NewBucket(
			conn.Database("myfiles"),
		)
		if err != nil {
			log.Fatal(err)
			os.Exit(1)
		}
		uploadStream, err := bucket.OpenUploadStream(
			filename, // this is the name of the file which will be saved in the database
		)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		// ioutil.ReadAll()
		fileSize, err := uploadStream.Write(readFile)
		if err != nil {
			log.Fatal(err)
			os.Exit(1)
		}
		log.Printf("Write file to DB was successful. File size: %d M\n", fileSize)
		uploadStream.Close()

		err = mongo.AddFileID(email, GetfileID2(filename))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": err.Error(),
			})
			return
		}
		os.Remove("temp/" + filename)

		c.JSON(http.StatusOK, gin.H{
			"message": "FileID is updated.",
		})
	})

	// get file id
	r.GET("/filename/:filename", func(c *gin.Context) {
		filename := c.Param("filename")
		var response model.FileIDResponseModel
		fileids, err := mongo.GetFileID(filename)
		if err != nil {
			fmt.Println(err)
			response.Message = "Internal Server Error"
			c.JSON(http.StatusOK, response)
		}
		response.Fileids = fileids
		response.Message = "Success"
		fmt.Println("%#v\n", response)

		c.JSON(http.StatusOK, response)
	})

	r.POST("/sendCV", func(c *gin.Context) {
		email := c.MustGet("email").(string)
		user, err := mongo.GetRegister(email)
		if err != nil {
			fmt.Println(err)
			c.JSON(500, gin.H{
				"Message": "cannot find email",
			})
		}
		conn := mongo.InitiateMongoClient()
		defer conn.Disconnect(context.Background())
		bucket, err := gridfs.NewBucket(
			conn.Database("myfiles"),
		)
		if err != nil {
			log.Fatal(err)
			c.JSON(500, gin.H{
				"message": err.Error(),
			})
		}

		// fmt.Println(user[0]["fileid"])
		// fileDoc, _ := mongo.GetFileFromFileId(fmt.Sprintf("%v", user[0]["fileid"]))
		// fmt.Println(fileDoc)
		objectId, _ := primitive.ObjectIDFromHex(fmt.Sprintf("%v", user[0]["fileid"]))

		downloadStream, err := bucket.OpenDownloadStream(objectId)
		if err != nil {
			c.JSON(500, gin.H{
				"message": err.Error(),
			})
		}
		defer func() {
			if err := downloadStream.Close(); err != nil {
				log.Fatal(err)
			}
		}()
		// fileBuffer := bytes.NewBuffer(nil)
		// if _, err := io.Copy(fileBuffer, downloadStream); err != nil {
		// 	log.Fatal(err)
		// }
		// fmt.Println(string(fileBuffer.Bytes()))
		var client = &http.Client{}

		// New multipart writer.
		body := &bytes.Buffer{}
		writer := multipart.NewWriter(body)
		fw, err := writer.CreateFormField("email")
		if err != nil {
		}
		_, err = io.Copy(fw, strings.NewReader(fmt.Sprintf("%v", user[0]["email"])))
		if err != nil {
		}
		fw, err = writer.CreateFormFile("file", fmt.Sprintf("%v-%v", user[0]["firstname"], user[0]["lastname"])+"-cv.pdf")
		if err != nil {
		}
		_, err = io.Copy(fw, downloadStream)
		if err != nil {
		}
		// Close multipart writer.
		writer.Close()

		req, err := http.NewRequest("POST", "http://localhost:8081/apply/cv", bytes.NewReader(body.Bytes()))
		if err != nil {
			fmt.Println("ERROR", err)
		}
		req.Header.Set("Content-Type", writer.FormDataContentType())
		rsp, err := client.Do(req)
		if err != nil {
			fmt.Println("send error:", err)
			c.JSON(500, gin.H{
				"message": "send error",
			})
		}
		if rsp.StatusCode != http.StatusOK {
			log.Printf("Request failed with response code: %d", rsp.StatusCode)
		}
		c.JSON(200, gin.H{
			"message": "sent success",
		})
	})
}

func GetfileID2(filename string) string {
	fileids, err := mongo.GetFileID(filename)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("%v\n", fileids)
	stringObjectID := fileids[0]["_id"].(primitive.ObjectID).Hex()
	return stringObjectID
}
