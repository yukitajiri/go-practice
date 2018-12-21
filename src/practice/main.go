package main


import (
	"net"
	"fmt"
)

func main(){
	// サーバー側のソケットを作成する
	// goでソケット通信するにはnetパッケージを使用する(プロトコル名, 受信ホスト名・ポート番号)
	listener, err := net.Listen("tcp", "0.0.0.0:1234")
	if err != nil {
		fmt.Printf("Listen error: %s\n", err)
		return
	}
	// net.listenによってソケット作成、
	defer listener.Close()

	for {
		// コネクションを取り出す(クライアントからの接続を待つ)
		con, err := listener.Accept()
		if err != nil {
			fmt.Printf("Accept error: %s\n", err)
		}
		go func(con net.Conn) {
			_, err = con.Write([]byte("Hello World!\r\n"))
			if err != nil {
				fmt.Printf("Write error: %s\n", err)
			}
			buf := make([]byte, 1024)
			// クライアントからの送信メッセージを受信
			n, err := con.Read(buf)
			if err != nil {
				fmt.Printf("Read error: %s\n", err)
			}
			fmt.Print(string(buf[:n]))
			con.Close()
		}(con)
	}
}
