package main

import (
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

type NewsReader struct {
	ID           int            `json:"id"`
	Name         string         `json:"name"`
	Email        string         `json:"email"`
	PhoneNumber  string         `json:"phone_number"`
	Designation  string         `json:"designation"`
	WorkSchedule []WorkSchedule `json:"work_schedule"`
}

type WorkSchedule struct {
	NewsPrgName string `json:"news_prg_name"`
	Date        string `json:"date"`
	StartTime   string `json:"start_time"`
	EndTime     string `json:"end_time"`
}

var newsReaders []NewsReader

func main() {
	r := gin.Default()

	// Add news reader details
	r.POST("/newsreader", addNewsReader)

	// Update news reader details
	r.PUT("/newsreader/:id", updateNewsReader)

	// Delete news reader details
	r.DELETE("/newsreader/:id", deleteNewsReader)

	// View all news reader details
	r.GET("/newsreaders", getAllNewsReaders)

	// Add work schedule details of a news reader
	r.POST("/newsreader/:id/workschedule", addWorkSchedule)

	// Display the total hours of work for a news reader
	r.GET("/newsreader/:id/totalhours", getTotalHoursOfWork)

	r.Run()
}

// Add news reader details
func addNewsReader(c *gin.Context) {
	var newsReader NewsReader
	if err := c.ShouldBindJSON(&newsReader); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Add news reader to the slice
	newsReaders = append(newsReaders, newsReader)
	c.JSON(http.StatusCreated, gin.H{"message": "News reader added successfully"})
}

// Update news reader details
func updateNewsReader(c *gin.Context) {
	id := c.Param("id")
	// Find the news reader by ID
	for i, newsReader := range newsReaders {
		if strconv.Itoa(newsReader.ID) == id {
			var updatedNewsReader NewsReader
			if err := c.ShouldBindJSON(&updatedNewsReader); err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
				return
			}
			// Update the news reader
			newsReaders[i] = updatedNewsReader
			c.JSON(http.StatusOK, gin.H{"message": "News reader updated successfully"})
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"message": "News reader not found"})
}

// Delete news reader details
func deleteNewsReader(c *gin.Context) {
	id := c.Param("id")
	// Find the news reader by ID and remove it from the slice
	for i, newsReader := range newsReaders {
		if strconv.Itoa(newsReader.ID) == id {
			newsReaders = append(newsReaders[:i], newsReaders[i+1:]...)
			c.JSON(http.StatusOK, gin.H{"message": "News reader deleted successfully"})
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"message": "News reader not found"})
}

// View all news reader details
func getAllNewsReaders(c *gin.Context) {
	c.JSON(http.StatusOK, newsReaders)
}

// Add work schedule details of a news reader
func addWorkSchedule(c *gin.Context) {
	id := c.Param("id")
	var workSchedule WorkSchedule
	if err := c.ShouldBindJSON(&workSchedule); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// Find the news reader by ID and add the work schedule
	for i, newsReader := range newsReaders {
		if strconv.Itoa(newsReader.ID) == id {
			newsReaders[i].WorkSchedule = append(newsReaders[i].WorkSchedule, workSchedule)
			c.JSON(http.StatusCreated, gin.H{"message": "Work schedule added successfully"})
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"message": "News reader not found"})
}

// Display the total hours of work for a news reader
func getTotalHoursOfWork(c *gin.Context) {
	id := c.Param("id")
	// Find the news reader by ID and calculate the total hours of work
	for _, newsReader := range newsReaders {
		if strconv.Itoa(newsReader.ID) == id {
			totalHours := 0
			for _, schedule := range newsReader.WorkSchedule {
				// Assuming start and end times are in "HH:MM" format
				startTime, _ := time.Parse("15:04", schedule.StartTime)
				endTime, _ := time.Parse("15:04", schedule.EndTime)
				totalHours += int(endTime.Sub(startTime).Hours())
			}
			c.JSON(http.StatusOK, gin.H{"total_hours": totalHours})
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"message": "News reader not found"})
}
