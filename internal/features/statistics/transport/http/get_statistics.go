package statistics_transport_http

import (
	"fmt"
	"net/http"
	"time"

	core_logger "github.com/PopovMarko/todo_app/internal/core/logger"
	core_http_response "github.com/PopovMarko/todo_app/internal/core/transport/http/respons"
	core_http_utils "github.com/PopovMarko/todo_app/internal/core/transport/http/utils"
)

// GetStatistics godoc
// @Summary 	Get statistics
// @Description Get statistics parameters from BD with optional filtering
// @Tags 		Statistics
// @Produce		json
// @Param		id query int false "Filter by author ID"
// @Param 		from query string false "Start date in format DD-MM-YYYY"
// @Param 		to query string false "End date in format DD-MM-YYYY"
// @Success		200 {object}  StatisticsDTOResponse "Statistics"
// @Failure		400 {object} core_http_response.ErrorResponse "Bad request"
// @Failure 	500 {object} core_http_response.ErrorResponse "internal server error"
// @Router 		/statistics [get]
func (h *StatisticsHTTPHandler) GetStatistics(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	logger := core_logger.LogFromContext(ctx)
	responseHandler := core_http_response.NewHTTPResponseHandler(logger, w)

	userID, from, to, err := GetUserFromToQueryParam(r)
	if err != nil {
		responseHandler.ErrorResponse("Failed to get query params", err)
		return
	}

	domainStatistics, err := h.statisticsService.GetStatistics(ctx, userID, from, to)
	if err != nil {
		responseHandler.ErrorResponse("Failed to get statistics", err)
		return
	}
	dtoResponse := dtoStatisticsFromDomain(domainStatistics)
	responseHandler.JsonResponse(dtoResponse, http.StatusOK)
}

func GetUserFromToQueryParam(r *http.Request) (*int, *time.Time, *time.Time, error) {
	const (
		userIdQueryParamKey = "id"
		fromQueryParamKey   = "from"
		toQueryParamKey     = "to"
	)

	user, err := core_http_utils.GetIntQueryParams(r, userIdQueryParamKey)
	if err != nil {
		return nil, nil, nil, fmt.Errorf(
			"failed to get user id by key %s: %w",
			userIdQueryParamKey, err,
		)
	}

	from, err := core_http_utils.GetTimeQueryParams(r, fromQueryParamKey)
	if err != nil {
		return nil, nil, nil, fmt.Errorf(
			"failed to get time by key %s: %w",
			fromQueryParamKey, err,
		)
	}

	to, err := core_http_utils.GetTimeQueryParams(r, toQueryParamKey)
	if err != nil {
		return nil, nil, nil, fmt.Errorf(
			"failed to get time by key %s: %w",
			toQueryParamKey, err,
		)
	}
	return user, from, to, nil
}
