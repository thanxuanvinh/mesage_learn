package main

import (
	"fmt"

	_ "github.com/thanxuanvinh/mesage/server/hub"
)

func main() {
	fmt.Println("Hàm init() trong server/hub/builder.go sẽ chạy trước để khởi tạo giá trị hub ban đầu")
}
