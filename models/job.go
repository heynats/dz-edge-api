package models

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"thingularity.co/dz-edge-api/services"
)

// Job represents a child job object of the iShopFloor system
type Job struct {
	WoID          string    `json:"woId"`
	Group         string    `json:"group"`
	PartNumber    string    `json:"partNumber"`
	ProdName      string    `json:"prodName"`
	Qty           int       `json:"qty"`
	Route         []Process `json:"route"`
	SurfacePhoto  []string  `json:"surfacePhoto"`
	ProductPhoto  []string  `json:"productPhoto"`
	EstimatedTime time.Time `json:"estimatedTime"`
	IsASAP        bool      `json:"isASAP"`
	StartTime     time.Time `json:"startTime"`
	EndTime       time.Time `json:"endTime"`
	Deadline      time.Time `json:"deadline"`
	Status        int       `json:"status"` // 製程狀態，0: 尚未開始、1: 進行中、2: 暫停、3: 已完成、4: 已完工，但未完成
	Yield         int       `json:"yield"`
	Remarks       string    `json:"remarks"`
	IsActive      bool      `json:"isActive"`
	UpdateTime    time.Time `json:"updateTime"`
	CreateTime    time.Time `json:"createTime"`
}

// Process represents the process information embedded in the 'Route' field of the SubJob struct
type Process struct {
	No             int            `json:"no"`           // 製程順序
	Number         string         `json:"number"`       // 製程編號
	Name           string         `json:"name"`         // 作業說明
	RoutingName    string         `json:"routingName"`  // 已棄用
	Status         int            `json:"status"`       // 製程狀態，0: 尚未開始、1: 進行中、2: 暫停、3: 已完成、4: 已完工，但未完成
	StaID          string         `json:"staID"`        // 工站編號
	DevID          string         `json:"devID"`        // 設備編號
	Mid            string         `json:"mid"`          // 工站 ID，可修改
	StartTime      time.Time      `json:"startTime"`    // 開始時間
	EndTime        time.Time      `json:"endTime"`      // 結束時間
	WorkHour       []string       `json:"workHour"`     // 異常工時、除外工時
	WorkNumber     string         `json:"workNumber"`   // 操作員工號，可修改
	WorkName       string         `json:"workName"`     // 操作員姓名，可修改
	StdTP          int            `json:"stdTP"`        // 標準前置作業時數(min)
	StdTS          int            `json:"stdTS"`        // 標準單位加工時數(min)
	StdWorkTime    int            `json:"stdWorkTime"`  // 標準作業時數(min)
	ActWorkTime    int            `json:"actWorkTime"`  // 實際作業時數(min)
	TotalTS        int            `json:"totalTS"`      // 實際加工時數(min) 可先空
	Qty            int            `json:"qty"`          // 實際數量 (API 回傳)
	ReturnAmount   int            `json:"returnAmount"` // 現場回報量 (人工填寫)
	Good           int            `json:"good"`         // 製程個別良品數量
	Defect         int            `json:"defect"`       // 製程個別不良品數量
	Experiment     int            `json:"experiment"`   // 製程個別試車數量
	Yield          int            `json:"yield"`        // 製程個別良率
	DefectRate     int            `json:"defectRate"`   // 製程個別不良率
	FaultyReason   string         `json:"faultyReason"` // 不良品原因
	IsAbnStart     bool           `json:"isAbnStart"`   // 是否為除外開始
	IsAbnEnd       bool           `json:"isAbnEnd"`     // 是否為除外結束
	IsExcStart     bool           `json:"isExcStart"`   // 是否為異常開始
	IsExcEnd       bool           `json:"isExcEnd"`     // 是否為異常結束
	Remarks        string         `json:"remarks"`      // 備註
	IsShift        bool           `json:"isShift"`      // 是否需要交班
	StartCount     int            `json:"startCount"`   // 機器的開工的製造數量
	EndCount       int            `json:"endCount"`     // 機器的完工的製造數量
	MeasuringPoint []Measurements `json:"measuringPoint"`
}

// Measurements represents the measurements embedded in the 'MeasuringPoints' field of the Process struct
type Measurements struct {
	Key         string `json:"key"`         // 鍵值名稱
	Name        string `json:"name"`        // 名稱
	Description string `json:"description"` // 說明
	Unit        string `json:"unit"`        // 單位
	Mode        int    `json:"mode"`        // 模式
	Value       []int  `json:"value"`       // 測量值
	Length      int    `json:"length"`      // 測量值數量
	Max         int    `json:"max"`         // 上限
	Standard    int    `json:"standard"`    // 標準值
	Min         int    `json:"min"`         // 下限
}

type jobQueueItem struct {
	jobID  string
	status int
}

// preprocQueue represents the buffer registers for the "pre-treatment" process
var preprocQueue = []jobQueueItem{
	{"", 0}, {"", 0}, {"", 0}, {"", 0}, {"", 0}}

// galvanizeQueue represents the buffer registers for the "galvanization" process
var galvanizeQueue = []jobQueueItem{{"", 0}, {"", 0}}

