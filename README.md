# "How to Create and Publish Your Own Go Package"
https://sachinborse.hashnode.dev/publish-your-own-go-package

```
go get github.com/devsachinborse/mathops
```
```go
package main

import (
    "fmt"
    "github.com/yourusername/mathops"
)

func main() {
    // Use the Add function from the mathops package
    sum := mathops.Add(10, 5)
    fmt.Println("10 + 5 =", sum)

    // Use the Subtract function from the mathops package
    difference := mathops.Subtract(10, 5)
    fmt.Println("10 - 5 =", difference)
}
```
![SOLID (1)](https://github.com/user-attachments/assets/65344efd-5b70-416d-92ee-0a6da6e35c24)
