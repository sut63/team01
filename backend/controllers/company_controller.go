package controllers
import (
   "context"
   "fmt"
   "strconv"
   "github.com/sut63/team01/ent"
   "github.com/sut63/team01/ent/company"
   "github.com/gin-gonic/gin"
)
// CompanyController defines the struct for the company controller
type CompanyController struct {
   client *ent.Client
   router gin.IRouter
}
// CreateCompany handles POST requests for adding company entities
// @Summary Create company
// @Description Create company
// @ID create-company  
// @Accept   json
// @Produce  json
// @Param company   body ent.Company true "Company entity"
// @Success 200 {object} ent.Company  
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /companys [post]
func (ctl *CompanyController) CreateCompany(c *gin.Context) {
	obj := ent.Company{}
	if err := c.ShouldBind(&obj); err != nil {
		c.JSON(400, gin.H{
			"error": "company binding failed",
		})
		return
	}
  
	u, err := ctl.client.Company.
		Create().
		SetName(obj.Name).
		Save(context.Background())
	if err != nil {
		c.JSON(400, gin.H{
			"error": "saving failed",
		})
		return
	}
  
	c.JSON(200, u)
 }
// GetCompany handles GET requests to retrieve a company entity
// @Summary Get a company entity by ID
// @Description get company by ID
// @ID get-company
// @Produce  json
// @Param id path int true "Company ID"
// @Success 200 {object} ent.Company
// @Failure 400 {object} gin.H
// @Failure 404 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /companys/{id} [get]
func (ctl *CompanyController) GetCompany(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}
  
	u, err := ctl.client.Company.
		Query().
		Where(company.IDEQ(int(id))).
		Only(context.Background())
	if err != nil {
		c.JSON(404, gin.H{
			"error": err.Error(),
		})
		return
	}
  
	c.JSON(200, u)
 }
 // ListCompany handles request to get a list of company entities
// @Summary List company entities
// @Description list company entities
// @ID list-company
// @Produce json
// @Param limit  query int false "Limit"
// @Param offset query int false "Offset"
// @Success 200 {array} ent.Company
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /companys [get]
func (ctl *CompanyController) ListCompany(c *gin.Context) {
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
  
	companys, err := ctl.client.Company.
		Query().
		Limit(limit).
		Offset(offset).
		All(context.Background())
		if err != nil {
		c.JSON(400, gin.H{"error": err.Error(),})
		return
	}
  
	c.JSON(200, companys)
 }
 // DeleteCompany handles DELETE requests to delete a company entity
// @Summary Delete a company entity by ID
// @Description get company by ID
// @ID delete-company
// @Produce  json
// @Param id path int true "Company ID"
// @Success 200 {object} gin.H
// @Failure 400 {object} gin.H
// @Failure 404 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /companys/{id} [delete]

func (ctl *CompanyController) DeleteCompany(c *gin.Context) {

	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
 
	if err != nil {
 
		c.JSON(400, gin.H{
 
			"error": err.Error(),
 
		})
 
		return
 
	}
 
 
	err = ctl.client.Company.
 
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
// UpdateCompany handles PUT requests to update a company entity
// @Summary Update a company entity by ID
// @Description update company by ID
// @ID update-company
// @Accept   json
// @Produce  json
// @Param id path int true "Company ID"
// @Param company body ent.Company true "Company entity"
// @Success 200 {object} ent.Company
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /companys/{id} [put]

func (ctl *CompanyController) UpdateCompany(c *gin.Context) {

	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
 
	if err != nil {
 
		c.JSON(400, gin.H{
 
			"error": err.Error(),
 
		})
 
		return
 
	}
 
 
	obj := ent.Company{}
 
	if err := c.ShouldBind(&obj); err != nil {
 
		c.JSON(400, gin.H{
 
			"error": "company binding failed",
 
		})
 
		return
 
	}
 
	obj.ID = int(id)
 
	u, err := ctl.client.Company.
 
		UpdateOne(&obj).
 
		Save(context.Background())
 
	if err != nil {
 
		c.JSON(400, gin.H{"error": "update failed",})
 
		return
 
	}
 
 
	c.JSON(200, u)
 
 }
 // NewCompanyController creates and registers handles for the company controller
func NewCompanyController(router gin.IRouter, client *ent.Client) *CompanyController {
	uc := &CompanyController{
		client: client,
		router: router,
	}
	uc.register()
	return uc
 }
  
 // InitCompanyController registers routes to the main engine
 func (ctl *CompanyController) register() {
	companys := ctl.router.Group("/companys")
	companys.GET("", ctl.ListCompany)
	
	// CRUD
	companys.POST("", ctl.CreateCompany)
	companys.GET(":id", ctl.GetCompany)
	companys.PUT(":id", ctl.UpdateCompany)
	companys.DELETE(":id", ctl.DeleteCompany)
 }
