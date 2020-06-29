package initializer

import (
	"flag"
	"log"

	"github.com/SoundRequest/backend/routes"
	"github.com/gin-gonic/gin"
)

// InitAndCheckArgs check args
func InitAndCheckArgs() {
	DEBUG := flag.Bool("DEBUG", false, "Run as DEBUG mode")
	PORT := flag.String("PORT", "9096", "Set Server's Port")

	flag.Parse()
	if *DEBUG {
		initAndRunServer(*PORT, gin.DebugMode)
	} else {
		initAndRunServer(*PORT, gin.ReleaseMode)
	}
}

func initAndRunServer(runPort string, runMode string) {
	gin.SetMode(runMode)
	r := gin.Default()
	routes.Auth(r)
	r.Run(":" + runPort) // listen and serve on 0.0.0.0:9096 (for windows "localhost:9096")
	log.Println("Server is ready.")
}
