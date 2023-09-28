### B1
-  tạo commons -> types.go

### B2
- io -> hub -> hub.go (tạo hub IOs)
- io -> builders.go

### B3
- server -> hub -> hub.go
- server -> hub -> builder.go

### B4
- server -> main.go
    ```go
    package main

    import (
        "fmt"

        _ "github.com/thanxuanvinh/mesage/server/hub"
    )

    func main() {
        fmt.Println("aaa")
    }
    ```

- khi import `_ "github.com/thanxuanvinh/mesage/server/hub"`
- sẽ chạy hàm init() trong server -> hub -> builder.go
- có tác dụng khởi tạo hub ban đầu với config là rỗng
    ```go
    func init() {
	io.RegisterHubBuilderOrDie(builderName, new(builder))
    }
    ```
