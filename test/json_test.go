package test

import (
	"bytes"
	"crypto/tls"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"testing"
	"time"
)

func TestJson(t *testing.T) {
	fmt.Println(time.Now())
	heartbeatInterval := time.NewTicker(3 * time.Second)

	defer heartbeatInterval.Stop()
	for {
		rand.Seed(time.Now().Unix())
		tps := rand.Intn(500)
		select {
		case <-heartbeatInterval.C:
			url := fmt.Sprintf("https://%s/api/1/explorer/updateTps", "127.0.0.1:16666")
			body := fmt.Sprintf("{\"tps\":\"%d\"}", tps)

			fmt.Printf(" will send data to url : %s , tps : %d", url, tps)
			req, err := http.NewRequest("POST", url, bytes.NewBuffer([]byte(body)))
			if err != nil {
				log.Fatalf(" can't create http request for data notification, err : %s", err.Error())
				continue
			}
			req.Header.Set("Content-Type", "application/json")
			tr := &http.Transport{
				TLSClientConfig: &tls.Config{
					InsecureSkipVerify: true,
				},
			}
			client := &http.Client{Transport: tr}
			//client := &http.Client{}
			res, err := client.Do(req)
			if err != nil {
				log.Fatalf(" can't notification data to DAG-Explorer, err: %s", err.Error())
				continue
			}
			res.Body.Close()
		}
	}
}
