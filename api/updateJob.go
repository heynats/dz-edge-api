package api

import (
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
		// WebAccess integration
		switch *r.EventID {
		case 10, 11, 12:
			err = services.SetTagValue("dotzero", models.WStatTags[index], *r.EventID)
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"code": "u003", "message": err.Error()})
				return
			}
			/*
				switch index {
				case 0:
					err = services.SetTagValue("dotzero", models.W01Status, *r.EventID)
					if err != nil {
						c.JSON(http.StatusInternalServerError, gin.H{"code": "u003", "message": err.Error()})
						return
					}
				case 1:
					err = services.SetTagValue("dotzero", models.W02Status, *r.EventID)
					if err != nil {
						c.JSON(http.StatusInternalServerError, gin.H{"code": "u003", "message": err.Error()})
						return
					}
				case 2:
					err = services.SetTagValue("dotzero", models.W03Status, *r.EventID)
					if err != nil {
						c.JSON(http.StatusInternalServerError, gin.H{"code": "u003", "message": err.Error()})
						return
					}
				case 3:
					err = services.SetTagValue("dotzero", models.W04Status, *r.EventID)
					if err != nil {
						c.JSON(http.StatusInternalServerError, gin.H{"code": "u003", "message": err.Error()})
						return
					}
				case 4:
					err = services.SetTagValue("dotzero", models.W05Status, *r.EventID)
					if err != nil {
						c.JSON(http.StatusInternalServerError, gin.H{"code": "u003", "message": err.Error()})
						return
					}
				default:
					c.JSON(http.StatusInternalServerError, gin.H{"code": "u002", "message": "update pre-treatment job index out of bound"})
					return
				}
			*/
			c.JSON(http.StatusOK, gin.H{"woID": r.JobID, "queueIndex": index})
			return
		case 0:
			// Job finish procedure (r.EventID = 0)
			val, err := services.GetTagValue("dotzero", models.GetJobMetrics(index))
			metrics := services.MetricsFields{
				EmptyBucketWeight: strconv.FormatFloat(val[0], 'f', -1, 64),
				FullBucketWeight:  strconv.FormatFloat(val[1], 'f', -1, 64),
				DegreasingTime:    strconv.FormatFloat(val[2], 'f', -1, 64) + "," + strconv.FormatFloat(val[3], 'f', -1, 64) + "," + strconv.FormatFloat(val[4], 'f', -1, 64),
				PicklingTime:      strconv.FormatFloat(val[5], 'f', -1, 64) + "," + strconv.FormatFloat(val[6], 'f', -1, 64) + "," + strconv.FormatFloat(val[7], 'f', -1, 64),
				FluxTime:          strconv.FormatFloat(val[8], 'f', -1, 64),
				Temp:              strconv.FormatFloat(val[9], 'f', -1, 64) + "," + strconv.FormatFloat(val[10], 'f', -1, 64) + "," + strconv.FormatFloat(val[11], 'f', -1, 64) + "," + strconv.FormatFloat(val[12], 'f', -1, 64) + "," + strconv.FormatFloat(val[13], 'f', -1, 64),
			}
			if err := services.UpdateMetrics(r.JobID, "1", metrics); err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"code": "u002", "message": err.Error()})
				return
			}
			err = services.SetJobDone(r.JobID, "1", "", "")
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"code": "u002", "message": err.Error()})
				return
			}
			// Set job complete
			err = services.SetTagValue("dotzero", models.W01Status, 0)
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"code": "u003", "message": err.Error()})
				return
			}
			/*
				switch index {
				case 0:
					// Call iShopFloor API
					wtTags := []string{models.W01WtBktEmpt, models.W01WtBktFull}
					degTags := []string{models.W01TimeDegrease1, models.W01TimeDegrease2, models.W01TimeDegrease3}
					pickTags := []string{models.W01TimePickling1, models.W01TimePickling2, models.W01TimePickling3}
					fluxTags := []string{models.W01TimeFlux}
					tempTags := []string{models.Preproc01Temp, models.Preproc02Temp, models.Preproc03Temp, models.Preproc04Temp, models.Preproc05Temp}
					tags := append(wtTags, degTags...)
					tagss := append(pickTags, fluxTags...)
					tags = append(tags, tagss...)
					tags = append(tags, tempTags...)
					val, err := services.GetTagValue("dotzero", tags)

					metrics := services.MetricsFields{
						EmptyBucketWeight: strconv.FormatFloat(val[0], 'f', -1, 64),
						FullBucketWeight:  strconv.FormatFloat(val[1], 'f', -1, 64),
						DegreasingTime:    strconv.FormatFloat(val[2], 'f', -1, 64) + "," + strconv.FormatFloat(val[3], 'f', -1, 64) + "," + strconv.FormatFloat(val[4], 'f', -1, 64),
						PicklingTime:      strconv.FormatFloat(val[5], 'f', -1, 64) + "," + strconv.FormatFloat(val[6], 'f', -1, 64) + "," + strconv.FormatFloat(val[7], 'f', -1, 64),
						FluxTime:          strconv.FormatFloat(val[8], 'f', -1, 64),
						Temp:              strconv.FormatFloat(val[9], 'f', -1, 64) + "," + strconv.FormatFloat(val[10], 'f', -1, 64) + "," + strconv.FormatFloat(val[11], 'f', -1, 64) + "," + strconv.FormatFloat(val[12], 'f', -1, 64) + "," + strconv.FormatFloat(val[13], 'f', -1, 64),
					}

					if err := services.UpdateMetrics(r.JobID, "1", metrics); err != nil {
						c.JSON(http.StatusInternalServerError, gin.H{"code": "u002", "message": err.Error()})
						return
					}
					// Set job complete
					err = services.SetTagValue("dotzero", models.W01Status, 0)
					if err != nil {
						c.JSON(http.StatusInternalServerError, gin.H{"code": "u003", "message": err.Error()})
						return
					}
				case 1:
					// Call iShopFloor API
					wtTags := []string{models.W02WtBktEmpt, models.W02WtBktFull}
					degTags := []string{models.W02TimeDegrease1, models.W02TimeDegrease2, models.W02TimeDegrease3}
					pickTags := []string{models.W02TimePickling1, models.W02TimePickling2, models.W02TimePickling3}
					fluxTags := []string{models.W02TimeFlux}
					tempTags := []string{models.Preproc01Temp, models.Preproc02Temp, models.Preproc03Temp, models.Preproc04Temp, models.Preproc05Temp}
					tags := append(wtTags, degTags...)
					tagss := append(pickTags, fluxTags...)
					tags = append(tags, tagss...)
					tags = append(tags, tempTags...)
					val, err := services.GetTagValue("dotzero", tags)

					metrics := services.MetricsFields{
						EmptyBucketWeight: strconv.FormatFloat(val[0], 'f', -1, 64),
						FullBucketWeight:  strconv.FormatFloat(val[1], 'f', -1, 64),
						DegreasingTime:    strconv.FormatFloat(val[2], 'f', -1, 64) + "," + strconv.FormatFloat(val[3], 'f', -1, 64) + "," + strconv.FormatFloat(val[4], 'f', -1, 64),
						PicklingTime:      strconv.FormatFloat(val[5], 'f', -1, 64) + "," + strconv.FormatFloat(val[6], 'f', -1, 64) + "," + strconv.FormatFloat(val[7], 'f', -1, 64),
						FluxTime:          strconv.FormatFloat(val[8], 'f', -1, 64),
						Temp:              strconv.FormatFloat(val[9], 'f', -1, 64) + "," + strconv.FormatFloat(val[10], 'f', -1, 64) + "," + strconv.FormatFloat(val[11], 'f', -1, 64) + "," + strconv.FormatFloat(val[12], 'f', -1, 64) + "," + strconv.FormatFloat(val[13], 'f', -1, 64),
					}

					if err := services.UpdateMetrics(r.JobID, "1", metrics); err != nil {
						c.JSON(http.StatusInternalServerError, gin.H{"code": "u002", "message": err.Error()})
						return
					}
					// Set job complete
					err = services.SetTagValue("dotzero", models.W02Status, 0)
					if err != nil {
						c.JSON(http.StatusInternalServerError, gin.H{"code": "u003", "message": err.Error()})
						return
					}
				case 2:
					// Call iShopFloor API
					wtTags := []string{models.W03WtBktEmpt, models.W03WtBktFull}
					degTags := []string{models.W03TimeDegrease1, models.W03TimeDegrease2, models.W03TimeDegrease3}
					pickTags := []string{models.W03TimePickling1, models.W03TimePickling2, models.W03TimePickling3}
					fluxTags := []string{models.W03TimeFlux}
					tempTags := []string{models.Preproc01Temp, models.Preproc02Temp, models.Preproc03Temp, models.Preproc04Temp, models.Preproc05Temp}
					tags := append(wtTags, degTags...)
					tagss := append(pickTags, fluxTags...)
					tags = append(tags, tagss...)
					tags = append(tags, tempTags...)
					val, err := services.GetTagValue("dotzero", tags)

					metrics := services.MetricsFields{
						EmptyBucketWeight: strconv.FormatFloat(val[0], 'f', -1, 64),
						FullBucketWeight:  strconv.FormatFloat(val[1], 'f', -1, 64),
						DegreasingTime:    strconv.FormatFloat(val[2], 'f', -1, 64) + "," + strconv.FormatFloat(val[3], 'f', -1, 64) + "," + strconv.FormatFloat(val[4], 'f', -1, 64),
						PicklingTime:      strconv.FormatFloat(val[5], 'f', -1, 64) + "," + strconv.FormatFloat(val[6], 'f', -1, 64) + "," + strconv.FormatFloat(val[7], 'f', -1, 64),
						FluxTime:          strconv.FormatFloat(val[8], 'f', -1, 64),
						Temp:              strconv.FormatFloat(val[9], 'f', -1, 64) + "," + strconv.FormatFloat(val[10], 'f', -1, 64) + "," + strconv.FormatFloat(val[11], 'f', -1, 64) + "," + strconv.FormatFloat(val[12], 'f', -1, 64) + "," + strconv.FormatFloat(val[13], 'f', -1, 64),
					}

					if err := services.UpdateMetrics(r.JobID, "1", metrics); err != nil {
						c.JSON(http.StatusInternalServerError, gin.H{"code": "u002", "message": err.Error()})
						return
					}
					// Set job complete
					err = services.SetTagValue("dotzero", models.W03Status, 0)
					if err != nil {
						c.JSON(http.StatusInternalServerError, gin.H{"code": "u003", "message": err.Error()})
						return
					}
				case 3:
					// Call iShopFloor API
					wtTags := []string{models.W04WtBktEmpt, models.W04WtBktFull}
					degTags := []string{models.W04TimeDegrease1, models.W04TimeDegrease2, models.W04TimeDegrease3}
					pickTags := []string{models.W04TimePickling1, models.W04TimePickling2, models.W04TimePickling3}
					fluxTags := []string{models.W04TimeFlux}
					tempTags := []string{models.Preproc01Temp, models.Preproc02Temp, models.Preproc03Temp, models.Preproc04Temp, models.Preproc05Temp}
					tags := append(wtTags, degTags...)
					tagss := append(pickTags, fluxTags...)
					tags = append(tags, tagss...)
					tags = append(tags, tempTags...)
					val, err := services.GetTagValue("dotzero", tags)

					metrics := services.MetricsFields{
						EmptyBucketWeight: strconv.FormatFloat(val[0], 'f', -1, 64),
						FullBucketWeight:  strconv.FormatFloat(val[1], 'f', -1, 64),
						DegreasingTime:    strconv.FormatFloat(val[2], 'f', -1, 64) + "," + strconv.FormatFloat(val[3], 'f', -1, 64) + "," + strconv.FormatFloat(val[4], 'f', -1, 64),
						PicklingTime:      strconv.FormatFloat(val[5], 'f', -1, 64) + "," + strconv.FormatFloat(val[6], 'f', -1, 64) + "," + strconv.FormatFloat(val[7], 'f', -1, 64),
						FluxTime:          strconv.FormatFloat(val[8], 'f', -1, 64),
						Temp:              strconv.FormatFloat(val[9], 'f', -1, 64) + "," + strconv.FormatFloat(val[10], 'f', -1, 64) + "," + strconv.FormatFloat(val[11], 'f', -1, 64) + "," + strconv.FormatFloat(val[12], 'f', -1, 64) + "," + strconv.FormatFloat(val[13], 'f', -1, 64),
					}

					if err := services.UpdateMetrics(r.JobID, "1", metrics); err != nil {
						c.JSON(http.StatusInternalServerError, gin.H{"code": "u002", "message": err.Error()})
						return
					}
					// Set job complete
					err = services.SetTagValue("dotzero", models.W04Status, 0)
					if err != nil {
						c.JSON(http.StatusInternalServerError, gin.H{"code": "u003", "message": err.Error()})
						return
					}
				case 4:
					// Call iShopFloor API
					wtTags := []string{models.W05WtBktEmpt, models.W05WtBktFull}
					degTags := []string{models.W05TimeDegrease1, models.W05TimeDegrease2, models.W05TimeDegrease3}
					pickTags := []string{models.W05TimePickling1, models.W05TimePickling2, models.W05TimePickling3}
					fluxTags := []string{models.W05TimeFlux}
					tempTags := []string{models.Preproc01Temp, models.Preproc02Temp, models.Preproc03Temp, models.Preproc04Temp, models.Preproc05Temp}
					tags := append(wtTags, degTags...)
					tagss := append(pickTags, fluxTags...)
					tags = append(tags, tagss...)
					tags = append(tags, tempTags...)
					val, err := services.GetTagValue("dotzero", tags)

					metrics := services.MetricsFields{
						EmptyBucketWeight: strconv.FormatFloat(val[0], 'f', -1, 64),
						FullBucketWeight:  strconv.FormatFloat(val[1], 'f', -1, 64),
						DegreasingTime:    strconv.FormatFloat(val[2], 'f', -1, 64) + "," + strconv.FormatFloat(val[3], 'f', -1, 64) + "," + strconv.FormatFloat(val[4], 'f', -1, 64),
						PicklingTime:      strconv.FormatFloat(val[5], 'f', -1, 64) + "," + strconv.FormatFloat(val[6], 'f', -1, 64) + "," + strconv.FormatFloat(val[7], 'f', -1, 64),
						FluxTime:          strconv.FormatFloat(val[8], 'f', -1, 64),
						Temp:              strconv.FormatFloat(val[9], 'f', -1, 64) + "," + strconv.FormatFloat(val[10], 'f', -1, 64) + "," + strconv.FormatFloat(val[11], 'f', -1, 64) + "," + strconv.FormatFloat(val[12], 'f', -1, 64) + "," + strconv.FormatFloat(val[13], 'f', -1, 64),
					}

					if err := services.UpdateMetrics(r.JobID, "1", metrics); err != nil {
						c.JSON(http.StatusInternalServerError, gin.H{"code": "u002", "message": err.Error()})
						return
					}
					// Set job complete
					err = services.SetTagValue("dotzero", models.W05Status, 0)
					if err != nil {
						c.JSON(http.StatusInternalServerError, gin.H{"code": "u003", "message": err.Error()})
						return
					}
				default:
					c.JSON(http.StatusInternalServerError, gin.H{"code": "u002", "message": "update pre-treatment job index out of bound"})
					return
				}
			*/
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
		// WebAccess integration
		switch *r.EventID {
		case 10, 11, 12:
			switch index {
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
				c.JSON(http.StatusInternalServerError, gin.H{"code": "u002", "message": "update galvanizing job index out of bound"})
				return
			}
			c.JSON(http.StatusOK, gin.H{"woID": r.JobID, "queueIndex": index})
			return
		case 0:
			// Job finish procedure (r.EventID = 0)
			switch index {
			case 0:
				// Call iShopFloor API
				tags := []string{models.G01WtBktEmpt, models.G01WtBktFull, models.Gal01TempAct, models.G01WtGood}
				wtLilBlu := models.GetTagArray(models.G01WtLilblu, 5)
				timeLilBlu := models.GetTagArray(models.G01TimeLilblu, 5)
				tempLilBlu := models.GetTagArray(models.G01TempLilblu, 5)
				tags = append(tags, wtLilBlu...)
				tagss := append(timeLilBlu, tempLilBlu...)
				tags = append(tags, tagss...)
				val, err := services.GetTagValue("dotzero", tags)

				metrics := services.MetricsFields{
					EmptyBucketWeight: strconv.FormatFloat(val[0], 'f', -1, 64),
					FullBucketWeight:  strconv.FormatFloat(val[1], 'f', -1, 64),
					FurnaceTtemp:      strconv.FormatFloat(val[2], 'f', -1, 64),
					LittleBlueWeight:  strconv.FormatFloat(val[4], 'f', -1, 64) + "," + strconv.FormatFloat(val[5], 'f', -1, 64) + "," + strconv.FormatFloat(val[6], 'f', -1, 64) + "," + strconv.FormatFloat(val[7], 'f', -1, 64) + "," + strconv.FormatFloat(val[8], 'f', -1, 64),
					LittleBlueTime:    strconv.FormatFloat(val[9], 'f', -1, 64) + "," + strconv.FormatFloat(val[10], 'f', -1, 64) + "," + strconv.FormatFloat(val[11], 'f', -1, 64) + "," + strconv.FormatFloat(val[12], 'f', -1, 64) + "," + strconv.FormatFloat(val[13], 'f', -1, 64),
					LittleBlueTemp:    strconv.FormatFloat(val[14], 'f', -1, 64) + "," + strconv.FormatFloat(val[15], 'f', -1, 64) + "," + strconv.FormatFloat(val[16], 'f', -1, 64) + "," + strconv.FormatFloat(val[17], 'f', -1, 64) + "," + strconv.FormatFloat(val[18], 'f', -1, 64),
				}

				if err := services.UpdateMetrics(r.JobID, "2", metrics); err != nil {
					c.JSON(http.StatusInternalServerError, gin.H{"code": "u002", "message": err.Error()})
					return
				}

				defectWt := val[1] - val[0] - val[3]
				err = services.SetJobDone(r.JobID, "2", strconv.FormatFloat(val[3], 'f', -1, 64), strconv.FormatFloat(defectWt, 'f', -1, 64))
				if err != nil {
					c.JSON(http.StatusInternalServerError, gin.H{"code": "u002", "message": err.Error()})
					return
				}

				// Mark job complete
				err = services.SetTagValue("dotzero", models.G01Status, 0)
				if err != nil {
					c.JSON(http.StatusInternalServerError, gin.H{"code": "u003", "message": err.Error()})
					return
				}
			case 1:
				// Call iShopFloor API
				tags := []string{models.G02WtBktEmpt, models.G02WtBktFull, models.Gal01TempAct, models.G02WtGood}
				wtLilBlu := models.GetTagArray(models.G02WtLilblu, 5)
				timeLilBlu := models.GetTagArray(models.G02TimeLilblu, 5)
				tempLilBlu := models.GetTagArray(models.G02TempLilblu, 5)
				tags = append(tags, wtLilBlu...)
				tagss := append(timeLilBlu, tempLilBlu...)
				tags = append(tags, tagss...)
				val, err := services.GetTagValue("dotzero", tags)

				metrics := services.MetricsFields{
					EmptyBucketWeight: strconv.FormatFloat(val[0], 'f', -1, 64),
					FullBucketWeight:  strconv.FormatFloat(val[1], 'f', -1, 64),
					FurnaceTtemp:      strconv.FormatFloat(val[2], 'f', -1, 64),
					LittleBlueWeight:  strconv.FormatFloat(val[4], 'f', -1, 64) + "," + strconv.FormatFloat(val[5], 'f', -1, 64) + "," + strconv.FormatFloat(val[6], 'f', -1, 64) + "," + strconv.FormatFloat(val[7], 'f', -1, 64) + "," + strconv.FormatFloat(val[8], 'f', -1, 64),
					LittleBlueTime:    strconv.FormatFloat(val[9], 'f', -1, 64) + "," + strconv.FormatFloat(val[10], 'f', -1, 64) + "," + strconv.FormatFloat(val[11], 'f', -1, 64) + "," + strconv.FormatFloat(val[12], 'f', -1, 64) + "," + strconv.FormatFloat(val[13], 'f', -1, 64),
					LittleBlueTemp:    strconv.FormatFloat(val[14], 'f', -1, 64) + "," + strconv.FormatFloat(val[15], 'f', -1, 64) + "," + strconv.FormatFloat(val[16], 'f', -1, 64) + "," + strconv.FormatFloat(val[17], 'f', -1, 64) + "," + strconv.FormatFloat(val[18], 'f', -1, 64),
				}

				if err := services.UpdateMetrics(r.JobID, "2", metrics); err != nil {
					c.JSON(http.StatusInternalServerError, gin.H{"code": "u002", "message": err.Error()})
					return
				}

				defectWt := val[1] - val[0] - val[3]
				err = services.SetJobDone(r.JobID, "2", strconv.FormatFloat(val[3], 'f', -1, 64), strconv.FormatFloat(defectWt, 'f', -1, 64))
				if err != nil {
					c.JSON(http.StatusInternalServerError, gin.H{"code": "u002", "message": err.Error()})
					return
				}

				// Mark job complete
				err = services.SetTagValue("dotzero", models.G02Status, 0)
				if err != nil {
					c.JSON(http.StatusInternalServerError, gin.H{"code": "u003", "message": err.Error()})
					return
				}
			default:
				c.JSON(http.StatusInternalServerError, gin.H{"code": "u002", "message": "update galvanizing job index out of bound"})
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
