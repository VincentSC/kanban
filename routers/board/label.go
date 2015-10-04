package board

import (
	"gitlab.com/kanban/kanban/models"
	"gitlab.com/kanban/kanban/modules/middleware"
	"net/http"
)

// ListLabels gets a list of label on board accessible by the authenticated user.
func ListLabels(ctx *middleware.Context) {
	labels, err := models.ListLabels(ctx.User, ctx.Provider, ctx.Query("board_id"))

	if err != nil {
		ctx.JSON(http.StatusUnauthorized, &models.ResponseError{
			Success: false,
			Message: err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, &models.Response{
		Data: labels,
	})
}
