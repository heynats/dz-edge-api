package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
	"thingularity.co/dz-edge-api/api"
	"thingularity.co/dz-edge-api/models"
)

func main() {
	asciiArt := `
	 ______   _____    ______   ______   ______   ______  __    
	/\  ___\ /\  __-. /\  ___\ /\  ___\ /\  __ \ /\  == \/\ \   
	\ \  __\ \ \ \/\ \\ \ \__ \\ \  __\ \ \  __ \\ \  _-/\ \ \  
	 \ \_____\\ \____- \ \_____\\ \_____\\ \_\ \_\\ \_\   \ \_\ 
	  \/_____/ \/____/  \/_____/ \/_____/ \/_/\/_/ \/_/    \/_/  v1.1.0 by Stanley Yeh, 2020
	`
	fmt.Println(asciiArt + "\n")

	gin.DisableConsoleColor()
	f, _ := os.Create("gin.log")
	gin.DefaultWriter = io.MultiWriter(f, os.Stdout)

	router := gin.Default()

	v1 := router.Group("/api/v1")
	{
		v1.GET("/", func(c *gin.Context) {
			time.Sleep(5 * time.Second)
			c.String(http.StatusOK, "DZ Edge API Server")
		})
		v1.POST("/job/create", api.CreateJob)
		v1.POST("/job/update", api.UpdateJob)
		v1.GET("/job/:jobId/:processType", api.GetJob)    // get job details
		v1.GET("/process/:operation", api.GetProcMetrics) // get process status details
		v1.DELETE("/job/all", api.DeleteAllJobs)
		v1.DELETE("/job/pre", api.DeletePreJobs)           // delete all preprocess jobs
		v1.DELETE("/job/pre/:jobId", api.DeletePreJobByID) // delete preprocess job by ID
		v1.DELETE("/job/gal", api.DeleteGalJobs)           // delete all galvanization jobs
		v1.DELETE("/job/gal/:jobId", api.DeleteGalJobByID) // delete galvanization job by ID
	}

	srv := &http.Server{
		Addr:    ":8080",
		Handler: router,
	}

	// Initializing the server in a goroutine so that it won't block the graceful shutdown handling below
	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	// Initializing job queues
	deviceDisconnected := true
	var err error
	var waitDuration int
	var waitFactor = 5
	var n = 1
	fmt.Print("\n")
	log.Println("Initializing process queues...")
	for deviceDisconnected {
		if waitDuration < 300 {
			waitDuration = waitFactor * n
		}
		if err = models.InitQueues(); err != nil {
			log.Println(err.Error())
			log.Println("Retry in " + strconv.Itoa(waitDuration) + " seconds...")
			time.Sleep(time.Duration(waitDuration) * time.Second)
		} else {
			deviceDisconnected = false
		}
		n++
	}
	fmt.Print("done!\n\n")

	// Wait for interrupt signal to gracefully shutdown the server with 5 sec. timeout
	quit := make(chan os.Signal)
	// kill (no param) default send syscall.SIGTERM
	// kill -2 is syscall.SIGINT
	// kill -9 is syscall.SIGKILL but can't be catch, so ignore this one
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Shutting down server...")

	// The context is used to inform the server that it has 5 seconds to finish the
	// request it is currently handling
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server forced to shutdown:", err)
	}

	log.Println("Server exiting")
}
