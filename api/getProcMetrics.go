package api

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"thingularity.co/dz-edge-api/models"
	"thingularity.co/dz-edge-api/services"
)

type preprocMetrics struct {
	Process      string  `json:"process"`
	WoID         string  `json:"woID"`
	Operation    int     `json:"operation"`
	Temp         float32 `json:"temp"`
	PH           float32 `json:"ph"`
	Conductivity float32 `json:"conductivity"`
}

type galMetrics struct {
	Process   string  `json:"process"`
	WoID      string  `json:"woID"`
	Operation int     `json:"operation"`
	SetTemp   float32 `json:"setTemp"`
	ActTemp   float32 `json:"actTemp"`
}

// GetProcMetrics ...
func GetProcMetrics(c *gin.Context) {
	operation, err := strconv.Atoi(c.Param("operation")) // matching uri definition
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	switch operation {
	case 1: // Degrease 1
		valTagStr := []string{models.Preproc01Temp, models.Preproc01Ph, models.Preproc01Conduct}
		id, err := services.GetTextTagValue("dotzero", []string{models.Preproc01Id})
		val, err := services.GetTagValue("dotzero", valTagStr)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
			return
		}
		c.JSON(http.StatusOK, preprocMetrics{"Pre-treatment", id[0], operation, float32(val[0]), float32(val[1]), float32(val[2])})
		return
	case 2: // Degrease 2
		valTagStr := []string{models.Preproc02Temp, models.Preproc02Ph, models.Preproc02Conduct}
		id, err := services.GetTextTagValue("dotzero", []string{models.Preproc02Id})
		val, err := services.GetTagValue("dotzero", valTagStr)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
			return
		}
		c.JSON(http.StatusOK, preprocMetrics{"Pre-treatment", id[0], operation, float32(val[0]), float32(val[1]), float32(val[2])})
		return
	case 3: // Pickling 1
		valTagStr := []string{models.Preproc03Temp, models.Preproc03Ph, models.Preproc03Conduct}
		id, err := services.GetTextTagValue("dotzero", []string{models.Preproc03Id})
		val, err := services.GetTagValue("dotzero", valTagStr)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
			return
		}
		c.JSON(http.StatusOK, preprocMetrics{"Pre-treatment", id[0], operation, float32(val[0]), float32(val[1]), float32(val[2])})
		return
	case 4: // Pickling 2
		valTagStr := []string{models.Preproc04Temp, models.Preproc04Ph, models.Preproc04Conduct}
		id, err := services.GetTextTagValue("dotzero", []string{models.Preproc04Id})
		val, err := services.GetTagValue("dotzero", valTagStr)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
			return
		}
		c.JSON(http.StatusOK, preprocMetrics{"Pre-treatment", id[0], operation, float32(val[0]), float32(val[1]), float32(val[2])})
		return
	case 5: // Pickling 3
		valTagStr := []string{models.Preproc05Temp, models.Preproc05Ph, models.Preproc05Conduct}
		id, err := services.GetTextTagValue("dotzero", []string{models.Preproc05Id})
		val, err := services.GetTagValue("dotzero", valTagStr)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
			return
		}
		c.JSON(http.StatusOK, preprocMetrics{"Pre-treatment", id[0], operation, float32(val[0]), float32(val[1]), float32(val[2])})
		return
	case 6: // Galvanizing 1
		valTagStr := []string{models.Gal01TempSet, models.Gal01TempAct}
		id, err := services.GetTextTagValue("dotzero", []string{models.Gal01Id})
		val, err := services.GetTagValue("dotzero", valTagStr)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
			return
		}
		c.JSON(http.StatusOK, galMetrics{"Galvanized", id[0], operation, float32(val[0]), float32(val[1])})
		return
	default:
		c.JSON(http.StatusBadRequest, gin.H{"message": "unknown operation"})
	}
}
