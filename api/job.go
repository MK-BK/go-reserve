package api

import (
	"go-reserve/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func listJobs(c *gin.Context) {
	jobs, err := GE.JobManager.List()
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	c.JSON(http.StatusOK, jobs)
	return
}

func createJob(c *gin.Context) {
	var job models.Job

	if err := c.ShouldBind(&job); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	if err := GE.JobManager.Create(&job); err != nil {
		c.Errors = append(c.Errors, &gin.Error{Err: err})
		return
	}
}

func getJob(c *gin.Context) {
	id := c.Param("id")
	job, err := GE.JobManager.Get(id)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, job)
	return
}

func updateJob(c *gin.Context) {
	var job *models.Job
	id := c.Param("id")
	if err := GE.JobManager.Update(id, job); err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}
}

func deleteJob(c *gin.Context) {
	id := c.Param("id")
	if err := GE.JobManager.Delete(id); err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}
}
