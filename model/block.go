package model

import (
	"time"
)

type Block struct {
	Id            string    `json:"id"`
	Height        int64     `json:"height"`
	Parent        string    `json:"parent"`
	Last_Key_Unit string    `json:"last_key_unit"`
	Block_Content string    `json:"block_content"`
	Block_Time    time.Time `json:"block_time"`
}
