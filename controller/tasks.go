package controller

import (
	"net/http"

	"github.com/SanketPatil6869/Task_Management_Gin1/config"
	"github.com/SanketPatil6869/Task_Management_Gin1/models"
	"github.com/gin-gonic/gin"
)

func CreateTask(c *gin.Context) {
	var task models.Task
	if err := c.ShouldBindJSON(&task); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	config.DB.Create(&task)
	c.JSON(http.StatusOK, gin.H{"Data": task})
}

func FindTasks(c *gin.Context) {
	var tasks []models.Task
	config.DB.Find(&tasks)
	c.JSON(http.StatusOK, gin.H{"Data": tasks})
}

func FindTask(c *gin.Context) {
	var task models.Task
	if err := config.DB.Where("id = ?", c.Param("id")).First(&task).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record Not Found!"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"Data": task})
}

func FindTaskStatus(c *gin.Context) {
	var tasks []models.Task
	if err := config.DB.Where("status = ?", c.Param("status")).Find(&tasks).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record Not Found!"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"Data": tasks})
}

func UpdateTask(c *gin.Context) {
	var task models.Task
	if err := config.DB.Where("id = ?", c.Param("id")).First(&task).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record Not Found!"})
		return
	}
	var taskUpdated models.Task
	if err := c.ShouldBindJSON(&taskUpdated); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	config.DB.Model(&task).Updates(taskUpdated)
	c.JSON(http.StatusOK, gin.H{"Data": task})
}

func DeleteTask(c *gin.Context) {
	var task models.Task
	if err := config.DB.Where("id =?", c.Param("id")).First(&task).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record Not Found!"})
		return
	}
	config.DB.Delete(&task)
	c.JSON(http.StatusOK, gin.H{"Data": true})
}

// func FindTasks_2(c *gin.Context) {
// 	rows, err := config.DB2.Query("SELECT * FROM tasks")
// 	if err != nil {
// 		//return nil, err
// 		panic(err)
// 	}
// 	defer rows.Close()

// 	task := make([]models.Task, 0)

// 	for rows.Next() {
// 		singleTask := models.Task{}
// 		err = rows.Scan(&singleTask.ID, &singleTask.Title, &singleTask.Description, &singleTask.Description, &singleTask.Due_Date, &singleTask.Status)

// 		if err != nil {
// 			//return nil, err
// 			panic("Error 2")
// 		}
// 		task = append(task, singleTask)
// 	}

// 	err = rows.Err()

// 	if err != nil {
// 		//return nil, err
// 		panic("Error 3")
// 	}

// 	//return book, err
// 	c.JSON(http.StatusOK, gin.H{"data": task})
// }
