package handlers

import (
	"net/http"
	"strconv"

	"healthCheck/config"
	"healthCheck/response"
	"healthCheck/services"

	"github.com/labstack/echo/v4"
)

func HealthCheck(c echo.Context) error {
	env, err := config.GetRequiredEnvVars("VALIDATOR_RPC", "LOCAL_RPC", "HEIGHT_DIFF_LIMIT")
	if err != nil {
		return response.JSON(c, http.StatusInternalServerError, err.Error(), nil)
	}
	validator := env["VALIDATOR_RPC"]
	local := env["LOCAL_RPC"]
	heightDiffStr := env["HEIGHT_DIFF_LIMIT"]
	heightDiff := int64(10) // default value
	if heightDiffStr != "" {
		if v, err := strconv.ParseInt(heightDiffStr, 10, 64); err == nil {
			heightDiff = v
		}
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
	if lCatch && diff > heightDiff {
		return response.JSON(c, http.StatusServiceUnavailable, "catching up height diff too large", data)
	} else if !lCatch || diff <= heightDiff {
		return response.JSON(c, http.StatusOK, "node is healthy", data)
	} else {
		return response.JSON(c, http.StatusServiceUnavailable, "height diff too large", data)
	}
}
