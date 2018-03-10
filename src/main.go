package main


import (

"fmt"
"net/http"

)

func main(){
	//　ここでハンドラをとうろくしてエントリーポイントとひもづける
	http.HandleFunc("/", handler)
	//　HTTPサーバの利用
	http.ListenAndServe(":8080", nil)
}

//　レスポンスを書き込む先とクライアントからのリクエストを受け取る
func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "hello, world")
}