// InitQueues initializes the process queues by sync'ing with the PLC upon system start
func InitQueues() error {
	if !services.GetDeviceStatus("dotzero", "tgnode01", "1") {
		return errors.New("PLC is disconnected")
	}

	wID, err := services.GetTextTagValue("dotzero", WIDTags)
	if err != nil {
		return err
	}
	if len(wID) == 0 {
		return errors.New("WebAccess GetTextTagValue returns empty response")
	}
	wStat, err := services.GetTagValue("dotzero", WStatTags)
	if err != nil {
		return err
	}
	if len(wStat) == 0 {
		return errors.New("WebAccess GetTagValue returns empty response")
	}
	gID, err := services.GetTextTagValue("dotzero", GIDTags)
	if err != nil {
		return err
	}
	if len(gID) == 0 {
		return errors.New("WebAccess GetTextTagValue returns empty response")
	}
	gStat, err := services.GetTagValue("dotzero", GStatTags)
	if err != nil {
		return err
	}
	if len(gStat) == 0 {
		return errors.New("WebAccess GetTagValue returns empty response")
	}

	for i, item := range preprocQueue {
		if wID[i] != "3030" && wID[i] != "" {
			item.jobID = wID[i]
			item.status = int(wStat[i])
			preprocQueue[i] = item
		}
	}
	printQueue("pre")

	for i, item := range galvanizeQueue {
		if gID[i] != "3030" && gID[i] != "" {
			item.jobID = gID[i]
			item.status = int(gStat[i])
			galvanizeQueue[i] = item
		}
	}
	printQueue("gal")

	return nil
}

func printQueue(procType string) {
	var str strings.Builder
	str.WriteString("[")
	if procType == "pre" {
		for _, item := range preprocQueue {
			status := item.status
			if status != 0 {
				str.WriteString(" (" + item.jobID + "," + strconv.Itoa(status) + ")")
			}
		}
		str.WriteString(" ]")
		fmt.Println("PQueue" + str.String())
	} else {
		for _, item := range galvanizeQueue {
			status := item.status
			if status != 0 {
				str.WriteString(" (" + item.jobID + "," + strconv.Itoa(status) + ")")
			}
		}
		str.WriteString(" ]")
		fmt.Println("GQueue" + str.String())
	}
}

// JobExists ...
func JobExists(procType string, id string) int {
	if procType == "pre" {
		for i, item := range preprocQueue {
			if item.jobID == id {
				return i
			}
		}
		return -1
	}
	for i, item := range galvanizeQueue {
		if item.jobID == id {
			return i
		}
	}
	return -1
}

// AddJob adds a job of the specified process type ('pre' or 'gal') into the corresponding job queue
func AddJob(procType string, id string, status int) (int, error) {
	if procType == "pre" {
		for i, item := range preprocQueue {
			if item.jobID == id && item.status != 0 {
				//return -1, errors.New("job exists in pre-treatment process queue")
				return i, nil
			}
			if item.status == 0 {
				item.status = status
				item.jobID = id
				preprocQueue[i] = item
				return i, nil
			}
		}
		printQueue(procType)
		return -1, errors.New("pre-treatment process queue is full")
	}
	for i, item := range galvanizeQueue {
		if item.jobID == id && item.status != 0 {
			//return -1, errors.New("job exists in galvanizing process queue")
			return i, nil
		}
		if item.status == 0 {
			item.status = status
			item.jobID = id
			galvanizeQueue[i] = item
			return i, nil
		}
	}
	printQueue(procType)
	return -1, errors.New("galvanize process queue is full")
}

// RemoveJob ...
func RemoveJob(procType string, index int) {
	if procType == "pre" {
		preprocQueue[index] = jobQueueItem{"", 0}
	} else {
		galvanizeQueue[index] = jobQueueItem{"", 0}
	}
}

// RemoveAllJobs ...
func RemoveAllJobs() {
	preprocQueue = []jobQueueItem{
		{"", 0}, {"", 0}, {"", 0}, {"", 0}, {"", 0},
		{"", 0}, {"", 0}, {"", 0}, {"", 0}, {"", 0},
		{"", 0}, {"", 0}, {"", 0}, {"", 0}, {"", 0},
		{"", 0}, {"", 0}, {"", 0}, {"", 0}, {"", 0}}
	galvanizeQueue = []jobQueueItem{{"", 0}, {"", 0}}
	printQueue("pre")
	printQueue("gal")
}

// RemovePreJobs ...
func RemovePreJobs() {
	preprocQueue = []jobQueueItem{
		{"", 0}, {"", 0}, {"", 0}, {"", 0}, {"", 0},
		{"", 0}, {"", 0}, {"", 0}, {"", 0}, {"", 0},
		{"", 0}, {"", 0}, {"", 0}, {"", 0}, {"", 0},
		{"", 0}, {"", 0}, {"", 0}, {"", 0}, {"", 0}}
	printQueue("pre")
}

// RemoveGalJobs ...
func RemoveGalJobs() {
	galvanizeQueue = []jobQueueItem{{"", 0}, {"", 0}}
	printQueue("gal")
}

// UpdateJob updates a job of the specified process type ('pre' or 'gal') and jobID in the corresponding queue
func UpdateJob(procType string, id string, status int) (int, error) {
	if procType == "pre" {
		for i, item := range preprocQueue {
			if item.jobID == id {
				item.status = status
				preprocQueue[i] = item
				if gin.Mode() == gin.DebugMode {
					printQueue(procType)
				}
				return i, nil
			}
		}
		return -1, errors.New("jobID not found in pre-treatment process queue")
	}
	for i, item := range galvanizeQueue {
		if item.jobID == id {
			item.status = status
			if status == 0 {
				item.jobID = ""
			}
			galvanizeQueue[i] = item
			if gin.Mode() == gin.DebugMode {
				printQueue(procType)
			}
			return i, nil
		}
	}
	return -1, errors.New("jobID not found in galvanize process queue")
}
