package main

import (
    "fmt"
    "net/http"
    "io"
    "encoding/json"
)

type ApiResponse struct {
    Message string
    Data map[string]string
}

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
        fmt.Fprintf(w, "Message: %s\n", response.Message)
        fmt.Fprintf(w, "Data:\n") 
        for key, val := range response.Data {
            fmt.Fprintf(w, "%s: %s \n", key, val)
        }
    }
    
}

func apiStatusCheck (id string) (ApiResponse, error){
    var respJSON ApiResponse = ApiResponse{}
    resp, err := http.Get(fmt.Sprintf("%sstatusabfrage?orderId=%s",apiEndpoint, id))
    if err != nil {
      return respJSON, err
    }

    body, err := io.ReadAll(resp.Body)
    if err != nil {
        return respJSON, err
    }
    //Umarshalling the response to JSON struct
    err = json.Unmarshal(body, &respJSON) 
    if err != nil {
        return respJSON, err
    }
    return respJSON, nil
}

func main() {

    mux := http.NewServeMux()
	mux.HandleFunc("/order/{id}", orderHandler)

    http.ListenAndServe(":8080", mux)
}