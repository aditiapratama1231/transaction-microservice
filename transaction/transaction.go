package transaction

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"

	"transaction/model"
)

type Transaction struct {
	DB *gorm.DB
}

func (t Transaction) GetTransactions(c *gin.Context) {
	var transaction []model.Transaction
	db := t.DB

	db.Find(&transaction)

	c.JSON(200, gin.H{
		"data": transaction,
	})
}

func (t Transaction) CreateTransactions(c *gin.Context) {
	var request model.Transaction
	db := t.DB

	if err := c.ShouldBind(&request); err != nil {
		c.JSON(500, gin.H{
			"message": err.Error(),
		})
	}

	db.Create(&request)

	c.JSON(200, gin.H{
		"message": "create transaction success",
		"data":    request,
	})
}
