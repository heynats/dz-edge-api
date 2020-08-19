package api

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"thingularity.co/dz-edge-api/models"
	"thingularity.co/dz-edge-api/services"
)

// CreateJob ...
func CreateJob(c *gin.Context) {
	// Handle request body with gin's wonderful marshalling feature
	type req struct {
		JobID   string `json:"woID" binding:"required"`
		Process string `json:"process" binding:"required"`
	}
	var r req
	if err := c.ShouldBindJSON(&r); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	fmt.Println("JobID: " + r.JobID + "  Process: " + r.Process)

	// Handling acceptable process types 'pre-treatment' (pre) and 'galvanization' (gal) ONLY
	switch r.Process {
	case "pre":
		index, err := models.AddJob(r.Process, r.JobID, 1)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
			return
		}
		if index < 0 {
			c.JSON(http.StatusInternalServerError, gin.H{"message": "add pre-treatment job index out of bound"})
			return
		}
		// Implement WebAccess integration, create job procedure:
		// 1) set status tag of the next empty buffer area to 1, 2) set job ID tag
		err = services.SetTagValue("dotzero", models.WStatTags[index], 1)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
			models.RemoveJob(r.Process, index)
			return
		}
		err = services.SetTextTagValue("dotzero", models.WIDTags[index], r.JobID)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
			models.RemoveJob(r.Process, index)
			return
		}

		c.JSON(http.StatusOK, gin.H{"woID": r.JobID, "queueIndex": index})
		return

	case "gal":
		index, err := models.AddJob(r.Process, r.JobID, 1)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
			return
		}
		if index < 0 {
			c.JSON(http.StatusInternalServerError, gin.H{"message": "add galvanizing job index out of bound"})
			return
		}
		err = services.SetTagValue("dotzero", models.GStatTags[index], 1)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
			models.RemoveJob(r.Process, index)
			return
		}
		err = services.SetTextTagValue("dotzero", models.GIDTags[index], r.JobID)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
			models.RemoveJob(r.Process, index)
			return
		}

		c.JSON(http.StatusOK, gin.H{"woID": r.JobID, "queueIndex": index})
		return

	default:
		c.JSON(http.StatusBadRequest, gin.H{"message": "Unknown process. Process must be 'pre' (pre-treatment) or 'gal' (galvanization)."})
	}
}
