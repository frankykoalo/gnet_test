package model

type IndexMessage struct {
	Id             string `json:"id"`
	UnitId         string `json:"unit_id"`
	DataId         string `json:"data_id"`
	DataVersion    string `json:"data_version"`
	BizCode        string `json:"biz_code"`
	MessageContent string `json:"message_content"`
}
