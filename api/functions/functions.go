package functions

import (
	"context"
	"errors"
	"net/http"

	"github.com/galo/moloon/models"
	"github.com/go-chi/chi"
	"github.com/go-chi/render"
)

// The list of error types returned from account resource.
var (
	ErrFunctionValidation = errors.New("function validation error")
)

// FunctionStore defines database operations for account.
type FunctionStore interface {
	Get(name string) (*models.Function, error)
	Create(models.Function) error
	Delete(models.Function) error
}

// FunctionResource implements account management handler.
type FunctionResource struct {
	Store FunctionStore
}

// NewFunctionResource creates and returns an account resource.
func NewFunctionResource(store FunctionStore) *FunctionResource {
	return &FunctionResource{
		Store: store,
	}
}

func (rs *FunctionResource) router() *chi.Mux {
	//auth, err := jwt.NewTokenAuth()
	//if err != nil {
	//	logging.Logger.Panic(err)
	//}

	r := chi.NewRouter()
	//r.Use(auth.Verifier())
	//r.Use(jwt.Authenticator)
	//r.Use(rs.LinkCtx)

	r.Post("/", rs.create)

	r.Route("/{functionName}", func(r chi.Router) {
		r.Use(rs.functionCtx)
		r.Get("/", rs.get)

	})

	return r
}

func (rs *FunctionResource) functionCtx(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		functionName := chi.URLParam(r, "functionName")
		if functionName == "" {
			_ = render.Render(w, r, ErrNotFound)
			return
		}
		ctx := context.WithValue(r.Context(), "functionName", functionName)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

type functionRequest struct {
	*models.Function
}

func (d *functionRequest) Bind(r *http.Request) error {
	//d.Kind = "function"
	//d.ApiVersion = "v1"
	if d.Metadata.Name == "" {
		return ErrFunctionValidation
	}

	if d.Spec.Image == "" {
		return ErrFunctionValidation
	}
	return nil
}

type functionResponse struct {
	*models.Function
}

func newFunctionResponse(f *models.Function) *functionResponse {
	resp := &functionResponse{Function: f}
	return resp
}

func (rs *FunctionResource) get(w http.ResponseWriter, r *http.Request) {
	function := r.Context().Value(rs.functionCtx).(*models.Function)

	render.Respond(w, r, newFunctionResponse(function))
}

func (rs *FunctionResource) create(w http.ResponseWriter, r *http.Request) {
	data := &functionRequest{}
	if err := render.Bind(r, data); err != nil {
		render.Render(w, r, ErrInvalidRequest(err))
		return
	}

	if err := rs.Store.Create(*data.Function); err != nil {
		render.Render(w, r, ErrInvalidRequest(err))
	}

	render.Respond(w, r, newFunctionResponse(data.Function))
}

func (rs *FunctionResource) delete(w http.ResponseWriter, r *http.Request) {
	f := r.Context().Value(rs.functionCtx).(*models.Function)
	if err := rs.Store.Delete(*f); err != nil {
		render.Render(w, r, ErrRender(err))
		return
	}
	render.Respond(w, r, http.NoBody)
}
