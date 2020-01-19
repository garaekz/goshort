package url

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

// API has the Service functions
type API struct {
	Service Service
}

// ProvideAPI provides Service functionalities
func ProvideAPI(s Service) API {
	return API{Service: s}
}

// FindByCode calls service function to find URL by Code
func (api *API) FindByCode(c *gin.Context) {
	url, status := api.Service.FindByCode(c.Param("code"))
	if status == false {
		c.HTML(404, "index.html", nil)
	} else {
		c.Redirect(http.StatusMovedPermanently, url.OriginalURL)
	}
}

// Create provides Service functionalities
func (api *API) Create(c *gin.Context) {
	var dto DTO
	err := c.BindJSON(&dto)
	if err != nil {
		log.Fatalln(err)
		c.Status(http.StatusBadRequest)
	} else {
		createdURL, status := api.Service.Save(ToURL(dto))
		if status == true {
			c.JSON(http.StatusOK, gin.H{"url": ToDTO(createdURL)})
		} else {
			c.JSON(http.StatusBadRequest, nil)
		}
	}
}
