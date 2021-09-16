package ghttp

import (
	"encoding/hex"
	"encoding/json"
	"fmt"
	"github.com/SHDMT/gravity/platform/gpow/commonstructs"
	"gnet_test/model"
	"gnet_test/server"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"time"
)

func HttpHandler(ss server.SqlServer) {
	http.HandleFunc("/api/1/list", func(w http.ResponseWriter, r *http.Request) {
		d := server.ListServer(ss)
		j, err := json.MarshalIndent(d, "", "  ")
		if err != nil {
			log.Fatalf("Marshal failed,%v\n ", err)
		}
		fmt.Fprintf(w, string(j))
	})

	http.HandleFunc("/api/1/add", func(w http.ResponseWriter, r *http.Request) {
		r.ParseForm()
		fmt.Printf(r.Method)
		body, err := ioutil.ReadAll(r.Body)
		var a model.Dollar
		if err = json.Unmarshal(body, &a); err != nil {
			log.Fatalf("Unmarshal err,%v\n", err)
		}
		p, _ := strconv.ParseFloat(a.Price, 32)
		err = server.InsertServer(a.Item, float32(p), ss)
		if err != nil {
			log.Fatalf("Add item failed, %v", err)
		}
	})

	http.HandleFunc("/api/1/search", func(w http.ResponseWriter, r *http.Request) {
		r.ParseForm()
		fmt.Printf(r.Method)
		body, err := ioutil.ReadAll(r.Body)
		var a struct {
			Item string `json:"item"`
		}
		if err = json.Unmarshal(body, &a); err != nil {
			log.Fatalf("Unmarshal err,%v\n", err)
		}
		d := server.GetDataOfItemServer(a.Item, ss)
		j, _ := json.MarshalIndent(d, "", "  ")
		fmt.Fprintf(w, string(j))
	})

	http.HandleFunc("/api/1/delete", func(w http.ResponseWriter, r *http.Request) {
		r.ParseForm()
		fmt.Printf(r.Method)
		body, err := ioutil.ReadAll(r.Body)
		var a struct {
			Id int `json:"id"`
		}
		if err = json.Unmarshal(body, &a); err != nil {
			log.Fatalf("Unmarshal err,%v\n", err)
		}
		err = server.DeleteServer(a.Id, ss)
		if err != nil {
			log.Fatalf("Deleted failed,%v\n", err)
		}
		d := server.ListServer(ss)
		j, err := json.MarshalIndent(d, "", "  ")
		if err != nil {
			log.Fatalf("Marshal failed,%v\n ", err)
		}
		fmt.Fprintf(w, string(j))
	})

	http.HandleFunc("/api/1/update", func(w http.ResponseWriter, r *http.Request) {
		r.ParseForm()
		fmt.Printf(r.Method)
		body, err := ioutil.ReadAll(r.Body)
		var a struct {
			Id    int    `json:"id"`
			Price string `json:"price"`
		}
		if err = json.Unmarshal(body, &a); err != nil {
			log.Fatalf("Unmarshal err,%v\n", err)
		}
		price, _ := strconv.ParseFloat(a.Price, 32)
		err = server.UpdatePrice(a.Id, float32(price), ss)
		if err != nil {
			log.Fatalf("Update failed, %v\n", err)
		}
		d := server.GetDataOfIdServer(a.Id, ss)
		j, _ := json.MarshalIndent(d, "", "  ")
		fmt.Fprintf(w, string(j))
	})

	http.HandleFunc("/api/1/average", func(w http.ResponseWriter, r *http.Request) {
		d := server.ListServer(ss)
		sum := make(map[string]float32, 1)
		sum["average"] = 0
		for _, v := range d {
			price, _ := strconv.ParseFloat(v.Price, 32)
			sum["average"] += float32(price)
		}
		sum["average"] = sum["average"] / float32(len(d))
		result, _ := json.MarshalIndent(sum, "", "  ")
		fmt.Fprintf(w, string(result))
	})

	http.HandleFunc("/api/1/explorer/updateTps", func(w http.ResponseWriter, r *http.Request) {
		r.ParseForm()
		body, _ := ioutil.ReadAll(r.Body)
		fmt.Println(string(body))
		fmt.Fprintf(w, string(body))
	})

	http.HandleFunc("/api/1/explorer/updateblockinfo", func(w http.ResponseWriter, r *http.Request) {
		b := server.ListBlock(ss)
		var block *commonstructs.Block
		var content []byte
		type blockwithoutcontent struct {
			Id          string    `json:"id"`
			Height      int64     `json:"height"`
			Parent      string    `json:"parent"`
			LastKeyUnit string    `json:"last_key_unit"`
			BlockTime   time.Time `json:"block_time"`
		}
		for _, v := range b {
			content, _ = hex.DecodeString(v.BlockContent)
			block = commonstructs.DecodeBlock(content)
			a, _ := json.MarshalIndent(blockwithoutcontent{v.Id, v.Height,
				v.Parent, v.LastKeyUnit, v.BlockTime}, "", "  ")
			c, _ := json.MarshalIndent(block, "", "  ")
			fmt.Fprintf(w, string(a))
			fmt.Fprintf(w, string(c))
		}
	})

	http.HandleFunc("/api/1/explorer/chainstatus", func(w http.ResponseWriter, r *http.Request) {
		c := server.ListChainStatus(ss)
		list, _ := json.MarshalIndent(c, "", "  ")
		fmt.Fprintf(w, string(list))
	})

	http.HandleFunc("/api/1/explorer/consumeNewBlock", func(w http.ResponseWriter, r *http.Request) {
		r.ParseForm()
		body, _ := ioutil.ReadAll(r.Body)
		fmt.Printf("consume a new block , %s\n", body)
		url := fmt.Sprintf("http://%s/api/1/explorer/updateblockinfo", "127.0.0.1:16666")
		req, err := http.NewRequest("POST", url, nil)
		if err != nil {
			log.Fatalf(" can't create http request for data notification, err : %s", err.Error())
			return
		}
		req.Header.Set("Content-Type", "application/json")
		client := &http.Client{}
		res, _ := client.Do(req)
		res.Body.Close()
	})
	ListenAndServe(":16666", nil)

}
