package apiserver

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/render"
	"github.com/mateuszmidor/FlightFinder/pkg/application"
)

func GetRoutes() Routes {
	return routes
}

// Index is the index handler.
func Index(c *gin.Context) {
	c.HTML(http.StatusOK, "map.html", nil)
}

func GetAirports(c *gin.Context) {
	_airportsSVC, ok := c.Get("airports")
	if !ok {
		err := errors.New("airports svc is missing")
		log.Print(err)
		c.AbortWithStatusJSON(http.StatusInternalServerError, errorToJSON(err))
		return
	}

	airportsSVC, ok := _airportsSVC.(*application.Airports)
	if !ok {
		err := fmt.Errorf("airports svc is of wrong type: %T", _airportsSVC)
		log.Print(err)
		c.AbortWithStatusJSON(http.StatusInternalServerError, errorToJSON(err))
		return
	}

	getAirports(airportsSVC, c)
}

// FindFromToConnection - Flight connections by from and to airport IATA codes
func FindFromToConnection(c *gin.Context) {
	_finder, ok := c.Get("finder")
	if !ok {
		err := errors.New("finder svc is missing")
		log.Print(err)
		c.AbortWithStatusJSON(http.StatusInternalServerError, errorToJSON(err))
		return
	}

	finder, ok := _finder.(*application.ConnectionFinder)
	if !ok {
		err := fmt.Errorf("finder svc is of wrong type: %T", _finder)
		log.Print(err)
		c.AbortWithStatusJSON(http.StatusInternalServerError, errorToJSON(err))
		return
	}

	find(finder, c)
}

func getAirports(svc *application.Airports, c *gin.Context) {
	apiAirports := []Airport{}
	for _, a := range svc.AllAirports() {
		airports := Airport{Code: a.Code(), Name: a.Name(), Nation: a.Nation(), NationFullName: "", Lon: float32(a.Longitude()), Lat: float32(a.Latitude())}
		apiAirports = append(apiAirports, airports)
	}
	buff := &bytes.Buffer{}
	json.NewEncoder(buff).Encode(apiAirports)
	c.Render(
		http.StatusOK,
		render.Data{
			ContentType: "application/json",
			Data:        buff.Bytes(),
		})
}

func find(finder *application.ConnectionFinder, c *gin.Context) {
	from := getFromAirportCode(c.Request)
	to := getToAirportCode(c.Request)
	maxSegmentCount := getMaxSegmentsCount(c.Request)
	log.Printf("%s -> %s...\n", from, to)

	buff := &bytes.Buffer{}
	pathRenderer := NewPathRendererAsJSON(buff)
	err := finder.Find(from, to, maxSegmentCount, pathRenderer)

	// FIND ERROR
	if err != nil {
		log.Printf("%s -> %s: ERROR: %v\n", from, to, err)
		c.AbortWithStatusJSON(http.StatusBadRequest, errorToJSON(err))
		return
	}

	// FIND OK
	log.Printf("%s -> %s: OK\n", from, to)
	c.Render(
		http.StatusOK,
		render.Data{
			ContentType: "application/json",
			Data:        buff.Bytes(),
		})
}

func getFromAirportCode(r *http.Request) string {
	return strings.ToUpper(r.FormValue("from"))
}

func getToAirportCode(r *http.Request) string {
	return strings.ToUpper(r.FormValue("to"))
}

func getMaxSegmentsCount(r *http.Request) int {
	count, _ := strconv.Atoi(r.FormValue("maxsegmentcount"))
	return count
}

func errorToJSON(err error) interface{} {
	type ErrorJSON struct {
		Error string `json:"error"`
	}

	return ErrorJSON{Error: err.Error()}
}
