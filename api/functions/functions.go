package functions

import (
	"context"
	"errors"
	"net/http"

	"github.ccom/galo/moloon/models"

	"github.com/dhax/go-base/auth/pwdless"
	"github.com/go-chi/chi"
	"github.com/go-chi/render"
)

// The list of error types returned from account resource.
var (
	ErrFunctionValidation = errors.New("function validation error")
	ErrNotFound = errors.New("function not found error")
	ErrInvalidRequest = errors.New("invalid request")
)

// FunctionStore defines database operations for account.
type FunctionStore interface {
	Get(uid string) (*Function, error)
	Update(Function) error
	Delete(Function) error
}

// FunctionResource implements account management handler.
type FunctionResource struct {
	Store FunctionStore
}

// NewFunctionResource creates and returns an account resource.
func NewFunctionResource(store FunctionStore) *FunctionStore {
	return &NewFunctionResource{
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

	r.Post("/", rs.createFunction)

	r.Route("/{functionName}", func(r chi.Router) {
		r.Use(functionCtx)
		r.Get("/", rs.getFunction)

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
	*Function
}

func (d *functionRequest) Bind(r *http.Request) error {
	// d.ProtectedActive = true
	// d.ProtectedRoles = []string{}
	return nil
}

type functionResponse struct {
	*Function
}

func newFunctionResponse(a *Function) *functionResponse {
	resp := &functionResponse{Function: a}
	return resp
}

func (rs *FunctionResource) get(w http.ResponseWriter, r *http.Request) {
	function := r.Context().Value(functionCtx).(*Function)
	render.Respond(w, r, newFunctionResponse(function))
}

func (rs *FunctionResource) update(w http.ResponseWriter, r *http.Request) {
	acc := r.Context().Value(functionCtx).(*Function)
	data := &functionRequest{Function: acc}
	if err := render.Bind(r, data); err != nil {
		render.Render(w, r, ErrInvalidRequest(err))
		return
	}

	//if err := rs.Store.Update(acc); err != nil {
	//	switch err.(type) {
	//	case validation.Errors:
	//		render.Render(w, r, ErrValidation(ErrAccountValidation, err.(validation.Errors)))
	//		return
	//	}
	//	render.Render(w, r, ErrRender(err))
	//	return
	//}

	render.Respond(w, r, newFunctionResponse(acc))
}

func (rs *FunctionResource) delete(w http.ResponseWriter, r *http.Request) {
	acc := r.Context().Value(functionCtx).(*pwdless.Account)
	if err := rs.Store.Delete(acc); err != nil {
		render.Render(w, r, ErrRender(err))
		return
	}
	render.Respond(w, r, http.NoBody)
}
