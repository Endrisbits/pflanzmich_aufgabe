package main

import (
    "fmt"
    "net/http"
    "io"
)

var apiEndpoint = "http://chatapi.freshdesk.intranet.pflanzmich.de:8000/"

func headers(w http.ResponseWriter, req *http.Request) {

    for name, headers := range req.Header {
        for _, h := range headers {
            fmt.Fprintf(w, "%v: %v\n", name, h)
        }
    }
}

func orderHandler (w http.ResponseWriter, req *http.Request) {
    order_id := req.PathValue("id")
    //Ask the API
    response, err := apiStatusCheck(order_id)
    if err != nil {
      fmt.Fprintf(w, "An error occured while processing request: %s", err)
    } else {
        fmt.Fprintf(w, "%s %s", order_id, response["id"], response["body"])
    }
    
}

func apiStatusCheck (id string) (map[string]string, error){
    resp, err := http.Get(fmt.Sprintf("%sstatusabfrage?orderId=%s",apiEndpoint, id))
    if err != nil {
      return nil, err
    }

    body, err := io.ReadAll(resp.Body)
    if err != nil {
        return nil, err
    }
    //Convert the body to type string
    sb := string(body)
    response := make(map[string]string)
    response["id"] = id
    response["body"] = sb
    return response, nil
}

func main() {

    mux := http.NewServeMux()
	mux.HandleFunc("/order/{id}", orderHandler)

    http.ListenAndServe(":8080", mux)
}