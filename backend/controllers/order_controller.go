package controllers
import (
   "context"
   "fmt"
   "strconv"
   "time"
   "github.com/gin-gonic/gin"
   "github.com/sut63/team01/ent"
   "github.com/sut63/team01/ent/order"
   "github.com/sut63/team01/ent/company"
   "github.com/sut63/team01/ent/pharmacist"
   "github.com/sut63/team01/ent/medicine"


) 
// OrderController defines the struct for the order controller
type OrderController struct {
   client *ent.Client
   router gin.IRouter
}

type Order struct{
	Pharmacistid       int
	Companyid          int
	Medicineid         int
	Price			   int
	Hospitalid         string
	Amount			   int
	Addedtime       string
}

// CreateOrder handles POST requests for adding order entities
// @Summary Create order
// @Description Create order
// @ID create-order 
// @Accept   json
// @Produce  json
// @Param order body Order true "Order entity"
// @Success 200 {object} Order 
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
  
	u, err := ctl.client.Pharmacist.
		Query().
		Where(pharmacist.IDEQ(int(obj.Pharmacistid))).
		Only(context.Background())

	if err != nil {
		c.JSON(400, gin.H{
			"error": "pharmacist not found",
		})
		return
	}

	pay, err := ctl.client.Company.
		Query().
		Where(company.IDEQ(int(obj.Companyid))).
		Only(context.Background())

	if err != nil {
		c.JSON(400, gin.H{
			"error": "company not found",
		})
		return
	}

	p, err := ctl.client.Medicine.
		Query().
		Where(medicine.IDEQ(int(obj.Medicineid))).
		Only(context.Background())

	if err != nil {
		c.JSON(400, gin.H{
			"error": "medicine not found",
		})
		return
	}



	times, err := time.Parse(time.RFC3339, obj.Addedtime)

	ol, err := ctl.client.Order.
		Create().
		SetPharmacist(u).
		SetCompany(pay).
		SetMedicine(p).
		SetPrice(obj.Price).
		SetHospitalid(obj.Hospitalid).
		SetAmount(obj.Amount).
        SetAddedtime(times).
		Save(context.Background())
		
		if err != nil {
			fmt.Println(err)
			c.JSON(400, gin.H{
				"status": false,
				"error": err,
			})
			return
		}
		c.JSON(200, gin.H{
			"status": true,
			"data": ol,
		})
	  
	
	 }
// GetOrder handles GET requests to retrieve a order entity
// @Summary Get a order entity by ID
// @Description get order by ID
// @ID get-order
// @Produce  json
// @Param id path int true "Order ID"
// @Success 200 {array} ent.Order
// @Failure 400 {object} gin.H
// @Failure 404 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /orders/{id} [get]
func (ctl *OrderController) GetOrder(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}
	pa, err := ctl.client.Order.
		Query().
		WithMedicine().
		WithCompany().
		WithPharmacist().
		Where(order.HasMedicineWith(medicine.IDEQ(int(id)))).
		All(context.Background())
	if err != nil {
		c.JSON(404, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(200, pa)
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
		if err == nil {limit = int(limit64)}
	}
	offsetQuery := c.Query("offset")
	offset := 0
	if offsetQuery != "" {
		offset64, err := strconv.ParseInt(offsetQuery, 10, 64)
		if err == nil {offset = int(offset64)}
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
		c.JSON(400, gin.H{"error": err.Error(),})
		return
	}
	c.JSON(200, orders)
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
// UpdateOrder handles PUT requests to update a order entity
// @Summary Update a order entity by ID
// @Description update order by ID
// @ID update-order
// @Accept   json
// @Produce  json
// @Param id path int true "Order ID"
// @Param order body ent.Order true "Order entity"
// @Success 200 {object} ent.Order
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /orders/{id} [put]

func (ctl *OrderController) UpdateOrder(c *gin.Context) {

	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
 
	if err != nil {
 
		c.JSON(400, gin.H{
 
			"error": err.Error(),
 
		})
 
		return
 
	}
 
 
	obj := ent.Order{}
 
	if err := c.ShouldBind(&obj); err != nil {
 
		c.JSON(400, gin.H{
 
			"error": "order binding failed",
 
		})
 
		return
 
	}
 
	obj.ID = int(id)
 
	u, err := ctl.client.Order.
 
		UpdateOne(&obj).
 
		Save(context.Background())
 
	if err != nil {
 
		c.JSON(400, gin.H{"error": "update failed",})
 
		return
 
	}
 
 
	c.JSON(200, u)
 
 }
// NewOrderController creates and registers handles for the order controller
func NewOrderController(router gin.IRouter, client *ent.Client) *OrderController {
	uc := &OrderController{
		client: client,
		router: router,
	}
	uc.register()
	return uc
 }
// InitOrderController registers routes to the main engine
 func (ctl *OrderController) register() {
	orders := ctl.router.Group("/orders")
	
	orders.GET("", ctl.ListOrder)
	// CRUD
	orders.POST("", ctl.CreateOrder)
	orders.GET(":id", ctl.GetOrder)
	orders.PUT(":id", ctl.UpdateOrder)
	orders.DELETE(":id", ctl.DeleteOrder)
 }