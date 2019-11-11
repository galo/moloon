package controller

import (
	"github.com/galo/moloon/database"
	"github.com/go-chi/chi"
	"net/http"
)

type ControllerResource struct {
	store database.FunctionStore
}

func NewControllerResource(store database.FunctionStore) *ControllerResource {
	return &ControllerResource{store}
}

func (rs *ControllerResource) router() *chi.Mux {
	//auth, err := jwt.NewTokenAuth()
	//if err != nil {
	//	logging.Logger.Panic(err)
	//}

	r := chi.NewRouter()
	//r.Use(auth.Verifier())
	//r.Use(jwt.Authenticator)

	//r.Use(rs.VersionCtx)
	//r.Use(rs.NamespaceCtx)

	r.Post("/", rs.create)
	r.Get("/", rs.list)

	r.Route("/{functionName}", func(r chi.Router) {
		r.Get("/", rs.get)
		r.Delete("/", rs.delete)
	})

	return r
}

func (rs *ControllerResource) create(w http.ResponseWriter, r *http.Request) {

}

func (rs *ControllerResource) list(w http.ResponseWriter, r *http.Request) {

}

func (rs *ControllerResource) get(w http.ResponseWriter, r *http.Request) {

}

func (rs *ControllerResource) delete(w http.ResponseWriter, r *http.Request) {

}
