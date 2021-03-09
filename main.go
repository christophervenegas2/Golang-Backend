package main

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	// "strconv"
)

// CreditCard ...
type CreditCard struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Money int    `json:"money"`
}

var CreditCards []CreditCard

func getCreditCard(c *gin.Context) {
	c.JSON(200, CreditCards)
	// for _, item := range CreditCards {
	// 	c.JSON(200, gin.H{
	// 		"id":    item.ID,
	// 		"name":  item.Name,
	// 		"money": item.Money,
	// 	})
	// }

	// c.JSON(404, gin.H{
	// 	"message": "No encontrado",
	// })
}

func createCreditCard(c *gin.Context) {
	var reqBody CreditCard
	if err := c.ShouldBindJSON(&reqBody); err != nil {
		c.JSON(422, gin.H{
			"message": "No encontrado",
		})
		return
	}
	reqBody.ID = uuid.New().String()

	CreditCards = append(CreditCards, reqBody)

	c.JSON(200, gin.H{
		"message": "Tarjeta creada con exito",
	})

}

func updateCreditCard(c *gin.Context) {

}
func deleteCreditCard(c *gin.Context) {}

func main() {
	CreditCards = append(CreditCards, CreditCard{
		ID:    "1",
		Name:  "Paty maldonado",
		Money: 20000,
	})
	CreditCards = append(CreditCards, CreditCard{
		ID:    "2",
		Name:  "Alexis Sanchez",
		Money: 406000570,
	})
	CreditCards = append(CreditCards, CreditCard{
		ID:    "3",
		Name:  "Augusto Pinochet",
		Money: 1234567890,
	})

	r := gin.Default()
	userRoutes := r.Group("/api/examen/CreditCard")
	{
		userRoutes.GET("/", getCreditCard)
		userRoutes.POST("/", createCreditCard)
		userRoutes.PATCH("/:id", updateCreditCard)
		userRoutes.DELETE("/:id", deleteCreditCard)
	}

	r.Run()
}
