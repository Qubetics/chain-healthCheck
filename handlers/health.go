package handlers

import (
	"net/http"
	"os"

	"healthCheck/response"
	"healthCheck/services"

	"github.com/labstack/echo/v4"
)

func HealthCheck(c echo.Context) error {
	validator, local := os.Getenv("VALIDATOR_RPC"), os.Getenv("LOCAL_RPC")
	if validator == "" || local == "" {
		return response.JSON(c, http.StatusInternalServerError, "missing env vars", nil)
	}
	vHeight, vCatch, err := services.GetStatus(validator)
	if err != nil {
		return response.JSON(c, http.StatusServiceUnavailable, "validator unreachable", nil)
	}
	lHeight, lCatch, err := services.GetStatus(local)
	if err != nil {
		return response.JSON(c, http.StatusServiceUnavailable, "local unreachable", nil)
	}
	diff := vHeight - lHeight
	if diff < 0 {
		diff = -diff
	}
	data := map[string]interface{}{
		"validatorHeight": vHeight,
		"localnodeHeight": lHeight,
		"diff":            diff,
		"vcatch":          vCatch,
		"lcatch":          lCatch,
	}
	switch {
	case lCatch:
		return response.JSON(c, http.StatusServiceUnavailable, "catching up", data)
	case diff <= 10:
		return response.JSON(c, http.StatusOK, "ok", data)
	default:
		return response.JSON(c, http.StatusServiceUnavailable, "height diff too large", data)
	}
}
