<!-- Start SDK Example Usage -->


```go
package main

import(
	"context"
	"log"
	golxpets "github.com/speakeasy-sdks/golx_pets"
	"github.com/speakeasy-sdks/golx_pets/pkg/models/shared"
)

func main() {
    s := golxpets.New(
        golxpets.WithSecurity(shared.Security{
            PetstoreAuth: "",
        }),
    )

    ctx := context.Background()
    res, err := s.Pet.AddPetForm(ctx, shared.Pet{
        Category: &shared.Category{
            ID: golxpets.Int64(1),
            Name: golxpets.String("Dogs"),
        },
        ID: golxpets.Int64(10),
        Name: "doggie",
        PhotoUrls: []string{
            "corrupti",
        },
        Status: shared.PetStatusPending.ToPointer(),
        Tags: []shared.Tag{
            shared.Tag{
                ID: golxpets.Int64(715190),
                Name: golxpets.String("Stuart Stiedemann"),
            },
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