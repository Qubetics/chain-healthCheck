package services

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"
	"time"

	"healthCheck/models"
)

var httpClient = &http.Client{Timeout: 10 * time.Second}

func GetStatus(url string) (height int64, catchingUp bool, err error) {
	resp, err := httpClient.Get(url + "/status")
	if err != nil {
		return
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		err = errors.New("bad node response")
		return
	}
	var status models.NodeStatus
	if err = json.NewDecoder(resp.Body).Decode(&status); err != nil {
		return
	}
	height, err = strconv.ParseInt(status.Result.SyncInfo.LatestBlockHeight, 10, 64)
	if err != nil {
		return
	}
	catchingUp = status.Result.SyncInfo.CatchingUp
	return
}
