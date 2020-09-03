package services

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

const waUser = "admin"
const waPass = "qwER123!"

//const waAPIBaseURL = "https://tg-webaccess.dotzero.tech/wawebservice/json"
const waAPIBaseURL = "http://127.0.0.1/wawebservice/json"

type waResponse struct {
	Result waResponseResult  `json:"Result"`
	Values []waResponseValue `json:"Values"`
}

type waResponseResult struct {
	Ret   int `json:"Ret"`
	Total int `json:"Total"`
}

type waResponseValue struct {
	Name    string          `json:"Name"`
	Value   json.RawMessage `json:"Value"`
	Quality int             `json:"Quality"`
}

// GetDeviceStatus returns true if device is connected, false otherwise
func GetDeviceStatus(project string, node string, port string) bool {
	var j map[string]interface{}

	client := &http.Client{}
	req, _ := http.NewRequest("GET", waAPIBaseURL+"/GetDeviceStatus/"+project+"/"+node+"/"+port, bytes.NewBuffer([]byte{}))
	req.SetBasicAuth(waUser, waPass)
	req.Header.Set("Content-Type", "application/json; charset=utf-8")
	resp, err := client.Do(req)
	if err != nil {
		return false
	}
	defer resp.Body.Close()
	respBody, _ := ioutil.ReadAll(resp.Body)

	err = json.Unmarshal(respBody, &j)
	if err != nil {
		return false
	}
	if j["Result"].(map[string]interface{})["Total"] != "0" {
		return true
	}

	return false
}

// GetTagValue returns the 'Value' of the WebAccess response as an 'float'
func GetTagValue(project string, tags []string) ([]float64, error) {
	// build request string
	var reqStr = `{"Tags":[`
	for _, s := range tags {
		reqStr += `{"Name":"` + s + `"},`
	}
	reqStr = strings.TrimSuffix(reqStr, ",")
	reqStr += `]}`

	var reqBody = []byte(reqStr)
	var waResp waResponse
	var val []float64

	// Create request & execute
	client := &http.Client{}
	req, _ := http.NewRequest("POST", waAPIBaseURL+"/GetTagValue/"+project, bytes.NewBuffer(reqBody))
	req.SetBasicAuth(waUser, waPass)
	req.Header.Set("Content-Type", "application/json; charset=utf-8")
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	respBody, _ := ioutil.ReadAll(resp.Body)

	// Process API response
	if err = json.Unmarshal(respBody, &waResp); err != nil {
		return nil, err
	}
	for _, v := range waResp.Values {
		if v.Quality != 0 {
			return nil, errors.New("unable to read value, PLC might be offline")
		}
		//i, err := strconv.Atoi(string(v.Value))
		i, err := strconv.ParseFloat(string(v.Value), 64)
		if err != nil {
			return nil, err
		}
		val = append(val, i)
	}
	if gin.Mode() == gin.DebugMode {
		fmt.Printf("[debug] wa tag values: %v\n", val)
	}
	return val, nil
}

// GetTextTagValue ...
func GetTextTagValue(project string, tags []string) ([]string, error) {
	// build request string
	var reqStr = `{"Tags":[`
	for _, s := range tags {
		reqStr += `{"Name":"` + s + `"},`
	}
	reqStr = strings.TrimSuffix(reqStr, ",")
	reqStr += `]}`

	var reqBody = []byte(reqStr)
	var waResp waResponse
	var val []string

	// Create request & execute
	client := &http.Client{}
	req, _ := http.NewRequest("POST", waAPIBaseURL+"/GetTagValueText/"+project, bytes.NewBuffer(reqBody))
	req.SetBasicAuth(waUser, waPass)
	req.Header.Set("Content-Type", "application/json; charset=utf-8")
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	respBody, _ := ioutil.ReadAll(resp.Body)

	// Process API response
	if err = json.Unmarshal(respBody, &waResp); err != nil {
		return nil, err
	}
	for _, v := range waResp.Values {
		if v.Quality != 0 {
			return nil, errors.New("unable to read value, PLC might be offline")
		}
		s := string(v.Value[1 : len(v.Value)-1])
		val = append(val, s)
	}
	if gin.Mode() == gin.DebugMode {
		fmt.Printf("[debug] wa tag values: %v\n", val)
	}
	return val, nil
}

// SetTagValue ...
func SetTagValue(project string, tag string, value int) error {
	var reqBody = []byte(`{"Tags":[{"Name":"` + tag + `","Value":` + strconv.Itoa(value) + `}]}`)
	var waResp map[string]interface{}

	// Create request & execute
	client := &http.Client{}
	req, _ := http.NewRequest("POST", waAPIBaseURL+"/SetTagValue/"+project, bytes.NewBuffer(reqBody))
	req.SetBasicAuth(waUser, waPass)
	req.Header.Set("Content-Type", "application/json; charset=utf-8")
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Request ---> " + string(reqBody))
		return err
	}
	defer resp.Body.Close()
	respBody, _ := ioutil.ReadAll(resp.Body)

	// Process API response
	if err = json.Unmarshal(respBody, &waResp); err != nil {
		fmt.Println("Response <--- " + string(respBody))
		return err
	}
	switch waResp["Ret"].(float64) {
	case 0:
		return nil
	default:
		return errors.New("Unable to write value to WebAccess tag using SetTagValue API")
	}
}

// SetTextTagValue ...
func SetTextTagValue(project string, tag string, value string) error {
	var reqBody = []byte(`{"Tags":[{"Name":"` + tag + `","Value":"` + value + `"}]}`)
	var waResp map[string]interface{}

	// Create request & execute
	client := &http.Client{}
	req, _ := http.NewRequest("POST", waAPIBaseURL+"/SetTagValueText/"+project, bytes.NewBuffer(reqBody))
	req.SetBasicAuth(waUser, waPass)
	req.Header.Set("Content-Type", "application/json; charset=utf-8")
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Request ---> " + string(reqBody))
		return err
	}
	defer resp.Body.Close()
	respBody, _ := ioutil.ReadAll(resp.Body)

	// Process API response
	if err = json.Unmarshal(respBody, &waResp); err != nil {
		fmt.Println("Response <--- " + string(respBody))
		return err
	}
	switch waResp["Ret"].(float64) {
	case 0:
		return nil
	default:
		return errors.New("Unable to write value to WebAccess tag using SetTagValueText API")
	}
}
