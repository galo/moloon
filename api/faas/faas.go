package faas

import (
	"github.com/galo/moloon/api"
	"github.com/galo/moloon/database"
	"github.com/galo/moloon/models"
	"github.com/go-chi/chi"
	"github.com/go-chi/render"
	"net/http"
)

// FunctionResource implements account management handler.
type FaaSResource struct {
	Store database.FunctionStore
}

// NewFunctionResource creates and returns an account resource.
func NewFaaSResource(store database.FunctionStore) *FaaSResource {
	return &FaaSResource{
		Store: store,
	}
}

func (rs *FaaSResource) router() *chi.Mux {
	//auth, err := jwt.NewTokenAuth()
	//if err != nil {
	//	logging.Logger.Panic(err)
	//}

	r := chi.NewRouter()
	//r.Use(auth.Verifier())
	//r.Use(jwt.Authenticator)

	//r.Use(rs.VersionCtx)
	//r.Use(rs.NamespaceCtx)

	r.Route("/{functionName}", func(r chi.Router) {
		r.Put("/", rs.instantiate)
	})

	return r
}

// Execute the function
func (rs *FaaSResource) instantiate(w http.ResponseWriter, r *http.Request) {
	functionName := chi.URLParam(r, "functionName")
	if functionName == "" {
		_ = render.Render(w, r, api.ErrNotFound)
		return
	}

	f, err := rs.Store.Get(functionName)
	if err == models.ErrFunctionNotfound {
		_ = render.Render(w, r, api.ErrNotFound)
		return
	}
	if err != nil {
		_ = render.Render(w, r, api.ErrInternalServerError)
		return
	}

	// Execute teh function

}
