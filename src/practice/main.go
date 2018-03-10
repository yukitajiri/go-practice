package main


import (
"fmt"
"net/http"

)

func main(){

}

//　レスポンスを書き込む先とクライアントからのリクエストを受け取る
func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "hello, world")
}