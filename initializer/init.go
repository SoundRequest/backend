package initializer

import (
	"flag"
	"fmt"
	"log"
	"math/rand"
	"time"

	"github.com/SoundRequest/backend/db"
	"github.com/SoundRequest/backend/helper"
	"github.com/SoundRequest/backend/routes"
	"github.com/SoundRequest/backend/structure"
	"github.com/gin-gonic/gin"
)

// InitAndCheckArgs check args
func InitAndCheckArgs() {
	DEBUG := flag.Bool("DEBUG", false, "Run as DEBUG mode")
	PORT := flag.String("PORT", "9096", "Set Server's Port")

	flag.Parse()
	go initDB()

	if *DEBUG {
		initAndRunServer(*PORT, gin.DebugMode)
	} else {
		initAndRunServer(*PORT, gin.ReleaseMode)
	}
	log.Println("Server is ready.")
}

func initAndRunServer(runPort string, runMode string) {
	rand.Seed(time.Now().Unix())
	gin.SetMode(runMode)
	r := gin.Default()
	r.LoadHTMLGlob("templates/verify.html")
	routes.Auth(r)
	r.Run(":" + runPort) // listen and serve on 0.0.0.0:9096 (for windows "localhost:9096")
}

// initDB Only work for AutoMigrate
func initDB() {
	config, errLoadConfig := helper.LoadConfig("./config.json")
	if errLoadConfig != nil {
		log.Fatal(errLoadConfig)
	}
	connectionInfo := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True", config.Username, config.Password, config.Host, config.Port, config.Schema)
	dbConnection, errInitDB := db.InitDB(config.DBType, connectionInfo)
	if errInitDB != nil {
		log.Fatal(errInitDB)
	}

	fmt.Println("Performing AutoMigrate...")
	var models = []interface{}{&structure.User{}}
	dbConnection.AutoMigrate(models...)
	fmt.Println("Successfully performed AutoMigrate")

}
