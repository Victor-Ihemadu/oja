package handlers

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func (h *Handler) HandleGetSellerShopByProfileAndProduct() gin.HandlerFunc {
	return func(c *gin.Context) {
		//get seller path parameter
		sellerID := c.Param("id")

		//find seller with the retrieved ID and return the seller and its product
		Seller, err := h.DB.FindIndividualSellerShop(sellerID)

		if err != nil {
			log.Println("Error finding information in database:", err)
			c.IndentedJSON(http.StatusNotFound, gin.H{
				"Message": "Error Exist ; Seller with this ID not found in product table",
			})
			return
		}
		//5. return a json object of seller profile and product if found
		c.IndentedJSON(http.StatusOK, gin.H{
			"Message":     "Found Seller Shop by Profile and Product",
			"Seller Shop": Seller,
		})

	}
}
