package controllers

import (
	"ContactListTask/database"
	"ContactListTask/models"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func CreateContact(c *gin.Context) {
	var reqBody models.Contact
	if err := c.ShouldBindJSON(&reqBody); err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"err_message": err.Error(),
		})
		return
	}

	_, err := database.DBClient.Exec("INSERT INTO contact (first_name, last_name, phone, email, position ) VALUES ($1, $2, $3, $4, $5)",
		reqBody.FirstName,
		reqBody.LastName,
		reqBody.Phone,
		reqBody.Email,
		reqBody.Position,
	)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"err_message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "Created!",
	})
}
func GetContacts(c *gin.Context) {
	var contacts []models.Contact

	err := database.DBClient.Select(&contacts, "SELECT *FROM contact;")

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error_message": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, contacts)
}
func GetContact(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	var contact models.Contact
	err := database.DBClient.Get(&contact, "SELECT *FROM contact WHERE id=$1", id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error_message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, contact)
}
func UpdateContact(c *gin.Context) {
	var reqBody models.Contact
	c.ShouldBindJSON(&reqBody)
	id, _ := strconv.Atoi(c.Param("id"))

	_, err := database.DBClient.Exec("UPDATE contact SET first_name=$1, last_name=$2, phone=$3, email=$4, position=$5 WHERE id=$6;",
		reqBody.FirstName,
		reqBody.LastName,
		reqBody.Phone,
		reqBody.Email,
		reqBody.Position,
		id,
	)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error_message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, "update")
}
func DeleteContact(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	_, err := database.DBClient.Exec("DELETE FROM contact WHERE id=$1", id)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, "delete")
}
