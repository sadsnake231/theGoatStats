package main

import(
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"log"
	"fmt"
	"os"
	"theGoatStats/routes"
	"theGoatStats/scraper"
)

func Routes(incomingRoutes *gin.Engine){
	incomingRoutes.GET("/season", routes.GetSeasonStats())
	incomingRoutes.GET("/lastmatch", routes.GetLastMatchStats())
}


func main(){
	err := godotenv.Load()
	if err != nil {
		fmt.Println(err)
	}

	router := gin.New()
	Routes(router)

	scraper.InfoUpdate()
	log.Fatal(router.Run(os.Getenv("APP_PORT")))


}