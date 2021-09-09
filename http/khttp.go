package ghttp

import (
	"encoding/json"
	"fmt"
	"gnet_test/dollars"
	"gnet_test/model"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
)

func HttpHandler(ss dollars.SqlServer) {
	http.HandleFunc("/list", func(w http.ResponseWriter, r *http.Request) {
		d := dollars.ListServer(ss)
		j, err := json.MarshalIndent(d, "", "  ")
		if err != nil {
			log.Fatalf("Marshal failed,%v\n ", err)
		}
		fmt.Fprintf(w, string(j))
	})

	http.HandleFunc("/add", func(w http.ResponseWriter, r *http.Request) {
		r.ParseForm()
		fmt.Printf(r.Method)
		body, err := ioutil.ReadAll(r.Body)
		var a model.Dollar
		if err = json.Unmarshal(body, &a); err != nil {
			log.Fatalf("Unmarshal err,%v\n", err)
		}
		p, _ := strconv.ParseFloat(a.Price, 32)
		dollars.InsertServer(a.Item, float32(p), ss)
	})

	http.HandleFunc("/search", func(w http.ResponseWriter, r *http.Request) {
		r.ParseForm()
		fmt.Printf(r.Method)
		body, err := ioutil.ReadAll(r.Body)
		var a struct {
			Item string `json:"item"`
		}
		if err = json.Unmarshal(body, &a); err != nil {
			log.Fatalf("Unmarshal err,%v\n", err)
		}
		d := dollars.GetDataOfItemServer(a.Item, ss)
		j, _ := json.MarshalIndent(d, "", "  ")
		fmt.Fprintf(w, string(j))
	})

	http.HandleFunc("/delete", func(w http.ResponseWriter, r *http.Request) {
		r.ParseForm()
		fmt.Printf(r.Method)
		body, err := ioutil.ReadAll(r.Body)
		var a struct {
			Id int `json:"id"`
		}
		if err = json.Unmarshal(body, &a); err != nil {
			log.Fatalf("Unmarshal err,%v\n", err)
		}
		err = dollars.DeleteServer(a.Id, ss)
		if err != nil {
			log.Fatalf("Deleted failed,%v\n", err)
		}
		d := dollars.ListServer(ss)
		j, err := json.MarshalIndent(d, "", "  ")
		if err != nil {
			log.Fatalf("Marshal failed,%v\n ", err)
		}
		fmt.Fprintf(w, string(j))
	})

	http.HandleFunc("/update", func(w http.ResponseWriter, r *http.Request) {
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
		err = dollars.UpdatePrice(a.Id, float32(price), ss)
		if err != nil {
			log.Fatalf("Update failed, %v\n", err)
		}
		d := dollars.GetDataOfIdServer(a.Id, ss)
		j, _ := json.MarshalIndent(d, "", "  ")
		fmt.Fprintf(w, string(j))
	})
}
