package methods

import (
	"github.com/labstack/echo"
	"gobank/common/types"
	"gobank/echoMiddlewares"
	"gobank/models"
	"net/http"
	"strconv"
)

func RouteGet(g *echo.Group) {
	g.GET("/:userId", Get)
}

// Get ...
func Get(c echo.Context) error {
	userId, err := strconv.ParseUint(c.Param("userId"), 10, 64)

	ctx := c.Request().Context()
	traceId := ctx.Value(echoMiddlewares.ContextTraceId).(string)

	if err != nil {
		respError := types.ApiError{Code: -1, Message: err.Error()}
		resp := types.ApiResponse{Error:respError, TraceId: traceId}
		return c.JSON(http.StatusOK, &resp)
	}

	type Result struct {
		balance models.BalanceEntity
	}

	balance := models.BalanceEntity{}
	if dbRes, err := balance.GetById(ctx, userId); !dbRes || err != nil {
		respError := types.ApiError{Code: -1, Message: err.Error()}
		resp := types.ApiResponse{Error: respError, TraceId: traceId}
		return c.JSON(http.StatusOK, &resp)
	}

	result := Result { balance: balance }
	resp := types.ApiResponse{ Result: result, TraceId: traceId}

	return c.JSON(http.StatusOK, &resp)
}
