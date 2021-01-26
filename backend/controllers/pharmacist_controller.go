package controllers

import (
	"context"
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/sut63/team01/ent"
	"github.com/sut63/team01/ent/pharmacist"
	"github.com/sut63/team01/ent/positioninpharmacist"
)

// PharmacistController defines the struct for the pharmacist controller
type PharmacistController struct {
	client *ent.Client
	router gin.IRouter
}

//Pharmacist Structure
type Pharmacist struct {
	Email                string
	Password             string
	Name                 string
	PositionInPharmacist int
}

// CreatePharmacist handles POST requests for adding pharmacist entities
// @Summary Create pharmacist
// @Description Create pharmacist
// @ID create-pharmacist
// @Accept   json
// @Produce  json
// @Param pharmacist body ent.Pharmacist true "Pharmacist entity"
// @Success 200 {object} ent.Pharmacist
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /pharmacists [post]
func (ctl *PharmacistController) CreatePharmacist(c *gin.Context) {
	obj := Pharmacist{}
	if err := c.ShouldBind(&obj); err != nil {
		c.JSON(400, gin.H{
			"error": "pharmacist binding failed",
		})
		return
	}

	pip, err := ctl.client.PositionInPharmacist.
		Query().
		Where(positioninpharmacist.IDEQ(int(obj.PositionInPharmacist))).
		Only(context.Background())
	if err != nil {
		c.JSON(400, gin.H{
			"error": "positioninpharmacist not found",
		})
		return
	}

	p, err := ctl.client.Pharmacist.
		Create().
		SetEmail(obj.Email).
		SetPassword(obj.Password).
		SetName(obj.Name).
		SetPositioninpharmacist(pip).
		Save(context.Background())
	if err != nil {
		c.JSON(400, gin.H{
			"error": "saving pharmacist failed",
		})
		return
	}

	c.JSON(200, p)
}

// GetPharmacist handles GET requests to retrieve a pharmacist entity
// @Summary Get a pharmacist entity by ID
// @Description get pharmacist by ID
// @ID get-pharmacist
// @Produce  json
// @Param id path int true "Pharmacist ID"
// @Success 200 {object} ent.Pharmacist
// @Failure 400 {object} gin.H
// @Failure 404 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /pharmacists/{id} [get]
func (ctl *PharmacistController) GetPharmacist(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	p, err := ctl.client.Pharmacist.
		Query().
		WithPositioninpharmacist().
		Where(pharmacist.IDEQ(int(id))).
		Only(context.Background())
	if err != nil {
		c.JSON(404, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, p)
}

// ListPharmacist handles request to get a list of pharmacist entities
// @Summary List pharmacist entities
// @Description list pharmacist entities
// @ID list-pharmacist
// @Produce json
// @Param limit  query int false "Limit"
// @Param offset query int false "Offset"
// @Success 200 {array} ent.Pharmacist
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /pharmacists [get]
func (ctl *PharmacistController) ListPharmacist(c *gin.Context) {
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

	pharmacists, err := ctl.client.Pharmacist.
		Query().
		WithPositioninpharmacist().
		Limit(limit).
		Offset(offset).
		All(context.Background())
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, pharmacists)
}

// DeletePharmacist handles DELETE requests to delete a pharmacist entity
// @Summary Delete a pharmacist entity by ID
// @Description get pharmacist by ID
// @ID delete-pharmacist
// @Produce  json
// @Param id path int true "Pharmacist ID"
// @Success 200 {object} gin.H
// @Failure 400 {object} gin.H
// @Failure 404 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /pharmacists/{id} [delete]
func (ctl *PharmacistController) DeletePharmacist(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	err = ctl.client.Pharmacist.
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

// UpdatePharmacist handles PUT requests to update a pharmacist entity
// @Summary Update a pharmacist entity by ID
// @Description update pharmacist by ID
// @ID update-pharmacist
// @Accept   json
// @Produce  json
// @Param id path int true "Pharmacist ID"
// @Param pharmacist body ent.Pharmacist true "Pharmacist entity"
// @Success 200 {object} ent.Pharmacist
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /pharmacists/{id} [put]
func (ctl *PharmacistController) UpdatePharmacist(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	obj := Pharmacist{}
	if err := c.ShouldBind(&obj); err != nil {
		c.JSON(400, gin.H{
			"error": "pharmacist binding failed",
		})
		return
	}

	pip, err := ctl.client.PositionInPharmacist.
		Query().
		Where(positioninpharmacist.IDEQ(int(obj.PositionInPharmacist))).
		Only(context.Background())
	if err != nil {
		c.JSON(404, gin.H{
			"error": err.Error(),
		})
		return
	}

	p, err := ctl.client.Pharmacist.
		UpdateOneID(int(id)).
		SetEmail(obj.Email).
		SetPassword(obj.Password).
		SetName(obj.Name).
		SetPositioninpharmacist(pip).
		Save(context.Background())
	if err != nil {
		c.JSON(400, gin.H{
			"error": "update pharmacist failed",
		})
		return
	}

	c.JSON(200, p)
}

// NewPharmacistController creates and registers handles for the PharmacistController
func NewPharmacistController(router gin.IRouter, client *ent.Client) *PharmacistController {
	pc := &PharmacistController{
		client: client,
		router: router,
	}
	pc.register()
	return pc
}

// InitPharmacistController registers routes to the main engine
func (ctl *PharmacistController) register() {
	pharmacists := ctl.router.Group("/pharmacists")

	pharmacists.GET("", ctl.ListPharmacist)

	// CRUD
	pharmacists.POST("", ctl.CreatePharmacist)
	pharmacists.GET(":id", ctl.GetPharmacist)
	pharmacists.PUT(":id", ctl.UpdatePharmacist)
	pharmacists.DELETE(":id", ctl.DeletePharmacist)
}
