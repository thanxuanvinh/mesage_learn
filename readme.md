## Phần 1
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

## Phần 2
### B1
- io -> manager.go
    - func NewManager(configReaders <-chan io.Reader) (*Manager, error)
        - đọc dữ liệu từ file server/configs/hubs/...yaml
        - kiểm tra và buid các hubs
- server -> main.go
```go
package main

import (
	"fmt"
	"io"
	"log"
	"os"

	sio "github.com/thanxuanvinh/mesage/io"
	_ "github.com/thanxuanvinh/mesage/server/hub"
)

func main() {
	hubMan := newHubManager()
	fmt.Println(hubMan)
}

func newHubManager() *sio.Manager {
	configReader := make(chan io.Reader, 16)
	for _, s := range []string{"lemon", "mongo"} {
		f, err := os.Open(fmt.Sprintf("./configs/hubs/%sHub.yaml", s))
		if err != nil {
			log.Fatal("error opening config file")
		}
		configReader <- f
	}
	close(configReader)

	man, err := sio.NewManager(configReader)
	if err != nil {
		log.Fatalln("error creating new hub manager: ", err)
	}
	return man
}

```
