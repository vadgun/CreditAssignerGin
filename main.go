package main

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

// @title Mi API
// @version 1.0
// @description Esta es una API simple en Go con Swagger
// @BasePath /
func main() {
	router := gin.Default()
	// Ruta para servir el asignador de creditos
	router.POST("/credit-assignment/", assignerHandler)
	router.Run(":8080")
}

// @Summary Asignar Créditos segun el monto de inversión
// @Description Asignación de créditos de 300, 500 y 700, segun sea el monto de inversion.
// @Param investment path int32 true "Monto de la inversion"
// @Success 200 {int32} int32 int32 int32 err "0,0,0,nil"
// @Router /credit-assignment [get]
func assignerHandler(c *gin.Context) {
	var investmentAmount Investment
	if err := c.BindJSON(&investmentAmount); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	investment := investmentAmount.Investment
	if investment%100 != 0 || investment == 0 {
		err := errors.New("the investment is not a multiple of 100")
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	} else {
		assigner := NewCreditAssigner()
		credits300, credits500, credits700, err := assigner.Assign(investment)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		} else {
			c.JSON(http.StatusOK, gin.H{
				"credit_type_300": credits300,
				"credit_type_500": credits500,
				"credit_type_700": credits700,
			})
		}
	}
}

type Investment struct {
	Investment int32 `json:"investment"`
}

type CreditAssigner interface {
	Assign(investment int32) (int32, int32, int32, error)
}

type creditAssigner struct{}

func NewCreditAssigner() CreditAssigner {
	return &creditAssigner{}
}

func (ca *creditAssigner) Assign(investment int32) (int32, int32, int32, error) {

	max300 := investment / 300
	max500 := investment / 500
	max700 := investment / 700

	for i := int32(0); i <= max300; i++ {
		for j := int32(0); j <= max500; j++ {
			// This lines are commented because they slow the filter trying to equalize the credits.
			// if j > i {
			// 	break
			// }
			for k := int32(0); k <= max700; k++ {
				// if k > i && k > j {
				// 	break
				// }
				total := i*300 + j*500 + k*700
				if total == investment {
					return i, j, k, nil
				}
			}
		}
	}

	return 0, 0, 0, errors.New("the investment cannot be distributed")
}
