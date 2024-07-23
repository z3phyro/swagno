package handlers

import (
	. "github.com/go-swagno/swagno"
	"github.com/go-swagno/swagno-fiber/swagger"
	"github.com/go-swagno/swagno/example/fiber/models"
	"github.com/gofiber/fiber/v2"
)

type Handler struct {
}

func NewHandler() Handler {
	return Handler{}
}

func (h Handler) SetRoutes(a *fiber.App) {
	a.Get("/", func(c *fiber.Ctx) error {
		return c.Status(200).SendString("Hello World!")
	}).Name("index")
}

func (h *Handler) SetSwagger(a *fiber.App) {
	endpoints := []Endpoint{
		EndPoint(GET, "/product", "product", Params(),
			nil,
			models.Product{},
			map[string]ErrorSchema{"500": ErrorSchema{Model: models.ErrorResponse{}, Description: "Internal Error"}},
			"Get all products",
			nil),
	}

	sw := CreateNewSwagger("Swagger API", "1.0")
	AddEndpoints(endpoints)

	// set auth
	sw.SetBasicAuth()
	sw.SetApiKeyAuth("api_key", "query")
	sw.SetOAuth2Auth("oauth2_name", "password", "http://localhost:8080/oauth2/token", "http://localhost:8080/oauth2/authorize", Scopes(Scope("read:pets", "read your pets"), Scope("write:pets", "modify pets in your account")))

	// 3 alternative way for describing tags with descriptions
	sw.AddTags(Tag("product", "Product operations"), Tag("merchant", "Merchant operations"))
	sw.AddTags(SwaggerTag{Name: "WithStruct", Description: "WithStruct operations"})
	sw.Tags = append(sw.Tags, SwaggerTag{Name: "headerparams", Description: "headerparams operations"})

	// if you want to export your swagger definition to a file
	sw.ExportSwaggerDocs("swagger/doc.json") // optional

	swagger.SwaggerHandler(a, sw.GenerateDocs(), swagger.Config{Prefix: "/swagger"})
}
