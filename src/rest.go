package src

import (
	"github.com/gin-gonic/gin"
	"github.com/shopspring/decimal"
	"net/http"
	"strconv"
)

type BalanceReq struct {
	From   string          `json:"from"`
	To     string          `json:"to"`
	Amount decimal.Decimal `json:"amount"`
}

type RestAPI struct {
	service Service
}

func SetupRouter(service *Service) *gin.Engine {
	api := RestAPI{*service}
	router := gin.Default()
	router.POST("v1/transfer", api.postTransfer)
	router.GET("v1/:username/balance", api.getBalance)
	router.GET("health", api.health)
	return router
}

func (api *RestAPI) postTransfer(c *gin.Context) {
	var req BalanceReq
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Not enough input data"})
		return
	}

	result := api.service.Transaction.Transfer(req.From, req.To, req.Amount)
	c.JSON(http.StatusOK, gin.H{
		"success": strconv.FormatBool(result),
		"from":    req.From,
		"to":      req.To,
		"amount":  req.Amount,
	})
}

func (api *RestAPI) getBalance(c *gin.Context) {
	userId := c.Param("username")
	userBalance := api.service.Balance.Amount(userId)

	c.JSON(http.StatusOK, gin.H{
		"balance":  userBalance,
		"username": userId,
	})
}

func (api *RestAPI) health(c *gin.Context) {
	c.String(http.StatusOK, "UP")
}
