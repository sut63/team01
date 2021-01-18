package controllers

import (
	"context"
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/sut63/team01/ent"
	"github.com/sut63/team01/ent/bill"
	"github.com/sut63/team01/ent/pharmacist"
	"github.com/sut63/team01/ent/payment"
	"github.com/sut63/team01/ent/dispensemedicine"
)

// BillController defines the struct for the bill controller
type BillController struct {
	client *ent.Client
	router gin.IRouter
}

//Bill Structure
type Bill struct {
	Annotation   		 string
	Amount				 int
	Payment		  		 int
	DispenseMedicine     int
	Pharmacist	 		 int
}

// CreateBill handles POST requests for adding bill entities
// @Summary Create bill
// @Description Create bill
// @ID create-bill
// @Accept   json
// @Produce  json
// @Param bill body Bill true "Bill entity"
// @Success 200 {object} ent.Bill
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /bills [post]
func (ctl *BillController) CreateBill(c *gin.Context) {
	obj := Bill{}
	if err := c.ShouldBind(&obj); err != nil {
		c.JSON(400, gin.H{
			"error": "bill binding failed",
		})
		return
	}

	py, err := ctl.client.Payment.
		Query().
		Where(payment.IDEQ(int(obj.Payment))).
		Only(context.Background())
		if err != nil {
		c.JSON(400, gin.H{
			"error": "prescription not found",
		})
		return
	}

	dis, err := ctl.client.DispenseMedicine.
		Query().
		Where(dispensemedicine.IDEQ(int(obj.DispenseMedicine))).
		Only(context.Background())
	if err != nil {
		c.JSON(400, gin.H{
			"error": "annotation not found",
		})
		return
	}

	ph, err := ctl.client.Pharmacist.
		Query().
		Where(pharmacist.IDEQ(int(obj.Pharmacist))).
		Only(context.Background())
	if err != nil {
		c.JSON(400, gin.H{
			"error": "pharmacist not found",
		})
		return
	}

	dm, err := ctl.client.Bill.
		Create().
		SetAmount(obj.Amount).
		SetDispenseMedicines(dis).
		SetPayments(py).
		SetAnnotation(obj.Annotation).
		SetPharmacists(ph).
		Save(context.Background())
	if err != nil {
		c.JSON(400, gin.H{
			"error": "saving bill failed",
		})
		return
	}

	c.JSON(200, dm)
}

// GetBill handles GET requests to retrieve a bill entity
// @Summary Get a bill entity by ID
// @Description get bill by ID
// @ID get-bill
// @Produce  json
// @Param id path int true "Bill ID"
// @Success 200 {object} ent.Bill
// @Failure 400 {object} gin.H
// @Failure 404 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /bills/{id} [get]
func (ctl *BillController) GetBill(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	dm, err := ctl.client.Bill.
		Query().
		WithPharmacists().
		WithPayments().
		WithDispenseMedicines().
		Where(bill.IDEQ(int(id))).
		Only(context.Background())
	if err != nil {
		c.JSON(404, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, dm)
}

// ListBill handles request to get a list of bill entities
// @Summary List bill entities
// @Description list bill entities
// @ID list-bill
// @Produce json
// @Param limit  query int false "Limit"
// @Param offset query int false "Offset"
// @Success 200 {array} ent.Bill
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /bills [get]
func (ctl *BillController) ListBill(c *gin.Context) {
	limitQuery := c.Query("limit")
	limit := 10
	if limitQuery != "" {
		limit64, err := strconv.ParseInt(limitQuery, 10, 64)
		if err == nil {
			limit = int(limit64)
		}
	}

	offsetQuery := c.Query("offset")
	offset := 0
	if offsetQuery != "" {
		offset64, err := strconv.ParseInt(offsetQuery, 10, 64)
		if err == nil {
			offset = int(offset64)
		}
	}

	bills, err := ctl.client.Bill.
		Query().
		WithPharmacists().
		WithPayments().
		WithDispenseMedicines().
		Limit(limit).
		Offset(offset).
		All(context.Background())
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, bills)
}

// DeleteBill handles DELETE requests to delete a bill entity
// @Summary Delete a bill entity by ID
// @Description get bill by ID
// @ID delete-bill
// @Produce  json
// @Param id path int true "Bill ID"
// @Success 200 {object} gin.H
// @Failure 400 {object} gin.H
// @Failure 404 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /bills/{id} [delete]
func (ctl *BillController) DeleteBill(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	err = ctl.client.Bill.
		DeleteOneID(int(id)).
		Exec(context.Background())
	if err != nil {
		c.JSON(404, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, gin.H{"result": fmt.Sprintf("ok deleted %v", id)})
}

// UpdateBill handles PUT requests to update a bill entity
// @Summary Update a bill entity by ID
// @Description update bill by ID
// @ID update-bill
// @Accept   json
// @Produce  json
// @Param id path int true "Bill ID"
// @Param bill body ent.Bill true "Bill entity"
// @Success 200 {object} ent.Bill
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /bills/{id} [put]
func (ctl *BillController) UpdateBill(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	obj := Bill{}
	if err := c.ShouldBind(&obj); err != nil {
		c.JSON(400, gin.H{
			"error": "bill binding failed",
		})
		return
	}

	py, err := ctl.client.Payment.
		Query().
		Where(payment.IDEQ(int(obj.Payment))).
		Only(context.Background())
	if err != nil {
		c.JSON(400, gin.H{
			"error": "prescription not found",
		})
		return
	}

	dis, err := ctl.client.DispenseMedicine.
		Query().
		Where(dispensemedicine.IDEQ(int(obj.DispenseMedicine))).
		Only(context.Background())
	if err != nil {
		c.JSON(400, gin.H{
			"error": "annotation not found",
		})
		return
	}

	ph, err := ctl.client.Pharmacist.
		Query().
		Where(pharmacist.IDEQ(int(obj.Pharmacist))).
		Only(context.Background())
	if err != nil {
		c.JSON(400, gin.H{
			"error": "pharmacist not found",
		})
		return
	}
	
	dm, err := ctl.client.Bill.
		UpdateOneID(int(id)).
		SetAmount(obj.Amount).
		SetDispenseMedicines(dis).
		SetPayments(py).
		SetAnnotation(obj.Annotation).
		SetPharmacists(ph).
		Save(context.Background())
	if err != nil {
		c.JSON(400, gin.H{"error": "update bill failed"})
		return
	}

	c.JSON(200, dm)
}

// NewBillController creates and registers handles for the BillController
func NewBillController(router gin.IRouter, client *ent.Client) *BillController {
	dmc := &BillController{
		client: client,
		router: router,
	}
	dmc.register()
	return dmc
}

// InitBillController registers routes to the main engine
func (ctl *BillController) register() {
	bills := ctl.router.Group("/bills")

	bills.GET("", ctl.ListBill)

	// CRUD
	bills.POST("", ctl.CreateBill)
	bills.GET(":id", ctl.GetBill)
	bills.PUT(":id", ctl.UpdateBill)
	bills.DELETE(":id", ctl.DeleteBill)
}
