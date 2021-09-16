package model

import (
	"time"
)

type Block struct {
	Id           string    `json:"id"`
	Height       int64     `json:"height"`
	Parent       string    `json:"parent"`
	LastKeyUnit  string    `json:"last_key_unit"`
	BlockContent string    `json:"block_content"`
	BlockTime    time.Time `json:"block_time"`
}
