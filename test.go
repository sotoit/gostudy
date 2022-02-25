package main

import (
	"fmt"
)

func main() {
	fmt.Println("ユーザー名を入力してください")
	var userName string
	fmt.Scan(&userName)

	fmt.Println("メールアドレスを入力してください")
	var userEmail string
	fmt.Scan(&userEmail)

	fmt.Println("登録が完了しました。")
	fmt.Printf("ユーザー名：%s\n", userName)
	fmt.Printf("メールアドレス:%s\n", userEmail)
}
