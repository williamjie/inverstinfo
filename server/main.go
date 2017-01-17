package main

import (
    "net/http"
    "log"
	"encoding/json"
	"io/ioutil"
	//"io"
	"bytes"
)
type MessageInfo struct{
    Phone    string    `json:"phone"`
}

type Result struct{
    Result string  `json:"result"`
    Msg string `json:"msg"`
}


// hello world, the web server
func HelloServer(w http.ResponseWriter, r *http.Request) {
    log.Println("GetMsgSendInfo enter r.Method",r.Method)
	r.ParseForm()
    //if r.Method == "GET" {
        result, _:= ioutil.ReadAll(r.Body)
        r.Body.Close()
        var MsgInfo MessageInfo
        json.Unmarshal([]byte(result), &MsgInfo)

        log.Println("111111:",MsgInfo.Phone)
        pa := &Result{"success","success11111"}
        log.Println("pa",pa)
        js,_ := json.Marshal(pa)
        log.Println("json result:",bytes.NewBuffer(js).String())
        w.Write([]byte(js))
        //io.WriteString(w,bytes.NewBuffer(js).String())
    //}

    //io.WriteString(w, "hello, world!\n")
}

func main() {
	log.Println("main enter")
    http.HandleFunc("/api/GetInfo", HelloServer)
    err := http.ListenAndServe(":12345", nil)
    if err != nil {
        log.Fatal("ListenAndServe: ", err)
    }
}
