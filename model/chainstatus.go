package model

import (
	"time"
)

type ChainStatus struct {
	Id                 int       `json:"id"`
	Unit_Mci           int64     `json:"unit_mci"`
	Block_Height       int64     `json:"block_height"`
	Tips_Count         int       `json:"tips_count"`
	Blockchain_Version string    `json:"blockchain_version"`
	Update_Time        time.Time `json:"update_time"`
}
