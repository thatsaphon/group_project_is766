package main

import (
	"fmt"
	"group-project/firebase"
	"group-project/router"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

//register//

func main() {
	firebaseApp, err := firebase.InitFirebase()
	if err != nil {
		fmt.Println(err)
	}
	firebase.InitClientAuth(firebaseApp)

	r := gin.Default()

	r.Use(cors.Default())
	router.RouterBoat(r)
	router.RouterKaew(r)

	r.Run()

}
