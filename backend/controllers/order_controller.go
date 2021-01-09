package controllers

import (
	"context"
	"fmt"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/sut63/team01/ent"
	"github.com/sut63/team01/ent/medicine"
	"github.com/sut63/team01/ent/company"
	"github.com/sut63/team01/ent/pharmacist"
)

// OrderController defines the struct for the order controller
type OrderController struct {
	client *ent.Client
	router gin.IRouter
}

// Order defines the struct for the order
type Order struct {
	Company    int
	Medicine      int
	Pharmacist int
	amount     int
	price      float64
	total      float64
	Datetime   string
}

// CreateOrder handles POST requests for adding order entities
// @Summary Create order
// @Description Create order
// @ID create-order
// @Accept   json
// @Produce  json
// @Param order body Order true "Order entity"
// @Success 200 {object} ent.Order
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /orders [post]
func (ctl *OrderController) CreateOrder(c *gin.Context) {
	obj := Order{}
	if err := c.ShouldBind(&obj); err != nil {
		c.JSON(400, gin.H{
			"error": "order binding failed",
		})
		return
	}

	m, err := ctl.client.Meedicine.
		Query().
		Where(medicine.IDEQ(int(obj.Medicine))).
		Only(context.Background())

	if err != nil {
		c.JSON(400, gin.H{
			"error": "medicine not found",
		})
		return
	}
	p, err := ctl.client.Company.
		Query().
		Where(company.IDEQ(int(obj.Company))).
		Only(context.Background())

	if err != nil {
		c.JSON(400, gin.H{
			"error": "company not found",
		})
		return
	}
	r, err := ctl.client.Pharmacist.
		Query().
		Where(pharmacist.IDEQ(int(obj.Pharmacist))).
		Only(context.Background())

	if err != nil {
		c.JSON(400, gin.H{
			"error": "pharmacist not found",
		})
		return
	}
	time, err := time.Parse(time.RFC3339, obj.Datetime)


	or, err := ctl.client.Order.
		Create().
		SetMedicine(m).
		SetCompany(p).
		SetPharmacist(r).
		SetAmount(obj.amount).
		SetPrice(obj.price).
		SetTotal(obj.total).
		SetDatetime(time).
		Save(context.Background())

	if err != nil {
		c.JSON(400, gin.H{
			"error": "saving failed",
		})
		return
	}

	c.JSON(200, or)
}

// DeleteOrder handles DELETE requests to delete a order entity
// @Summary Delete a order entity by ID
// @Description get order by ID
// @ID delete-order
// @Produce  json
// @Param id path int true "Order ID"
// @Success 200 {object} gin.H
// @Failure 400 {object} gin.H
// @Failure 404 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /orders/{id} [delete]
func (ctl *OrderController) DeleteOrder(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	err = ctl.client.Order.
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

// ListOrder handles request to get a list of order entities
// @Summary List order entities
// @Description list order entities
// @ID list-order
// @Produce json
// @Param limit  query int false "Limit"
// @Param offset query int false "Offset"
// @Success 200 {array} ent.Order
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /orders [get]
func (ctl *OrderController) ListOrder(c *gin.Context) {
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

	orders, err := ctl.client.Order.
		Query().
		WithPharmacist().
		WithCompany().
		WithMedicine().
		Limit(limit).
		Offset(offset).
		All(context.Background())
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, orders)
}

// NewOrderController creates and registers handles for the order controller
func NewOrderController(router gin.IRouter, client *ent.Client) *OrderController {
	oc := &OrderController{
		client: client,
		router: router,
	}
	oc.register()
	return oc
}

// InitOrderController registers routes to the main engine
func (ctl *OrderController) register() {
	orders := ctl.router.Group("/orders")
	orders.GET("", ctl.ListOrder)

	// CRUD
	orders.POST("", ctl.CreateOrder)
	orders.GET(":id", ctl.DeleteOrder)

}