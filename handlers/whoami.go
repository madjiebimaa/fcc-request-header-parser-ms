package handlers

import (
	"fcc-request-header-parser-ms/helpers"
	"fcc-request-header-parser-ms/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func WhoAmIHandler(c *gin.Context) {
	var whoami models.WhoAmI
	ipAddress, err := helpers.GetIPAddress()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "can't get ip address",
		})
		return
	}
	whoami.IPAddress = ipAddress.String()

	lang := c.GetHeader("Accept-Language")
	if lang == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "can't get language from the header",
		})
		return
	}
	whoami.Language = lang

	software := c.GetHeader("User-Agent")
	if software == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "can't get software that used by user from the header",
		})
		return
	}
	whoami.Software = software

	c.JSON(http.StatusOK, whoami)
}
