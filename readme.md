# Phần 1

- Khởi tạo hub ban đầu khi chạy chương trình

## B1

- tạo commons -> types.go

## B2

- io -> hub -> hub.go (tạo hub IOs)
- io -> builders.go

## B3

- server -> hub -> hub.go
- server -> hub -> builder.go

## B4

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

- Check file yaml config của hub, và lưu config vào hub đã khởi tạo ở bước 1

### 2_B1

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

## Phần 3

- Khởi tạo Service ban đầu khi chạy chương trình

### 3_B1

- services/service/service.go
- services/builders.go

```go
var builderRegistry map[string]service.Builder = map[string]service.Builder{}

// ClearBuilders clears registered builders
func ClearBuilders() {
 builderRegistry = map[string]service.Builder{}
}

// RegisterBuilderOrDie registers a service builder or dies.
// This is expected to be called by init functions of builder implementations.
func RegisterBuilderOrDie(name string, builder service.Builder) {
 if _, ok := builderRegistry[name]; ok {
  log.Panicf("service builder name collision [%s]", name)
 }
if builder == nil {
  log.Panicf("nil service builder [%s]", name)
 }
 builderRegistry[name] = builder
}
```

### 3_B2

- server/service/service.go
- server/service/builder.go

```go
const builderName = "mock"

type builder int

// NewConfig implementation
func (b *builder) NewConfig() interface{} {
  return new(Config)
}

// Build implementation
func (b *builder) Build(name string, configInterface interface{}) (service.Service, error) {
  config, ok := configInterface.(*Config)
    if !ok {
      return nil, fmt.Errorf("incompatible config")
    }

  return &Service{
    name:   name,
    config: config,
  }, nil
}

func init() {
  services.RegisterBuilderOrDie(builderName, new(builder))
}
```

### 3_B3

- server/main.go
  - import `_ "github.com/thanxuanvinh/mesage/server/service"`

Để chạy hàm init() trong server/service/builder.go

## Phần 4

- check file yaml config của service, và lưu config service đã khởi tạo ở bước 3

### 4_B1

- services/manager.go

- services/builders.go
  - Thêm hàm `Build`
  