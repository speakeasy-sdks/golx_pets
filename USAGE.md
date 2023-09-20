<!-- Start SDK Example Usage -->


```go
package main

import(
	"context"
	"log"
	"github.com/speakeasy-sdks/golx_pets"
	"github.com/speakeasy-sdks/golx_pets/pkg/models/operations"
	"github.com/speakeasy-sdks/golx_pets/pkg/models/shared"
)

func main() {
    s := golx_pets.New()
    operationSecurity := operations.AddPetFormSecurity{
            PetstoreAuth: "",
        }

    ctx := context.Background()
    res, err := s.Pet.AddPetForm(ctx, shared.Pet{
        Category: &shared.Category{
            ID: golx_pets.Int64(1),
            Name: golx_pets.String("Dogs"),
        },
        ID: golx_pets.Int64(10),
        Name: "doggie",
        PhotoUrls: []string{
            "corrupti",
        },
        Status: shared.PetStatusPending.ToPointer(),
        Tags: []shared.Tag{
            shared.Tag{
                ID: golx_pets.Int64(715190),
                Name: golx_pets.String("Stuart Stiedemann"),
            },
        },
    }, operationSecurity)
    if err != nil {
        log.Fatal(err)
    }

    if res.Pet != nil {
        // handle response
    }
}
```
<!-- End SDK Example Usage -->