package main

import (
    "ose"
    "fmt"
    "net/http"
    "github.com/golang/protobuf/proto"
)

func main() {
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        fmt.Printf("reqeust from:%v\n", r.URL.Path)

        // read body
        r.ParseForm()
        content := r.PostForm["content"]
        fmt.Println("Read content field:", content)
        if content == nil || len(content) < 1 {
            fmt.Println("Invalid content")
            return
        }

        req := ose.Request{}
        err := proto.Unmarshal([]byte(content[0]), &req)
        if err != nil {
            fmt.Printf("unmarshal failed, err:%s", err.Error())
            return
        }
        fmt.Println("Request:", req)
        // write response
        rsp := ose.Response{}
        result := ose.Result{ResourceId: 11007052436330104591, ResourceType: 7, Source: 1}
        rsp.Results = append(rsp.Results, &result)
        rsp.Results = append(rsp.Results, &ose.Result{ResourceId: 2, ResourceType: 1, Source: 1})

        // serialize
        data, err := proto.Marshal(&rsp)
        if err != nil {
            fmt.Printf("Marshal failed, err:%s", err.Error())
            fmt.Fprintln(w, "Marshal Failed")
            return
        }
        _, err = w.Write(data)
        if err != nil {
            fmt.Printf("Failed to write response:%s", err.Error())
            fmt.Fprintln(w, "Write response failed")
            return
        }
        fmt.Println("Succ to write response:%v", rsp)
    })

    http.ListenAndServe(":6732", nil)
}
