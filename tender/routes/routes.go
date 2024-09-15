package routes

import (
	"net/http"
	"tender/controll"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()
	router.GET("/api/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "ok")
	})
	router.POST("/api/tenders/new", controll.CreateTender)
	router.GET("/api/tenders", controll.GetTender)
	router.PATCH("/api/tenders/:id/edit", controll.EditTend)
	router.GET("/api/tenders/my", controll.GetMyTenders)
	router.PUT("/api/tenders/:id/status", controll.StatusTender)

	router.POST("/api/bids/new", controll.CreateBid)
	router.GET("/api/bids/my", controll.GetMyBids)
	router.PATCH("/api/bids/:id/edit", controll.EditBids)
	router.GET("/api/bids/:tender_id/list", controll.BidsFromTender)

	return router
}
