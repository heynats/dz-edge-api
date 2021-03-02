package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"thingularity.co/dz-edge-api/models"
	"thingularity.co/dz-edge-api/services"
)

// UpdateJob ...
func UpdateJob(c *gin.Context) {
	type req struct {
		JobID   string `json:"woID" binding:"required"`
		Process string `json:"process" binding:"required"`
		EventID *int   `json:"eventID" binding:"required"`
	}
	var r req
	if err := c.ShouldBindJSON(&r); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"code": "u001", "message": err.Error()})
		return
	}
	fmt.Println("JobID: " + r.JobID + "  Process: " + r.Process + "  EventID: " + strconv.Itoa(*r.EventID))

	// Handling acceptable process types 'pre-treatment' (pre) and 'galvanization' (gal) ONLY
	switch r.Process {
	case "pre":
		index, err := models.UpdateJob(r.Process, r.JobID, *r.EventID)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"code": "u002", "message": err.Error()})
			return
		}
		if index < 0 {
			c.JSON(http.StatusInternalServerError, gin.H{"code": "u002", "message": "update pre-treatment job index out of bound"})
			return
		}

		switch *r.EventID {
		case 10, 11, 12:
			// Write status code to PLC using WebAccess API
			err = services.SetTagValue("dotzero", models.WStatTags[index], *r.EventID)
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"code": "u003", "message": err.Error()})
				return
			}

			c.JSON(http.StatusOK, gin.H{"woID": r.JobID, "queueIndex": index})
			return

		case 0:
			// Get metric value from SCADA system
			val, err := services.GetTagValue("dotzero", models.GetJobMetrics(index))
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"code": "u002", "message": err.Error()})
				return
			}
			metrics := services.Route03Metrics{
				PicklingSmallTime:   strconv.FormatFloat(val[0], 'f', 2, 64) + "," + strconv.FormatFloat(val[1], 'f', 2, 64) + "," + strconv.FormatFloat(val[2], 'f', 2, 64),
				DegreasingSmallTime: strconv.FormatFloat(val[3], 'f', 2, 64) + "," + strconv.FormatFloat(val[4], 'f', 2, 64) + "," + strconv.FormatFloat(val[5], 'f', 2, 64),
				HotWaterTime:        strconv.FormatFloat(val[6], 'f', 2, 64) + "," + strconv.FormatFloat(val[7], 'f', 2, 64) + "," + strconv.FormatFloat(val[8], 'f', 2, 64),
				DegreasingTime:      strconv.FormatFloat(val[9], 'f', 2, 64) + "," + strconv.FormatFloat(val[10], 'f', 2, 64) + "," + strconv.FormatFloat(val[11], 'f', 2, 64),
				DegreasingWaterTime: strconv.FormatFloat(val[12], 'f', 2, 64) + "," + strconv.FormatFloat(val[13], 'f', 2, 64) + "," + strconv.FormatFloat(val[14], 'f', 2, 64),
				FluxTime:            strconv.FormatFloat(val[15], 'f', 2, 64) + "," + strconv.FormatFloat(val[16], 'f', 2, 64) + "," + strconv.FormatFloat(val[17], 'f', 2, 64),
				DryingTime:          strconv.FormatFloat(val[18], 'f', 2, 64) + "," + strconv.FormatFloat(val[19], 'f', 2, 64) + "," + strconv.FormatFloat(val[20], 'f', 2, 64),
				PicklingTime:        strconv.FormatFloat(val[21], 'f', 2, 64) + "," + strconv.FormatFloat(val[22], 'f', 2, 64) + "," + strconv.FormatFloat(val[23], 'f', 2, 64) + "," + strconv.FormatFloat(val[24], 'f', 2, 64) + "," + strconv.FormatFloat(val[25], 'f', 2, 64) + "," + strconv.FormatFloat(val[26], 'f', 2, 64),
				PicklingWaterTime:   strconv.FormatFloat(val[27], 'f', 2, 64) + "," + strconv.FormatFloat(val[28], 'f', 2, 64) + "," + strconv.FormatFloat(val[29], 'f', 2, 64),
				LessFluxTime:        strconv.FormatFloat(val[30], 'f', 2, 64) + "," + strconv.FormatFloat(val[31], 'f', 2, 64) + "," + strconv.FormatFloat(val[31], 'f', 2, 64),
				TurnOverTime:        strconv.FormatFloat(val[33], 'f', 2, 64) + "," + strconv.FormatFloat(val[34], 'f', 2, 64) + "," + strconv.FormatFloat(val[35], 'f', 2, 64),
			}
			jsonstr, err := json.Marshal(metrics)
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"code": "u002", "message": err.Error()})
				return
			}
			// Update job metrics with iShopFloor API
			if err = services.UpdateRouteMetrics(r.JobID, "3", string(jsonstr)); err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"code": "u002", "message": err.Error()})
				return
			}
			// Update job status with iShopFloor API
			// err = services.SetJobDone(r.JobID, "1", "", "")
			// if err != nil {
			// 	c.JSON(http.StatusInternalServerError, gin.H{"code": "u002", "message": err.Error()})
			// 	return
			// }
			// Set job complete code in PLC with WebAccess API
			err = services.SetTagValue("dotzero", models.W01Status, 0)
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"code": "u003", "message": err.Error()})
				return
			}

			c.JSON(http.StatusOK, gin.H{"woID": r.JobID, "queueIndex": index})
			return

		default:
			c.JSON(http.StatusBadRequest, gin.H{"code": "u001", "message": "unknown EventID for the pre-treatment process"})
			return
		}

	case "gal":
		index, err := models.UpdateJob(r.Process, r.JobID, *r.EventID)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"code": "u002", "message": err.Error()})
			return
		}

		switch *r.EventID {
		case 10, 11, 12:
			switch index {
			// Write status code to PLC using WebAccess API
			case 0:
				err = services.SetTagValue("dotzero", models.G01Status, *r.EventID)
				if err != nil {
					c.JSON(http.StatusInternalServerError, gin.H{"code": "u003", "message": err.Error()})
					return
				}
			case 1:
				err = services.SetTagValue("dotzero", models.G02Status, *r.EventID)
				if err != nil {
					c.JSON(http.StatusInternalServerError, gin.H{"code": "u003", "message": err.Error()})
					return
				}
			default:
				c.JSON(http.StatusInternalServerError, gin.H{"code": "u001", "message": "update galvanizing job index out of bound"})
				return
			}
			c.JSON(http.StatusOK, gin.H{"woID": r.JobID, "queueIndex": index})
			return
		case 0:
			// Job finish procedure (r.EventID = 0)
			switch index {
			case 0:
				// Get metric value from SCADA system
				tags := []string{}
				wtLilBlu := models.GetTagArray(models.G01WtLilblu, 5)
				timeLilBlu := models.GetTagArray(models.G01TimeLilblu, 5)
				tempLilBlu := models.GetTagArray(models.G01TempLilblu, 5)
				tags = append(tags, wtLilBlu...)
				tagss := append(timeLilBlu, tempLilBlu...)
				tags = append(tags, tagss...)
				val, err := services.GetTagValue("dotzero", tags)
				if err != nil {
					c.JSON(http.StatusInternalServerError, gin.H{"code": "u002", "message": err.Error()})
					return
				}
				metrics := services.Route04Metrics{
					SubBasketWeight: strconv.FormatFloat(val[0], 'f', 2, 64) + "," + strconv.FormatFloat(val[1], 'f', 2, 64) + "," + strconv.FormatFloat(val[2], 'f', 2, 64) + "," + strconv.FormatFloat(val[3], 'f', 2, 64) + "," + strconv.FormatFloat(val[4], 'f', 2, 64),
					SubBasketTime:   strconv.FormatFloat(val[5], 'f', 2, 64) + "," + strconv.FormatFloat(val[6], 'f', 2, 64) + "," + strconv.FormatFloat(val[7], 'f', 2, 64) + "," + strconv.FormatFloat(val[8], 'f', 2, 64) + "," + strconv.FormatFloat(val[9], 'f', 2, 64),
					SubBasketTemp:   strconv.FormatFloat(val[10], 'f', 2, 64) + "," + strconv.FormatFloat(val[11], 'f', 2, 64) + "," + strconv.FormatFloat(val[12], 'f', 2, 64) + "," + strconv.FormatFloat(val[13], 'f', 2, 64) + "," + strconv.FormatFloat(val[14], 'f', 2, 64),
				}
				jsonstr, err := json.Marshal(metrics)
				if err != nil {
					c.JSON(http.StatusInternalServerError, gin.H{"code": "u002", "message": err.Error()})
					return
				}
				// Update job metrics with iShopFloor API
				if err = services.UpdateRouteMetrics(r.JobID, "4", string(jsonstr)); err != nil {
					c.JSON(http.StatusInternalServerError, gin.H{"code": "u002", "message": err.Error()})
					return
				}

				// defectWt := val[1] - val[0] - val[3]
				// err = services.SetJobDone(r.JobID, "2", strconv.FormatFloat(val[3], 'f', -1, 64), strconv.FormatFloat(defectWt, 'f', -1, 64))
				// if err != nil {
				// 	c.JSON(http.StatusInternalServerError, gin.H{"code": "u002", "message": err.Error()})
				// 	return
				// }

				// Mark job complete
				err = services.SetTagValue("dotzero", models.G01Status, 0)
				if err != nil {
					c.JSON(http.StatusInternalServerError, gin.H{"code": "u003", "message": err.Error()})
					return
				}
			case 1:
				// Get metric value from SCADA system
				tags := []string{}
				wtLilBlu := models.GetTagArray(models.G02WtLilblu, 5)
				timeLilBlu := models.GetTagArray(models.G02TimeLilblu, 5)
				tempLilBlu := models.GetTagArray(models.G02TempLilblu, 5)
				tags = append(tags, wtLilBlu...)
				tagss := append(timeLilBlu, tempLilBlu...)
				tags = append(tags, tagss...)
				val, err := services.GetTagValue("dotzero", tags)
				if err != nil {
					c.JSON(http.StatusInternalServerError, gin.H{"code": "u002", "message": err.Error()})
					return
				}
				metrics := services.Route04Metrics{
					SubBasketWeight: strconv.FormatFloat(val[0], 'f', 2, 64) + "," + strconv.FormatFloat(val[1], 'f', 2, 64) + "," + strconv.FormatFloat(val[2], 'f', 2, 64) + "," + strconv.FormatFloat(val[3], 'f', 2, 64) + "," + strconv.FormatFloat(val[4], 'f', 2, 64),
					SubBasketTime:   strconv.FormatFloat(val[5], 'f', 2, 64) + "," + strconv.FormatFloat(val[6], 'f', 2, 64) + "," + strconv.FormatFloat(val[7], 'f', 2, 64) + "," + strconv.FormatFloat(val[8], 'f', 2, 64) + "," + strconv.FormatFloat(val[9], 'f', 2, 64),
					SubBasketTemp:   strconv.FormatFloat(val[10], 'f', 2, 64) + "," + strconv.FormatFloat(val[11], 'f', 2, 64) + "," + strconv.FormatFloat(val[12], 'f', 2, 64) + "," + strconv.FormatFloat(val[13], 'f', 2, 64) + "," + strconv.FormatFloat(val[14], 'f', 2, 64),
				}
				jsonstr, err := json.Marshal(metrics)
				if err != nil {
					c.JSON(http.StatusInternalServerError, gin.H{"code": "u002", "message": err.Error()})
					return
				}
				// Update job metrics with iShopFloor API
				if err = services.UpdateRouteMetrics(r.JobID, "4", string(jsonstr)); err != nil {
					c.JSON(http.StatusInternalServerError, gin.H{"code": "u002", "message": err.Error()})
					return
				}

				// defectWt := val[1] - val[0] - val[3]
				// err = services.SetJobDone(r.JobID, "2", strconv.FormatFloat(val[3], 'f', -1, 64), strconv.FormatFloat(defectWt, 'f', -1, 64))
				// if err != nil {
				// 	c.JSON(http.StatusInternalServerError, gin.H{"code": "u002", "message": err.Error()})
				// 	return
				// }

				// Mark job complete
				err = services.SetTagValue("dotzero", models.G02Status, 0)
				if err != nil {
					c.JSON(http.StatusInternalServerError, gin.H{"code": "u003", "message": err.Error()})
					return
				}
			default:
				c.JSON(http.StatusInternalServerError, gin.H{"code": "u001", "message": "update galvanizing job index out of bound"})
				return
			}
			c.JSON(http.StatusOK, gin.H{"woID": r.JobID, "queueIndex": index})
			return

		default:
			c.JSON(http.StatusBadRequest, gin.H{"code": "u001", "message": "unknown EventID for the galvanizing process"})
			return
		}

	default:
		c.JSON(http.StatusBadRequest, gin.H{"code": "u001", "message": "Unknown process. Process must be 'pre' (pre-treatment) or 'gal' (galvanization)."})
	}
}
