# Structur data

## How to install
```bash
github.com/didikprabowo/structur-data
```

## Features
    - Binary Search Tree

### Binary Search Tree

API reference
-   `Insert(v float64)`
-   `Get() *Tree`
-   `BulkInsert([]float64)`
-   `Find(value float64) (node *Node, err error)`
-   `Min() *Node`
-   `Max() *Node`
-   `Delete(value float64)`

Example 
``` go
package main

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/didikprabowo/structur-data/bts"
)
func main() {
    var t bts.Tree

	ExampleData := []int{5, 3, 7, 1, 4, 6, 8, 9}

    // Insert data
    {
        for i := range ExampleData {
            t.Insert(float64(ExampleData[i]))
        }
    }
    // Get Tree
    {
        s, err := json.Marshal(t.Get().Root)
        if err != nil {
            log.Print(err)
        }

        fmt.Println(string(s))
    }

    // Delete node
    {
        t.Delete(7)
    }

    // Get Min value
    {
        t.Min()
    }

     // Get Max value
    {
        t.Max()
    }

    // Find 
    {
        g, err := t.Find(7)
        if err != nil {
            panic(err)
        }
        fmt.Println(g)
    }
}

```
