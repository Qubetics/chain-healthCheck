package models

type NodeStatus struct {
	Result struct {
		SyncInfo struct {
			CatchingUp        bool   `json:"catching_up"`
			LatestBlockHeight string `json:"latest_block_height"`
		} `json:"sync_info"`
	} `json:"result"`
}
