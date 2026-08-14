package main

import (
	"fmt"
	"log"
	"notes-api/internal/config"
	"notes-api/internal/db"
	"notes-api/internal/server"

	"github.com/gin-gonic/gin"
)



func main(){
	// load config
	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("ENV config load failed: %v", err)
	}

	// db connection
 	client, _, err := db.Connect(cfg)
	if err != nil {
		log.Fatalf("DB connect failed: %v", err)
	}

	defer func(){
		if err:=db.Disconnect(client); err != nil {
			log.Fatalf("DB disconnect failed: %v", err)
		}
	}()

	if cfg.GinEnv == "production" {
		gin.SetMode(gin.ReleaseMode)
	}

	router:=server.NewRouter()

	err = router.Run(":" + cfg.ServerPort)

	fmt.Println("Server started on port " + cfg.ServerPort)
	if err != nil {
			log.Fatalf("Server crashed: %v", err)
		}
}
