package test

import (
	"bytes"
	"fmt"
	"log"
	"net/http"
	"testing"
)

func TestSendBlock(t *testing.T) {
	url := fmt.Sprintf("http://%s/api/1/explorer/consumeNewBlock", "127.0.0.1:16666")
	body := fmt.Sprintf("{\"%s\":\"%s\"}", "blockHash", "fasiofjoi21j3io214jo21imf")
	req, err := http.NewRequest("POST", url, bytes.NewBuffer([]byte(body)))
	if err != nil {
		log.Fatalf("  err: %s", err.Error())
		return
	}
	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		log.Fatalf(" can't notification data to DAG-Explorer, err: %s", err.Error())
		return
	}
	res.Body.Close()
}
