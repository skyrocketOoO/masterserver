package rest

import (
	"encoding/json"
	"fmt"
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
	sortStr := c.Query("sort")
	sort := []string{}
	if err := json.Unmarshal([]byte(sortStr), &sort); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid filter parameter"})
		return
	}
	paginationStr := c.Query("range")
	Pagination := []int{}
	if err := json.Unmarshal([]byte(paginationStr), &Pagination); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid filter parameter"})
		return
	}
	filterStr := c.Query("filter")
	filter := make(map[string]interface{})
	if err := json.Unmarshal([]byte(filterStr), &filter); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid filter parameter"})
		return
	}

	users, err := d.usecase.GetUsers(c.Request.Context(), filter,
		domain.Sort{
			Field: sort[0],
			Order: sort[1],
		},
		domain.Pagination{
			Page:    Pagination[0],
			PerPage: Pagination[1],
		})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	type Response struct {
		Data  []postgres.User `json:"data"`
		Total int             `json:"total"`
	}
	c.JSON(http.StatusOK, Response{
		Data:  users,
		Total: len(users),
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
		Data postgres.User `json:"data"`
	}
	c.JSON(http.StatusOK, Response{
		Data: *user,
	})
}

func (d *RestDelivery) GetManyUsers(c *gin.Context) {
	// Parse the IDs from the query string
	idsStr := c.Query("ids")
	fmt.Println(idsStr)
	// var getManyParams GetManyParams
	// if err := json.Unmarshal([]byte(idsStr), &getManyParams); err != nil {
	// 	c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ids parameter"})
	// 	return
	// }

	// // Retrieve the users for the given IDs
	// users, err := d.usecase.GetManyUsers(c.Request.Context(), getManyParams.ids)
	// if err != nil {
	// 	c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	// 	return
	// }

	// // Return the users as a response
	// var responseData []Record // Assuming Record is the struct for your users
	// for _, user := range users {
	// 	// Convert each user to a Record and append to the responseData slice
	// 	record := Record{
	// 		ID:   user.ID,
	// 		Name: user.Name,
	// 		// Add other fields as needed
	// 	}
	// 	responseData = append(responseData, record)
	// }

	// // Construct the response
	// response := GetManyResult{
	// 	Data: responseData,
	// }
	// c.JSON(http.StatusOK, response)
}

func (d *RestDelivery) GetManyReference(c *gin.Context) {
	c.JSON(http.StatusInternalServerError, gin.H{"message": "not implemented"})
}

func (d *RestDelivery) CreateUser(c *gin.Context) {
	type requestBody struct {
		Data postgres.User `json:"data"`
	}
	reqBody := requestBody{}
	if err := c.BindJSON(&reqBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	fmt.Println(reqBody.Data)
	user, err := d.usecase.CreateUser(c.Request.Context(), reqBody.Data)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	type Response struct {
		Data postgres.User `json:"data"`
	}
	c.JSON(http.StatusOK, Response{
		Data: user,
	})
}

func (d *RestDelivery) UpdateUser(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid id parameter"})
	}
	type requestBody struct {
		Data         map[string]interface{} `json:"data"`
		PreviousData postgres.User          `json:"previous_data"`
	}
	reqBody := requestBody{}
	if err := c.BindJSON(&reqBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := d.usecase.UpdateUser(c.Request.Context(), uint(id),
		reqBody.PreviousData, reqBody.Data)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	type Response struct {
		Data postgres.User `json:"data"`
	}
	c.JSON(http.StatusOK, Response{
		Data: user,
	})
}

func (d *RestDelivery) DeleteUser(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid id parameter"})
	}
	type requestBody struct {
		PreviousData postgres.User `json:"previous_data"`
	}
	reqBody := requestBody{}
	if err := c.BindJSON(&reqBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := d.usecase.DeleteUser(c.Request.Context(), uint(id),
		reqBody.PreviousData)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	type Response struct {
		Data postgres.User `json:"data"`
	}
	c.JSON(http.StatusOK, Response{
		Data: user,
	})
}
