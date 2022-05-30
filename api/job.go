package api

import (
	"go-reserve/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func ListJob(c *gin.Context) {
	jobes, err := GE.JobManage.List()
	if err != nil {
		c.Errors = append(c.Errors, &gin.Error{Err: err})
		return
	}
	c.JSON(http.StatusOK, jobes)
	return
}

func CreateJob(c *gin.Context) {
	var job models.Job

	if err := c.ShouldBind(&job); err != nil {
		c.Errors = append(c.Errors, &gin.Error{Err: err})
		return
	}

	if err := GE.JobManage.Create(&job); err != nil {
		c.Errors = append(c.Errors, &gin.Error{Err: err})
		return
	}
}

func GetJob(c *gin.Context) {
	id := c.Param("id")
	job, err := GE.JobManage.Get(id)
	if err != nil {
		c.Errors = append(c.Errors, &gin.Error{Err: err})
		return
	}

	c.JSON(http.StatusOK, job)
	return
}

func UpdateJob(c *gin.Context) {
	var job *models.Job
	id := c.Param("id")
	if err := GE.JobManage.Update(id, job); err != nil {
		c.Errors = append(c.Errors, &gin.Error{Err: err})
		return
	}
}

func DeleteJob(c *gin.Context) {
	id := c.Param("id")
	if err := GE.JobManage.Delete(id); err != nil {
		c.Errors = append(c.Errors, &gin.Error{Err: err})
		return
	}
}
