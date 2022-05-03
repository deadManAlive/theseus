package controller

import (
	"net/http"

	"goctr/structs"

	"github.com/gin-gonic/gin"
)

func (ind *InDB) GetPerson(c *gin.Context) {
	var (
		person structs.Person
		result gin.H
	)

	id := c.Param("id")
	err := ind.DB.Where("id = ?", id).First(&person).Error

	if err != nil {
		result = gin.H{
			"result": err.Error(),
			"count":  0,
		}
	} else {
		result = gin.H{
			"result": person,
			"count":  1,
		}
	}

	c.JSON(http.StatusOK, result)

}

func (ind *InDB) GetPersons(c *gin.Context) {
	var (
		persons []structs.Person
		result  gin.H
	)
	ind.DB.Find(&persons)
	if len(persons) <= 0 {
		result = gin.H{
			"result": nil,
			"count":  0,
		}
	} else {
		result = gin.H{
			"result": persons,
			"count":  len(persons),
		}
	}

	c.JSON(http.StatusOK, result)

}

func (ind *InDB) CreatePerson(c *gin.Context) {
	var (
		person structs.Person
		result gin.H
	)

	first_name := c.PostForm("first_name")
	last_name := c.PostForm("last_name")

	person.FirstName = first_name
	person.LastName = last_name

	ind.DB.Create(&person)

	result = gin.H{
		"result": person,
	}

	c.JSON(http.StatusOK, result)

}

func (ind *InDB) UpdatePerson(c *gin.Context) {
	id := c.Query("id")
	first_name := c.PostForm("first_name")
	last_name := c.PostForm("last_name")

	var (
		person     structs.Person
		new_person structs.Person
		result     gin.H
	)

	err := ind.DB.First(&person, id).Error

	if err != nil {
		result = gin.H{
			"result": "data not found",
		}
	}

	new_person.FirstName = first_name
	new_person.LastName = last_name

	err = ind.DB.Model(&person).Updates(new_person).Error

	if err != nil {
		result = gin.H{
			"result": "update failed",
		}
	} else {
		result = gin.H{
			"result": "update succed",
		}
	}

	c.JSON(http.StatusOK, result)

}

func (ind *InDB) DeletePerson(c *gin.Context) {
	var (
		person structs.Person
		result gin.H
	)

	id := c.Param("id")
	err := ind.DB.First(&person, id).Error

	if err != nil {
		result = gin.H{
			"result": "data not found",
		}
	}

	err = ind.DB.Delete(&person).Error

	if err != nil {
		result = gin.H{
			"result": "data deletion failed",
		}
	} else {
		result = gin.H{
			"result": "data deletion succed",
		}
	}

	c.JSON(http.StatusOK, result)
}
