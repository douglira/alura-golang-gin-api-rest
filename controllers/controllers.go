package controllers

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/douglira/alura-golang-gin-api-rest/database"
	"github.com/douglira/alura-golang-gin-api-rest/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func AllStudents(c *gin.Context) {
	var students []models.Student
	database.DB.Find(&students)
	c.JSON(200, students)
}

func RegisterStudent(c *gin.Context) {
	var student models.Student
	if err := c.ShouldBindJSON(&student); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	result := database.DB.Where(&student, "IdentityNumber").Find(&student)
	if result.RowsAffected > 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid identityNumber",
		})
		return
	}
	database.DB.Create(&student)
	c.Writer.WriteHeader(http.StatusCreated)
}

func FindStudentById(c *gin.Context) {
	studentId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid parameter",
		})
		return
	}
	var student models.Student
	result := database.DB.Find(&student, studentId)
	if result.RowsAffected == 0 {
		c.Writer.WriteHeader(http.StatusNotFound)
		return
	}
	c.JSON(http.StatusOK, student)
}

func DeleteStudent(c *gin.Context) {
	studentId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid parameter",
		})
		return
	}
	result := database.DB.Delete(&models.Student{}, studentId)
	if result.RowsAffected != 1 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid parameter",
		})
		return
	}
	c.Writer.WriteHeader(http.StatusNoContent)
}

func UpdateStudent(c *gin.Context) {
	var student models.Student
	studentId := c.Param("id")
	database.DB.First(&student, studentId)

	if err := c.ShouldBindJSON(&student); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	result := database.DB.Clauses(clause.Returning{}).Model(&student).UpdateColumns(student)
	if result.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}
	c.JSON(http.StatusOK, student)
}

func FindStudentByIdentityNumber(c *gin.Context) {
	identityNumber := c.Query("identityNumber")
	var student models.Student
	if result := database.DB.Take(&student, "identity_number = ?", identityNumber); errors.Is(result.Error, gorm.ErrRecordNotFound) {
		c.Writer.WriteHeader(http.StatusNoContent)
		return
	}
	c.JSON(http.StatusOK, student)
}
