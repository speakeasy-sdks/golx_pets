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
    s := petst.New()
    operationSecurity := operations.AddPetFormSecurity{
            PetstoreAuth: "",
        }

    ctx := context.Background()
    res, err := s.Pet.AddPetForm(ctx, shared.Pet{
        Category: &shared.Category{
            ID: petst.Int64(1),
            Name: petst.String("Dogs"),
        },
        ID: petst.Int64(10),
        Name: "doggie",
        PhotoUrls: []string{
            "corrupti",
        },
        Status: shared.PetStatusPending.ToPointer(),
        Tags: []shared.Tag{
            shared.Tag{
                ID: petst.Int64(715190),
                Name: petst.String("Stuart Stiedemann"),
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