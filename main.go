package main

import (
	"embed"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

//go:embed dist/*
var dist embed.FS

var dataPath = "./my.db"

func main() {
	http.Handle("/", http.FileServer(http.FS(dist)))
	http.Handle("/api/db", http.HandlerFunc(HandleDbExec))
	http.Handle("/api/download", http.HandlerFunc(HandleDownloadWrapper()))

	var port string
	if len(os.Args) > 1 {
		port = os.Args[1]
		if len(os.Args) > 2 {
			dataPath = os.Args[2]
		}
	} else {
		port = "86"
	}
	Init(dataPath)
	log.Println("Listening...", port)
	err := http.ListenAndServe(":"+port, nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}

}

type Resp struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

type Param struct {
	Type   string        `json:"type"`
	Sql    string        `json:"sql"`
	Params []interface{} `json:"params"`
}

func HandleDbExec(w http.ResponseWriter, r *http.Request) {
	// 异常处理
	defer exceptHandler(w)
	if r.Method == "OPTIONS" {
		write(w, Resp{})
		return
	}

	var err error
	param := Param{}
	body, _ := io.ReadAll(r.Body)
	json.Unmarshal(body, &param)
	var results []Result
	bytes, _ := json.Marshal(param)
	log.Println(string(bytes))
	switch param.Type {
	case "QUERY":
		results, err = ExecQuery(param)
	case "DELETE":
		ExecDel(param)
	case "UPDATE", "INSERT":
		err = ExecUpdate(param)
	default:
		log.Println("错误")
	}
	resp := Resp{}
	if err != nil {
		resp = Resp{
			Code: -1,
			Msg:  err.Error(),
			Data: results,
		}
	} else {
		resp = Resp{
			Code: 1,
			Msg:  "success",
			Data: results,
		}
	}

	write(w, resp)

}

func HandleDownloadWrapper() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "OPTIONS" {
			write(w, Resp{})
			return
		}
		file, err := os.Open("./my.db")
		if err != nil {
			log.Println("Error opening file:", err)
			return
		}
		defer file.Close()                                         // 确保在函数返回时关闭文件
		w.Header().Set("Content-Type", "application/octet-stream") // 设置响应头为JSON
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		w.Header().Set("Content-Disposition", "attachment;filename=my.db")
		_, err = io.Copy(w, file)
		if err != nil {
			log.Println("Error copying file:", err)
		}
	}
}

func write(w http.ResponseWriter, resp Resp) {
	w.Header().Set("Content-Type", "application/json") // 设置响应头为JSON
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
	json.NewEncoder(w).Encode(resp)
}

func exceptHandler(w http.ResponseWriter) {
	if err := recover(); err != nil {
		e := errors.New("an error occurred")
		log.Println("exceptHandler", err.(error).Error())
		resp := Resp{
			Code: -1,
			Msg:  fmt.Sprintf("exception: %v", e.Error()),
			Data: nil,
		}
		w.Header().Set("Content-Type", "application/json") // 设置响应头为JSON
		json.NewEncoder(w).Encode(resp)
		return
	}
}
