package controll

import (
	"fmt"
	"net/http"
	"tender/db"
	"tender/models"

	"github.com/gin-gonic/gin"
)

func CreateTender(c *gin.Context) {
	var tender models.Tender
	if err := c.ShouldBindJSON(&tender); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	query := `
        INSERT INTO tender (name, description, service_type, status, organization_id, creator_username)
        VALUES ($1, $2, $3, $4, $5, $6) RETURNING id
    `

	err := db.DB.QueryRowx(query, tender.Name, tender.Description, tender.ServiceType, tender.Status, tender.OrganizationID, tender.CreatorUsername).
		Scan(&tender.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, tender)
}

func GetTender(c *gin.Context) {
	query := "SELECT * FROM tender"

	var tender []models.Tender

	err := db.DB.Select(&tender, query)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, tender)
}

func EditTend(c *gin.Context) {
	var tender models.Tender
	id := c.Param("id")

	queryS := "select * from tender where id=$1"
	err := db.DB.Get(&tender, queryS, id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Тендер не найден"})
		return
	}
	if err := c.ShouldBindJSON(&tender); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	queryUpd := `
		update tender set name=$1, description=$2, status=$3, organization_id=$4, updated_at=NOW() where id=$5
	`

	_, err = db.DB.Exec(queryUpd, tender.Name, tender.Description, tender.Status, tender.OrganizationID, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Не удалось обновить тендер"})
		return
	}

	c.JSON(http.StatusOK, tender)
}

func GetMyTenders(c *gin.Context) {
	username := c.Query("username")
	var tenders []models.Tender
	query := "select * from tender where creator_username=$1"
	err := db.DB.Select(&tenders, query, username)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, tenders)
}

func StatusTender(c *gin.Context) {
	id := c.Param("id")
	status := c.Query("status")
	var tender models.Tender

	queryS := "select * from tender where id=$1"
	err := db.DB.Get(&tender, queryS, id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Tender не найден"})
		fmt.Println(id)
		return
	}

	queryUpd := `
		update tender set status=$1, updated_at=NOW() where id=$2
	`
	_, err = db.DB.Exec(queryUpd, status, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Не удалось обновить статус tender"})
		return
	}

	err = db.DB.Get(&tender, queryS, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка при получении обновленного тендера"})
		return
	}

	c.JSON(http.StatusOK, tender)
}

func CreateBid(c *gin.Context) {
	var bid models.Bid
	if err := c.ShouldBindJSON(&bid); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	query := `
        insert into bid (name, description, status, tender_id, organization_id, creator_username)
        values ($1, $2, $3, $4, $5, $6) RETURNING id
    `
	err := db.DB.QueryRowx(query, bid.Name, bid.Description, bid.Status, bid.TenderID, bid.OrganizationID, bid.CreatorUsername).Scan(&bid.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка создания bid"})
		return
	}
	c.JSON(http.StatusOK, bid)
}

func GetMyBids(c *gin.Context) {
	username := c.Query("username")
	var bids []models.Bid
	query := "select * from bid where creator_username=$1"
	err := db.DB.Select(&bids, query, username)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, bids)
}

func EditBids(c *gin.Context) {
	var bids models.Bid
	id := c.Param("id")

	queryS := "select * from bid where id=$1"
	err := db.DB.Get(&bids, queryS, id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Bid не найден"})
		return
	}
	if err := c.ShouldBindJSON(&bids); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	queryUpd := `
		update bid set name=$1, description=$2, status=$3, organization_id=$4, updated_at=NOW() where id=$5
	`

	_, err = db.DB.Exec(queryUpd, bids.Name, bids.Description, bids.Status, bids.OrganizationID, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Не удалось обновить bid"})
		return
	}

	c.JSON(http.StatusOK, bids)
}

func BidsFromTender(c *gin.Context) {
	var bids models.Bid
	tender := c.Param("tender_id")

	queryS := "select * from bid where tender_id=$1"
	err := db.DB.Get(&bids, queryS, tender)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Bid не найден"})
		return
	}
	c.JSON(http.StatusOK, bids)
}
