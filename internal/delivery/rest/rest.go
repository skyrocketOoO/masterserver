package rest

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/skyrocketOoO/masterserver/domain"
	"github.com/skyrocketOoO/masterserver/internal/infra/postgres"
	"github.com/skyrocketOoO/masterserver/internal/usecase"
)

type RestDelivery struct {
	usecase *usecase.Usecase
}

func NewRestDelivery(usecase *usecase.Usecase) *RestDelivery {
	return &RestDelivery{
		usecase: usecase,
	}
}

// @Summary Check the server started
// @Accept json
// @Produce json
// @Success 200 {obj} domain.Response
// @Router /ping [get]
func (d *RestDelivery) Ping(c *gin.Context) {
	c.JSON(http.StatusOK, domain.Response{Message: "pong"})
}

// @Summary Check the server healthy
// @Accept json
// @Produce json
// @Success 200 {obj} domain.Response
// @Failure 503 {obj} domain.Response
// @Router /healthy [get]
func (d *RestDelivery) Healthy(c *gin.Context) {
	// do something check
	if err := d.usecase.Healthy(c.Request.Context()); err != nil {
		c.JSON(http.StatusServiceUnavailable, domain.Response{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, domain.Response{Message: "healthy"})
}

func (d *RestDelivery) GetUsers(c *gin.Context) {
	filterStr := c.Query("filter")
	filter := make(map[string]interface{})
	if err := json.Unmarshal([]byte(filterStr), &filter); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid filter parameter"})
		return
	}

	users, err := d.usecase.GetUsers(c.Request.Context(), filter)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	type Response struct {
		Users []postgres.User `json:"users"`
	}
	c.JSON(http.StatusOK, Response{
		Users: users,
	})
}

func (d *RestDelivery) GetUser(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid id parameter"})
	}

	user, err := d.usecase.GetUser(c.Request.Context(), uint(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	type Response struct {
		User postgres.User `json:"user"`
	}
	c.JSON(http.StatusOK, Response{
		User: *user,
	})
}

func (d *RestDelivery) CreateUser(c *gin.Context) {
	// Handle create record action
	// Parse request body to get the new post data
	var postData map[string]interface{}
	if err := c.BindJSON(&postData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Perform the necessary actions (e.g., save data to database)
	// Return the appropriate response
	c.JSON(http.StatusOK, gin.H{
		"message": "Create a record",
		"data":    postData,
	})
}

func (d *RestDelivery) UpdateUser(c *gin.Context) {
	// Handle update record action
	id := c.Param("id")

	// Parse request body to get the updated post data
	var updateData map[string]interface{}
	if err := c.BindJSON(&updateData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Perform the necessary actions (e.g., update data in database)
	// Return the appropriate response
	c.JSON(http.StatusOK, gin.H{
		"message": "Update a record",
		"id":      id,
		"data":    updateData,
	})
}

func (d *RestDelivery) DeleteUser(c *gin.Context) {
	// Handle delete record action
	id := c.Param("id")

	// Perform the necessary actions (e.g., delete data from database)
	// Return the appropriate response
	c.JSON(http.StatusOK, gin.H{
		"message": "Delete a record",
		"id":      id,
	})
}
