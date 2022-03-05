package app

import (
	"log"
	"path"

	"github.com/gin-gonic/gin"
	apiserver "github.com/mateuszmidor/FlightFinder/cmd/finder_web/app/apiserver/go"
	"github.com/mateuszmidor/FlightFinder/pkg/application"
	"github.com/mateuszmidor/FlightFinder/pkg/infrastructure/csv"
)

func Run(http_port, flights_data_dir, web_data_dir string) {
	router := gin.Default()
	router.LoadHTMLGlob(path.Join(web_data_dir, "*.html"))
	router.StaticFile("favicon.ico", path.Join(web_data_dir, "favicon.ico"))
	router.Use(allowLocalSwaggerPreviewCORS)
	router.Use(finder(flights_data_dir))
	router.Use(airports(flights_data_dir))
	for _, r := range apiserver.GetRoutes() {
		router.Handle(r.Method, r.Pattern, r.HandlerFunc)
	}
	log.Fatal(router.Run(":" + http_port))
}

func allowLocalSwaggerPreviewCORS(c *gin.Context) {
	const SwaggerViewerLocalhost = "http://localhost:18512"
	c.Header("Access-Control-Allow-Origin", SwaggerViewerLocalhost)
	c.Header("Access-Control-Allow-Methods", "PUT, POST, GET, DELETE, OPTIONS")
}

func finder(csv_dir string) func(*gin.Context) {
	repo := csv.NewFlightsDataRepoCSV(csv_dir)
	finder := application.NewConnectionFinder(repo)

	return func(c *gin.Context) {
		c.Set("finder", finder)
	}

}

func airports(csv_dir string) func(*gin.Context) {
	repo := csv.NewFlightsDataRepoCSV(csv_dir)
	airports := application.NewAirports(repo)

	return func(c *gin.Context) {
		c.Set("airports", airports)
	}

}
