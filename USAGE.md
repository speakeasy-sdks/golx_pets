<!-- Start SDK Example Usage -->
```go
package main

import (
	"context"
	golxpets "github.com/speakeasy-sdks/golx_pets"
	"github.com/speakeasy-sdks/golx_pets/pkg/models/shared"
	"log"
)

func main() {
	s := golxpets.New()

	ctx := context.Background()
	res, err := s.Pet.AddPetForm(ctx, shared.Pet{
		Category: &shared.Category{
			ID:   golxpets.Int64(1),
			Name: golxpets.String("Dogs"),
		},
		ID:   golxpets.Int64(10),
		Name: "doggie",
		PhotoUrls: []string{
			"string",
		},
		Tags: []shared.Tag{
			shared.Tag{},
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
<!-- End SDK Example Usage -->