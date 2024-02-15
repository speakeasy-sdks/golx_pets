<!-- Start SDK Example Usage [usage] -->
```go
package main

import (
	"context"
	golxpets "github.com/speakeasy-sdks/golx_pets"
	"github.com/speakeasy-sdks/golx_pets/pkg/models/shared"
	"log"
)

func main() {
	s := golxpets.New(
		golxpets.WithSecurity("Bearer <YOUR_ACCESS_TOKEN_HERE>"),
	)

	ctx := context.Background()
	res, err := s.Pet.AddPetForm(ctx, shared.Pet{
		ID:   golxpets.Int64(10),
		Name: "doggie",
		PhotoUrls: []string{
			"<value>",
		},
	})
	if err != nil {
		log.Fatal(err)
	}

	if res.Pet != nil {
		// handle response
	}
}

```
<!-- End SDK Example Usage [usage] -->