package services

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

const iUser = "root"
const iPass = "qwer1234"
const apiBaseURL = "https://dev-tgnet.orangicetech.com"

// MetricsFields ...
type MetricsFields struct {
	EmptyBucketWeight string
	FullBucketWeight  string
	DegreasingTime    string
	PicklingTime      string
	FluxTime          string
	Temp              string
	FurnaceTtemp      string
	LittleBlueWeight  string
	LittleBlueTime    string
	LittleBlueTemp    string
}

func getAccessToken() (string, error) {
	var jsonResp map[string]interface{}
	data := url.Values{}
	data.Set("uid", iUser)
	data.Set("username", iUser)
	data.Set("password", iPass)

	req, _ := http.NewRequest("POST", apiBaseURL+"/1/app/login", strings.NewReader(data.Encode()))
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("Content-Length", strconv.Itoa(len(data.Encode())))

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	r, _ := ioutil.ReadAll(resp.Body)
	if err = json.Unmarshal(r, &jsonResp); err != nil {
		return "", err
	}

	return jsonResp["token"].(string), nil
}

// UpdateMetrics ...
func UpdateMetrics(jobID string, procType string, metrics MetricsFields) error {
	var jsonResp map[string]interface{}
	accessToken, err := getAccessToken()
	if err != nil {
		return err
	}

	data := url.Values{}
	switch procType {
	case "1":
		data.Set("degreasingTime", metrics.DegreasingTime)
		data.Set("picklingTime", metrics.PicklingTime)
		data.Set("fluxTime", metrics.FluxTime)
		data.Set("temp", metrics.Temp)
		fallthrough
	case "2":
		data.Set("furnaceTtemp", metrics.FurnaceTtemp)
		data.Set("littleBlueWeight", metrics.LittleBlueWeight)
		data.Set("littleBlueTime", metrics.LittleBlueTime)
		data.Set("littleBlueTemp", metrics.LittleBlueTemp)
		fallthrough
	default:
		data.Set("emptyBucketWeight", metrics.EmptyBucketWeight)
		data.Set("fullBucketWeight", metrics.FullBucketWeight)
		data.Set("status", "4")
		data.Set("endTime", strconv.FormatInt(time.Now().Unix(), 10))
	}
	fmt.Printf("[EdgeAPI] Request body: %v\n", data)

	url := apiBaseURL + "/1/app/work_order/" + jobID + "/route/" + procType + "/parameter"
	if gin.Mode() == gin.DebugMode {
		fmt.Println("[debug] API URL: " + url)
	}
	req, _ := http.NewRequest("PUT", url, strings.NewReader(data.Encode()))
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("Content-Length", strconv.Itoa(len(data.Encode())))
	req.Header.Add("accessToken", accessToken)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	r, _ := ioutil.ReadAll(resp.Body)
	if gin.Mode() == gin.DebugMode {
		fmt.Println("[debug] iShopFloor API response: " + string(r))
	}
	if err = json.Unmarshal(r, &jsonResp); err != nil {
		fmt.Println("[debug] iShopFloor api respond with an error!")
		return err
	}
	if resp.StatusCode >= 400 || jsonResp["success"] != true {
		return errors.New("failed to update metrics fields with iShopFloor API")
	}

	return nil
}

// SetJobDone ...
func SetJobDone(jobID string, procType string, goodWt string, defectWt string) error {
	var jsonResp map[string]interface{}
	accessToken, err := getAccessToken()
	if err != nil {
		return err
	}

	data := url.Values{}
	switch procType {
	case "2":
		data.Set("good", goodWt)
		data.Set("defect", defectWt)
		fallthrough
	default:
		data.Set("status", "3")
		data.Set("endTime", strconv.FormatInt(time.Now().Unix(), 10))
	}
	fmt.Printf("[EdgeAPI] Request body: %v\n", data)

	url := apiBaseURL + "/1/app/work_order/" + jobID + "/route/" + procType
	if gin.Mode() == gin.DebugMode {
		fmt.Println("[debug] API URL: " + url)
	}
	req, _ := http.NewRequest("PUT", url, strings.NewReader(data.Encode()))
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("Content-Length", strconv.Itoa(len(data.Encode())))
	req.Header.Add("accessToken", accessToken)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	r, _ := ioutil.ReadAll(resp.Body)
	if gin.Mode() == gin.DebugMode {
		fmt.Println("[debug] iShopFloor API response: " + string(r))
	}
	if err = json.Unmarshal(r, &jsonResp); err != nil {
		fmt.Println("[debug] iShopFloor api respond with an error!")
		return err
	}
	if resp.StatusCode >= 400 || jsonResp["success"] != true {
		return errors.New("failed to update job status with iShopFloor API")
	}

	return nil
}
