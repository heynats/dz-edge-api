package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"thingularity.co/dz-edge-api/models"
	"thingularity.co/dz-edge-api/services"
)

// DeleteJob ...
func DeleteJob(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "DeleteJob API coming soon!"})
}

// DeleteAllJobs ...
func DeleteAllJobs(c *gin.Context) {
	models.RemoveAllJobs()
	// preproc
	err := services.SetTagValue("dotzero", models.W01Status, 0)
	err = services.SetTagValue("dotzero", models.W02Status, 0)
	err = services.SetTagValue("dotzero", models.W03Status, 0)
	err = services.SetTagValue("dotzero", models.W04Status, 0)
	err = services.SetTagValue("dotzero", models.W05Status, 0)
	// err = services.SetTagValue("dotzero", models.W06Status, 0)
	// err = services.SetTagValue("dotzero", models.W07Status, 0)
	// err = services.SetTagValue("dotzero", models.W08Status, 0)
	// err = services.SetTagValue("dotzero", models.W09Status, 0)
	// err = services.SetTagValue("dotzero", models.W10Status, 0)
	// err = services.SetTagValue("dotzero", models.W11Status, 0)
	// err = services.SetTagValue("dotzero", models.W12Status, 0)
	// err = services.SetTagValue("dotzero", models.W13Status, 0)
	// err = services.SetTagValue("dotzero", models.W14Status, 0)
	// err = services.SetTagValue("dotzero", models.W15Status, 0)
	// err = services.SetTagValue("dotzero", models.W16Status, 0)
	// err = services.SetTagValue("dotzero", models.W17Status, 0)
	// err = services.SetTagValue("dotzero", models.W18Status, 0)
	// err = services.SetTagValue("dotzero", models.W19Status, 0)
	// err = services.SetTagValue("dotzero", models.W20Status, 0)
	// galvanization
	err = services.SetTagValue("dotzero", models.G01Status, 0)
	err = services.SetTagValue("dotzero", models.G02Status, 0)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "removed all jobs"})
}

// DeletePreJobs ...
func DeletePreJobs(c *gin.Context) {
	models.RemovePreJobs()
	// preproc
	err := services.SetTagValue("dotzero", models.W01Status, 0)
	err = services.SetTagValue("dotzero", models.W02Status, 0)
	err = services.SetTagValue("dotzero", models.W03Status, 0)
	err = services.SetTagValue("dotzero", models.W04Status, 0)
	err = services.SetTagValue("dotzero", models.W05Status, 0)
	// err = services.SetTagValue("dotzero", models.W06Status, 0)
	// err = services.SetTagValue("dotzero", models.W07Status, 0)
	// err = services.SetTagValue("dotzero", models.W08Status, 0)
	// err = services.SetTagValue("dotzero", models.W09Status, 0)
	// err = services.SetTagValue("dotzero", models.W10Status, 0)
	// err = services.SetTagValue("dotzero", models.W11Status, 0)
	// err = services.SetTagValue("dotzero", models.W12Status, 0)
	// err = services.SetTagValue("dotzero", models.W13Status, 0)
	// err = services.SetTagValue("dotzero", models.W14Status, 0)
	// err = services.SetTagValue("dotzero", models.W15Status, 0)
	// err = services.SetTagValue("dotzero", models.W16Status, 0)
	// err = services.SetTagValue("dotzero", models.W17Status, 0)
	// err = services.SetTagValue("dotzero", models.W18Status, 0)
	// err = services.SetTagValue("dotzero", models.W19Status, 0)
	// err = services.SetTagValue("dotzero", models.W20Status, 0)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "removed all preprocess jobs"})
}

// DeleteGalJobs ...
func DeleteGalJobs(c *gin.Context) {
	models.RemoveGalJobs()
	// galvanization
	err := services.SetTagValue("dotzero", models.G01Status, 0)
	err = services.SetTagValue("dotzero", models.G02Status, 0)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "removed all galvanization jobs"})
}

// DeletePreJobByID ...
func DeletePreJobByID(c *gin.Context) {
	if c.Param("jobId") == "" {
		c.JSON(http.StatusBadRequest, gin.H{"message": "missing job ID"})
		return
	}
	id := c.Param("jobId")
	index := models.JobExists("pre", id)
	if index < 0 {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "job #" + id + " does not exist"})
		return
	}
	err := services.SetTagValue("dotzero", models.WStatTags[index], 0)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	models.RemoveJob("pre", index)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "removed preprocess job #" + id})
}

// DeleteGalJobByID ...
func DeleteGalJobByID(c *gin.Context) {
	if c.Param("jobId") == "" {
		c.JSON(http.StatusBadRequest, gin.H{"message": "missing job ID"})
		return
	}
	id := c.Param("jobId")
	index := models.JobExists("gal", id)
	if index < 0 {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "job #" + id + " does not exist"})
		return
	}
	err := services.SetTagValue("dotzero", models.GStatTags[index], 0)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	models.RemoveJob("gal", index)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "removed galvanization job #" + id})
}
