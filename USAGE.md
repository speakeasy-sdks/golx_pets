<!-- Start SDK Example Usage -->


```go
package main

import(
	"context"
	"log"
	"github.com/speakeasy-sdks/golx_pets"
	"github.com/speakeasy-sdks/golx_pets/pkg/models/shared"
	"github.com/speakeasy-sdks/golx_pets/pkg/models/operations"
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
            "provident",
            "distinctio",
            "quibusdam",
        },
        Status: shared.PetStatusPending.ToPointer(),
        Tags: []shared.Tag{
            shared.Tag{
                ID: petst.Int64(544883),
                Name: petst.String("Ben Mueller"),
            },
            shared.Tag{
                ID: petst.Int64(437587),
                Name: petst.String("Raquel Bednar"),
            },
            shared.Tag{
                ID: petst.Int64(383441),
                Name: petst.String("Alexandra Schulist"),
            },
            shared.Tag{
                ID: petst.Int64(568045),
                Name: petst.String("Mrs. Sophie Smith MD"),
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