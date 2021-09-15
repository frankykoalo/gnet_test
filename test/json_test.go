package test

import (
	"bytes"
	"fmt"
	"log"
	"net/http"
	"testing"
	"time"
)

func TestJson(t *testing.T) {
	heartbeatInterval := time.NewTicker(5 * time.Second)

	defer heartbeatInterval.Stop()
	tps := 2000
	for {
		tps++
		select {
		case <-heartbeatInterval.C:
			url := fmt.Sprintf("http://%s/api/1/explorer/updateTps", "127.0.0.1:10232")
			body := fmt.Sprintf("{\"tps\":\"%d\"}", tps)

			fmt.Printf(" will send data to url : %s , tps : %d", url, tps)
			req, err := http.NewRequest("POST", url, bytes.NewBuffer([]byte(body)))
			if err != nil {
				log.Fatalf(" can't create http request for data notification, err : %s", err.Error())
				continue
			}
			req.Header.Set("Content-Type", "application/json")
			client := &http.Client{}
			res, err := client.Do(req)
			if err != nil {
				log.Fatalf(" can't notification data to DAG-Explorer, err: %s", err.Error())
				continue
			}
			res.Body.Close()
		}
	}
}
