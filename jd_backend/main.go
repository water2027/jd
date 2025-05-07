package main

import (
	_ "jd/database"
	"jd/config"
	"jd/route"
)

func main() {
	config.InitConfig()

	r := route.SetupRouter()

	r.Run(":8080")
}
