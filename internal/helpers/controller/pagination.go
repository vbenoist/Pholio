package controller

import (
	"strconv"

	"github.com/gin-gonic/gin"
	apimodels "github.com/vbenoist/pholio/internal/models/api"
)

func GetPaginationParameters(c *gin.Context) apimodels.PaginationQuery {
	page, err := strconv.ParseInt(c.DefaultQuery("page", "1"), 10, 64)
	if err != nil {
		page = 1
	}
	if page < 1 {
		page = 1
	}

	perPage, err := strconv.ParseInt(c.DefaultQuery("perPage", "10"), 10, 64)
	if err != nil {
		perPage = 10
	}

	return apimodels.PaginationQuery{Page: page, PerPage: perPage}
}
