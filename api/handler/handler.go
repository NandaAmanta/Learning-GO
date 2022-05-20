package handler

import (
	"api/config"
	"api/model"

	"github.com/gin-gonic/gin"
)

func GetAllUser(c *gin.Context) {

	// Get the Database Connection
	db := config.GetDB()

	// Execute query to get data
	results, err := db.Query("SELECT * FROM user")
	if err != nil {
		// error handler
		panic(err.Error())
	}

	// convert data rows to list Object
	users := []*model.User{}
	for results.Next() {
		user := new(model.User)
		err2 := results.Scan(&user.Id, &user.FirstName, &user.LastName, &user.Email, &user.Password)
		if err2 != nil {
			// error handler
			c.JSON(500, gin.H{
				"ok":         false,
				"statusCode": "INTERNAL_SERVER_ERROR",
				"data":       nil,
			})
			panic(err2.Error())

		} else {
			users = append(users, user)
		}
	}

	// return the result
	c.JSON(200, gin.H{
		"ok":         true,
		"statusCode": "SUCCESS_GET_DATA",
		"data":       users,
	})
}

func Query(c *gin.Context) {
	email := c.Query("email")

	// Get the Database Connection
	db := config.GetDB()

	// Execute query to get data
	results, err := db.Query("SELECT * FROM user WHERE email=?", email)
	if err != nil {
		// error handler
		panic(err.Error())
	}

	// convert data rows to list Object
	users := []*model.User{}
	for results.Next() {
		user := new(model.User)
		err2 := results.Scan(&user.Id, &user.FirstName, &user.LastName, &user.Email, &user.Password)
		if err2 != nil {
			// error handler
			c.JSON(500, gin.H{
				"ok":         false,
				"statusCode": "INTERNAL_SERVER_ERROR",
				"data":       nil,
			})
			panic(err2.Error())

		} else {
			users = append(users, user)
		}
	}
	c.JSON(200, gin.H{
		"data": users,
	})
}

func GetDetailUser(c *gin.Context) {
	id := c.Param("id")

	// Get the Database Connection
	db := config.GetDB()

	// Execute query to get data
	var user model.User
	err := db.QueryRow("SELECT * FROM user WHERE id=?", id).Scan(&user.Id, &user.FirstName, &user.LastName, &user.Email, &user.Password)
	if err != nil {
		// error handler
		panic(err.Error())
	}

	c.JSON(200, gin.H{
		"data": user,
	})
}
