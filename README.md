# wakago

go client for Wakatime

## Sample

```go
package main

import (
	"context"
	"encoding/base64"
	"fmt"
	"net/http"

	"yamash723/wakago/wakago"
)

func main() {
	ctx := context.Background()
	wakagoClient := wakago.NewClient(nil)

	encryptedSecretApiKey := base64.StdEncoding.EncodeToString([]byte("xxxxxxxxxxxxxx"))
	header := http.Header{"Authorization": []string{"Basic " + encryptedSecretApiKey}}
	wakagoClient.DefaultHeader = &header

	res, err := wakagoClient.AllTimeSinceTodayService.Get(ctx, "current", nil)
	if err != nil {
		panic(err)
	}

	fmt.Printf("%+v", res)
}
```
