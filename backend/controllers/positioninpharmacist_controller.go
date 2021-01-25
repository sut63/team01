package controllers

import (
	"context"
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/sut63/team01/ent"
	"github.com/sut63/team01/ent/positioninpharmacist"
)

// PositionInPharmacistController defines the struct for the positioninpharmacist controller
type PositionInPharmacistController struct {
	client *ent.Client
	router gin.IRouter
}

// CreatePositionInPharmacist handles POST requests for adding positioninpharmacist entities
// @Summary Create positioninpharmacist
// @Description Create positioninpharmacist
// @ID create-positioninpharmacist
// @Accept   json
// @Produce  json
// @Param positioninpharmacist body ent.PositionInPharmacist true "PositionInPharmacist entity"
// @Success 200 {object} ent.PositionInPharmacist
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /positioninpharmacists [post]
func (ctl *PositionInPharmacistController) CreatePositionInPharmacist(c *gin.Context) {
	obj := ent.PositionInPharmacist{}
	if err := c.ShouldBind(&obj); err != nil {
		c.JSON(400, gin.H{
			"error": "positioninpharmacist binding failed",
		})
		return
	}

	pip, err := ctl.client.PositionInPharmacist.
		Create().
		SetPosition(obj.Position).
		Save(context.Background())
	if err != nil {
		c.JSON(400, gin.H{
			"error": "saving positioninpharmacist failed",
		})
		return
	}

	c.JSON(200, pip)
}

// GetPositionInPharmacist handles GET requests to retrieve a positioninpharmacist entity
// @Summary Get a positioninpharmacist entity by ID
// @Description get positioninpharmacist by ID
// @ID get-positioninpharmacist
// @Produce  json
// @Param id path int true "PositionInPharmacist ID"
// @Success 200 {object} ent.PositionInPharmacist
// @Failure 400 {object} gin.H
// @Failure 404 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /positioninpharmacists/{id} [get]
func (ctl *PositionInPharmacistController) GetPositionInPharmacist(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	pip, err := ctl.client.PositionInPharmacist.
		Query().
		Where(positioninpharmacist.IDEQ(int(id))).
		Only(context.Background())
	if err != nil {
		c.JSON(404, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, pip)
}

// ListPositionInPharmacist handles request to get a list of positioninpharmacist entities
// @Summary List positioninpharmacist entities
// @Description list positioninpharmacist entities
// @ID list-positioninpharmacist
// @Produce json
// @Param limit  query int false "Limit"
// @Param offset query int false "Offset"
// @Success 200 {array} ent.PositionInPharmacist
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /positioninpharmacists [get]
func (ctl *PositionInPharmacistController) ListPositionInPharmacist(c *gin.Context) {
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

	positioninpharmacists, err := ctl.client.PositionInPharmacist.
		Query().
		Limit(limit).
		Offset(offset).
		All(context.Background())
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, positioninpharmacists)
}

// DeletePositionInPharmacist handles DELETE requests to delete a positioninpharmacist entity
// @Summary Delete a positioninpharmacist entity by ID
// @Description get positioninpharmacist by ID
// @ID delete-positioninpharmacist
// @Produce  json
// @Param id path int true "PositionInPharmacist ID"
// @Success 200 {object} gin.H
// @Failure 400 {object} gin.H
// @Failure 404 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /positioninpharmacists/{id} [delete]
func (ctl *PositionInPharmacistController) DeletePositionInPharmacist(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	err = ctl.client.PositionInPharmacist.
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

// UpdatePositionInPharmacist handles PUT requests to update a positioninpharmacist entity
// @Summary Update a positioninpharmacist entity by ID
// @Description update positioninpharmacist by ID
// @ID update-positioninpharmacist
// @Accept   json
// @Produce  json
// @Param id path int true "PositionInPharmacist ID"
// @Param positioninpharmacist body ent.PositionInPharmacist true "PositionInPharmacist entity"
// @Success 200 {object} ent.PositionInPharmacist
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /positioninpharmacists/{id} [put]
func (ctl *PositionInPharmacistController) UpdatePositionInPharmacist(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	obj := ent.PositionInPharmacist{}
	if err := c.ShouldBind(&obj); err != nil {
		c.JSON(400, gin.H{
			"error": "positioninpharmacist binding failed",
		})
		return
	}

	pip, err := ctl.client.PositionInPharmacist.
		UpdateOneID(int(id)).
		SetPosition(obj.Position).
		Save(context.Background())
	if err != nil {
		c.JSON(400, gin.H{
			"error": "update positioninpharmacist failed",
		})
		return
	}

	c.JSON(200, pip)
}

// NewPositionInPharmacistController creates and registers handles for the PositionInPharmacistController
func NewPositionInPharmacistController(router gin.IRouter, client *ent.Client) *PositionInPharmacistController {
	pipc := &PositionInPharmacistController{
		client: client,
		router: router,
	}
	pipc.register()
	return pipc
}

// InitPositionInPharmacistController registers routes to the main engine
func (ctl *PositionInPharmacistController) register() {
	positioninpharmacists := ctl.router.Group("/positioninpharmacists")

	positioninpharmacists.GET("", ctl.ListPositionInPharmacist)

	// CRUD
	positioninpharmacists.POST("", ctl.CreatePositionInPharmacist)
	positioninpharmacists.GET(":id", ctl.GetPositionInPharmacist)
	positioninpharmacists.PUT(":id", ctl.UpdatePositionInPharmacist)
	positioninpharmacists.DELETE(":id", ctl.DeletePositionInPharmacist)
}
