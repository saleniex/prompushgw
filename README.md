Prometheus push gateway client
===

Helper for sending custom metrics onto prometheus push gateway

### Example

```go
package main

import (
	"github.com/saleniex/prompushgw"
)

keyValues := prompushgw.NewKeyValueList("unique_prefix_")
keyValues.AddStr("key1", "val1")
keyValues.AddInt("key2", 123)

pg := prompushgw.NewPushGateway(os.Getenv("PUSH_GW_URI"), "my_test_job")
pg.Push(keyValues)

```