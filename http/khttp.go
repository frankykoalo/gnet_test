package ghttp

import (
	"encoding/hex"
	"encoding/json"
	"fmt"
	"github.com/SHDMT/gravity/platform/consensus/structure"
	"github.com/SHDMT/gravity/platform/gpow/commonstructs"
	"gnet_test/model"
	_ "gnet_test/model"
	"gnet_test/server"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"time"
)

func HttpHandler(ss server.SqlServer) {
	http.HandleFunc("/api/1/list", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("content-type", "text/json")
		d := server.ListServer(ss)
		type out struct {
			Id    int     `json:"id"`
			Item  string  `json:"item"`
			Price float32 `json:"price"`
		}
		var a out
		var con model.Content
		for _, content := range d {
			decodeStr, _ := hex.DecodeString(content.Content)
			json.Unmarshal(decodeStr, &con)
			a.Id = content.Id
			a.Item = con.Item
			a.Price = con.Price
			j, _ := json.MarshalIndent(&a, "", "  ")
			fmt.Fprintf(w, string(j))
			fmt.Fprintf(w, "\n")
		}
	})

	http.HandleFunc("/api/1/add", func(w http.ResponseWriter, r *http.Request) {
		r.ParseForm()
		body, _ := ioutil.ReadAll(r.Body)
		var content model.Content
		json.Unmarshal(body, &content)
		server.InsertServer(content, ss)
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

	http.HandleFunc("/api/1/explorer/updatetps", func(w http.ResponseWriter, r *http.Request) {
		r.ParseForm()
		body, _ := ioutil.ReadAll(r.Body)
		fmt.Println(string(body))
		fmt.Fprintf(w, string(body))
	})

	http.HandleFunc("/api/1/explorer/updateblockinfo", func(w http.ResponseWriter, r *http.Request) {
		b := server.ListBlock(ss)
		var block *commonstructs.Block
		var content []byte
		type blockWithoutContent struct {
			Id          string    `json:"id"`
			Height      int64     `json:"height"`
			Parent      string    `json:"parent"`
			LastKeyUnit string    `json:"last_key_unit"`
			BlockTime   time.Time `json:"block_time"`
		}
		for _, v := range b {
			content, _ = hex.DecodeString(v.Block_Content)
			block = commonstructs.DecodeBlock(content)
			a, _ := json.MarshalIndent(blockWithoutContent{v.Id, v.Height,
				v.Parent, v.Last_Key_Unit, v.Block_Time}, "", "  ")
			c, _ := json.MarshalIndent(block, "", "  ")
			fmt.Fprintf(w, string(a), string(c))
		}
	})

	http.HandleFunc("/api/1/explorer/chainstatus", func(w http.ResponseWriter, r *http.Request) {
		c := server.ListChainStatus(ss)
		list, _ := json.MarshalIndent(c, "", "  ")
		fmt.Fprintf(w, string(list))
	})

	http.HandleFunc("/api/1/explorer/consumeNewBlock", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("consume a new block")
		r.ParseForm()
		body, _ := ioutil.ReadAll(r.Body)
		body, err := json.MarshalIndent(body, "", "  ")
		if err != nil {
			log.Fatalf("can't mashal block body, err:%s\n", err)
		}
		fmt.Fprintf(w, string(body))
	})
	http.HandleFunc("/api/1/explorer/listindexmessage", func(w http.ResponseWriter, r *http.Request) {
		var indexMessage structure.IndexMessage
		type indexMessageWithoutContent struct {
			Id          string `json:"id"`
			UnitId      string `json:"unit_id"`
			DataId      string `json:"data_id"`
			DataVersion string `json:"data_version"`
			BizCode     string `json:"biz_code"`
		}
		for _, im := range server.ListIndexMessage(ss) {
			msgContent, _ := hex.DecodeString(im.MessageContent)
			indexMessage.Deserialize(msgContent)
			json.MarshalIndent(indexMessage, "", "  ")
		}

	})
	fmt.Printf("Https server is listening on port %s\n", "16666")
	http.ListenAndServe(":16666", nil)
}
