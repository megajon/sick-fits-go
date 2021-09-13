package graph

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.
import "github.com/megajon/sick-fits-go/postgres"

type Resolver struct {
	UsersRepo    postgres.UsersRepo
	ProductsRepo postgres.ProductsRepo
}
