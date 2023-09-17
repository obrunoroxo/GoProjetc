package main

import (
	"goprojects/TestesGoLang/database"

	// "goprojects/utils"

	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

type personalData struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Nickname string `json:"nickname"`
	Email    string `json:"email"`
	Password string `json:"password"`
	// Password string  `json:"<PASSWORD>"` // TODO: make it a password field in the future
}

type postPersonalData struct {
	Name     string `json:"name"`
	Nickname string `json:"nickname"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func myDatas() []personalData {
	var personal []personalData
	personalDatas := database.GetData()

	for _, values := range personalDatas {
		var data personalData // Create a new datas instance for each value in database
		data.ID = values.ID
		data.Name = values.Name
		data.Nickname = values.Nickname
		data.Email = values.Email
		data.Password = values.Password

		personal = append(personal, data) // Append the current row to the slice
	}

	return personal

}

func getAllDatas(c *gin.Context) {
	personal := myDatas()
	c.IndentedJSON(http.StatusOK, personal)
}

func postNewRegister(c *gin.Context) {
	var personal postPersonalData

	// Call BindJSON to bind the received JSON to postPersonalData.
	if err := c.BindJSON(&personal); err != nil {
		return
	}

	values := database.NewValues{
		Name:     personal.Name,
		Nickname: personal.Nickname,
		Email:    personal.Email,
		Password: personal.Password,
	}

	database.AddNewValues(values)

	myDatas := myDatas()

	c.IndentedJSON(http.StatusCreated, myDatas)

}

func deleteRegister(c *gin.Context) {
	email := c.Query("email")
	password := c.Query("password")

	personal := myDatas()

	for key, value := range personal {
		if ok := value.Email == email && value.Password == password; ok {
			personal = append(personal[:key], personal[key+1:]...)

			delete := database.DeleteValues{
				Email:    email,
				Password: password,
			}

			database.DeleteValue(delete)

			c.JSON(http.StatusOK, gin.H{"delete": "Register deleted successfully!"})

		} else if value.Email != email && value.Password != password {
			c.JSON(http.StatusNotFound, gin.H{"error": "Ops! Register not found..."})
			return
		}
	}

}

func main() {
	router := gin.Default()

	// Configure CORS
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"*"} // Can be specify specific origins if needed
	config.AllowMethods = []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}
	router.Use(cors.New(config))

	router.GET("/api/datas", getAllDatas)

	router.POST("/api/datas", postNewRegister)

	router.DELETE("/api/datas", deleteRegister)

	router.Run("localhost:8080")
}
