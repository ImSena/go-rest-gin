package controllers

import (
	"errors"
	"go-rest-gin/database"
	"go-rest-gin/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func AllStudents(c *gin.Context) {
	var students []models.Aluno
	database.DB.Find(&students)
	c.JSON(http.StatusOK, students)
}

func GetById(c *gin.Context) {
	var aluno models.Aluno
	id := c.Params.ByName("id")

	result := database.DB.First(&aluno, id)

	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, gin.H{
				"error": "Aluno não encontrado",
			})
			return
		}

		c.JSON(http.StatusInternalServerError, gin.H{
			"error": result.Error.Error(),
		})

		return
	}
	c.JSON(http.StatusOK, aluno)
}

func CreateStudents(c *gin.Context) {
	var students models.Aluno
	if err := c.ShouldBindJSON(&students); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	if err := models.ValidateStudent(&students); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	database.DB.Create(&students)
	c.JSON(http.StatusOK, students)
}

func DeleteStudents(c *gin.Context) {
	var students models.Aluno

	id := c.Params.ByName("id")

	result := database.DB.First(&students, id)

	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, gin.H{
				"error": "Student Not found",
			})
			return
		}

		c.JSON(http.StatusInternalServerError, gin.H{
			"error": result.Error.Error(),
		})

		return
	}

	database.DB.Delete(&students, id)

	c.JSON(http.StatusOK, gin.H{
		"data": gin.H{
			"id":   students.ID,
			"name": students.Nome,
		},
	})
}

func EditStudent(c *gin.Context) {
	var student models.Aluno

	id := c.Params.ByName("id")

	result := database.DB.First(&student, id)

	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, gin.H{
				"error": "Student Not found",
			})
			return
		}

		c.JSON(http.StatusInternalServerError, gin.H{
			"error": result.Error.Error(),
		})

		return
	}

	if err := c.ShouldBindJSON(&student); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	if err := models.ValidateStudent(&student); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	database.DB.Model(&student).UpdateColumns(student)

	c.JSON(http.StatusOK, student)

}

func GetByCPF(c *gin.Context) {
	var student models.Aluno

	cpf := c.Param("cpf")

	result := database.DB.Where(&models.Aluno{CPF: cpf}).First(&student)

	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, gin.H{
				"error": "Aluno não encontrado",
			})
			return
		}

		c.JSON(http.StatusInternalServerError, gin.H{
			"error": result.Error.Error(),
		})

		return
	}

	c.JSON(http.StatusOK, student)
}

func ShowPageIndex(c *gin.Context) {
	var alunos []models.Aluno
	database.DB.Find(&alunos)
	c.HTML(http.StatusOK, "index.html", gin.H{
		"alunos": alunos,
	})
}

func RouteNotFound(c *gin.Context) {
	c.HTML(http.StatusNotFound, "404.html", nil)
}
