package handlers

import "github.com/gin-gonic/gin"

func (h Handler) Checkout(context *gin.Context) {
	//TODO: create a struct to handle response from the userservice
	//TODO: Create a function that returns service address and port
	//TODO: Make a request to user-service to fetch the stripe customer id
	// 	and unmarshal that into the struct created in step 1
	//TODO: authorizationHeader := c.Request.Header.Get("Authorization")
	//    req.Header.Set("Authorization", authorizationHeader)
	// Print the customer Id if fetched successfully
}
