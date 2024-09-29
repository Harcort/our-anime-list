package utils

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"our-anime-list/backend/constants"
	"our-anime-list/backend/datatransfers"
)

func AuthOnly(c *gin.Context) {
	if !c.GetBool(constants.IsAuthenticatedKey) {
		c.AbortWithStatusJSON(http.StatusUnauthorized, datatransfers.Response{Error: "user not authenticated"})
	}
}
