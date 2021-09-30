package test

import (
	"encoding/json"
	"fmt"
	"github.com/SHDMT/gravity/platform/consensus/structure"
	"testing"
)

func TestSerialize(t *testing.T) {
	dataId := []byte("21312dasczxvas")
	dataVersion := []byte("1")
	bizCode := []byte("1")

	indexMessage := structure.IndexMessage{
		DataID:      dataId,
		DataVersion: dataVersion,
		BizCode:     bizCode,
	}
	a, _ := json.MarshalIndent(indexMessage, "", "  ")
	fmt.Printf("%s", a)
}
