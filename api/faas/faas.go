package faas

import (
	error2 "github.com/galo/moloon/api/error"
	"github.com/galo/moloon/database"
	"github.com/galo/moloon/models"
	"github.com/galo/moloon/rte"
	"github.com/go-chi/chi"
	"github.com/go-chi/render"
	"net/http"
)

// FunctionResource implements account management handler.
type FaaSResource struct {
	store database.FunctionStore
	rte   rte.Runtime
}

// NewFunctionResource creates and returns an account resource.
func NewFaaSResource(store database.FunctionStore, rte rte.Runtime) *FaaSResource {
	return &FaaSResource{
		store: store,
		rte:   rte,
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
		_ = render.Render(w, r, error2.ErrNotFound)
		return
	}

	f, err := rs.store.Get(functionName)
	if err == models.ErrFunctionNotfound {
		_ = render.Render(w, r, error2.ErrNotFound)
		return
	}
	if err != nil {
		_ = render.Render(w, r, error2.ErrInternalServerError)
		return
	}

	// Execute the function
	rs.rte.Execute(*f)

	render.Respond(w, r, http.NoBody)
}
