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

// Route01Metrics defines 前處理量空桶 key-values
type Route01Metrics struct {
	EmptyBucketWeight string
}

// Route02Metrics defines 前處理量滿桶 key-values
type Route02Metrics struct {
	FullBucketWeight string
}

// Route03Metrics defines 鍍鋅前處理 key-values
type Route03Metrics struct {
	PicklingSmallTime   string
	DegreasingSmallTime string
	HotWaterTime        string
	DegreasingTime      string
	DegreasingWaterTime string
	FluxTime            string
	DryingTime          string
	PicklingTime        string
	PicklingWaterTime   string
	LessFluxTime        string
	TurnOverTime        string
}

// Route04Metrics defines 鍍鋅 key-values
type Route04Metrics struct {
	SubBasketWeight string
	SubBasketTime   string
	SubBasketTemp   string
}

// Route05Metrics defines 鍍鋅量空桶 key-values
type Route05Metrics struct {
	EmptyBucketWeight string
}

// Route06Metrics defines 鍍鋅量滿桶 key-values
type Route06Metrics struct {
	FullBucketWeight string
}

// Route07Metrics defines 鍍鋅量良品 key-values
type Route07Metrics struct {
	ProductWeight string
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

// UpdateRouteMetrics ...
func UpdateRouteMetrics(jobID string, routeNo string, metricsJSON string) error {
	var jsonResp map[string]interface{}
	accessToken, err := getAccessToken()
	if err != nil {
		return err
	}

	data := url.Values{}
	switch routeNo {
	case "1":
		m := Route01Metrics{}
		e := json.Unmarshal([]byte(metricsJSON), &m)
		if e != nil {
			return e
		}
		data.Set("emptyBucketWeight", m.EmptyBucketWeight)
		data.Set("status", "3")
		data.Set("endTime", strconv.FormatInt(time.Now().Unix(), 10))
	case "2":
		m := Route02Metrics{}
		e := json.Unmarshal([]byte(metricsJSON), &m)
		if e != nil {
			return e
		}
		data.Set("fullBucketWeight", m.FullBucketWeight)
		data.Set("status", "3")
		data.Set("endTime", strconv.FormatInt(time.Now().Unix(), 10))
	case "3":
		m := Route03Metrics{}
		e := json.Unmarshal([]byte(metricsJSON), &m)
		if e != nil {
			return e
		}
		data.Set("picklingSmallTime", m.PicklingSmallTime)
		data.Set("degreasingSmallTime", m.DegreasingSmallTime)
		data.Set("hotWaterTime", m.HotWaterTime)
		data.Set("degreasingTime", m.DegreasingTime)
		data.Set("degreasingWaterTime", m.DegreasingWaterTime)
		data.Set("fluxTime", m.FluxTime)
		data.Set("dryingTime", m.DryingTime)
		data.Set("picklingTime", m.PicklingTime)
		data.Set("picklingWaterTime", m.PicklingWaterTime)
		data.Set("lessFluxTime", m.LessFluxTime)
		data.Set("turnOverTime", m.TurnOverTime)
		data.Set("status", "3")
		data.Set("endTime", strconv.FormatInt(time.Now().Unix(), 10))
	case "4":
		m := Route04Metrics{}
		e := json.Unmarshal([]byte(metricsJSON), &m)
		if e != nil {
			return e
		}
		data.Set("subBasketWeight", m.SubBasketWeight)
		data.Set("subBasketTime", m.SubBasketTime)
		data.Set("subBasketTemp", m.SubBasketTemp)
		data.Set("status", "3")
		data.Set("endTime", strconv.FormatInt(time.Now().Unix(), 10))
	case "5":
		m := Route05Metrics{}
		e := json.Unmarshal([]byte(metricsJSON), &m)
		if e != nil {
			return e
		}
		data.Set("emptyBucketWeight", m.EmptyBucketWeight)
		data.Set("status", "3")
		data.Set("endTime", strconv.FormatInt(time.Now().Unix(), 10))
	case "6":
		m := Route06Metrics{}
		e := json.Unmarshal([]byte(metricsJSON), &m)
		if e != nil {
			return e
		}
		data.Set("fullBucketWeight", m.FullBucketWeight)
		data.Set("status", "3")
		data.Set("endTime", strconv.FormatInt(time.Now().Unix(), 10))
	case "7":
		m := Route07Metrics{}
		e := json.Unmarshal([]byte(metricsJSON), &m)
		if e != nil {
			return e
		}
		data.Set("productWeight", m.ProductWeight)
		data.Set("status", "3")
		data.Set("endTime", strconv.FormatInt(time.Now().Unix(), 10))
	default:
		return errors.New("Unknown process: route #" + routeNo)
	}
	fmt.Printf("[EdgeAPI] Request body: %v\n", data)

	url := apiBaseURL + "/1/app/work_order/" + jobID + "/route/" + routeNo + "/parameter"
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
		errMsg := jsonResp["errors"].([]interface{})
		return errors.New("failed to update metrics fields with iShopFloor API: " + errMsg[0].(string))
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
