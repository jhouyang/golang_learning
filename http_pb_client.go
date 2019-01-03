package main

import (
    "ose"
    "fmt"
    "net/http"
    "io/ioutil"
    "github.com/golang/protobuf/proto"
)

func main() {
    rsp, err := http.Get("http://127.0.0.1:6732/test/")
    if err != nil {
        fmt.Println(err)
        return
    }
    defer rsp.Body.Close()
    body, err := ioutil.ReadAll(rsp.Body)
    if err != nil {
        fmt.Printf("ioutil.ReadAll err:%s", err.Error())
        return
    }

    rsep := ose.Response{}
    err = proto.Unmarshal(body, &rsep)
    if err != nil {
        fmt.Printf("unmarshal failed, err:%s", err.Error())
        return
    }
    fmt.Printf("response from server:%v", rsep)
}
