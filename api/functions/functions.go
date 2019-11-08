package functions

import (
	"github.com/galo/moloon/database"
	"net/http"

	"github.com/galo/moloon/models"
	"github.com/go-chi/chi"
	"github.com/go-chi/render"
)

// FunctionResource implements account management handler.
type FunctionResource struct {
	Store database.FunctionStore
}

// NewFunctionResource creates and returns an account resource.
func NewFunctionResource(store database.FunctionStore) *FunctionResource {
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

//func (rs *FunctionResource) functionCtx(next http.Handler) http.Handler {
//	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
//		functionName := chi.URLParam(r, "functionName")
//		if functionName == "" {
//			_ = render.Render(w, r, ErrNotFound)
//			return
//		}
//		ctx := context.WithValue(r.Context(), "functionName", functionName)
//		next.ServeHTTP(w, r.WithContext(ctx))
//	})
//}

// Request and Response Render helpers https://github.com/go-chi/render

// This is the request Render
type newFunctionRequest struct {
	*models.Function
}

func (d *newFunctionRequest) Bind(r *http.Request) error {
	//d.Kind = "function"
	//d.ApiVersion = "v1"
	if d.Metadata.Name == "" {
		return models.ErrFunctionValidation
	}

	if d.Spec.Image == "" {
		return models.ErrFunctionValidation
	}
	return nil
}

// This is the response Render https://github.com/go-chi/render
type functionResponse struct {
	*models.Function
}

func (rd *functionResponse) Render(w http.ResponseWriter, r *http.Request) error {
	// Pre-processing before a response is marshalled and sent across the wire
	return nil
}

func newFunctionResponse(f *models.Function) *functionResponse {
	resp := &functionResponse{Function: f}
	return resp
}

func newFunctionListResponse(fns []*models.Function) []render.Renderer {
	list := []render.Renderer{}
	for _, f := range fns {
		list = append(list, newFunctionResponse(f))
	}
	return list
}

func (rs *FunctionResource) get(w http.ResponseWriter, r *http.Request) {
	functionName := chi.URLParam(r, "functionName")
	if functionName == "" {
		_ = render.Render(w, r, ErrNotFound)
		return
	}

	f, err := rs.Store.Get(functionName)
	if err == models.ErrFunctionNotfound {
		_ = render.Render(w, r, ErrNotFound)
		return
	}
	if err != nil {
		_ = render.Render(w, r, ErrInternalServerError)
		return
	}

	render.Respond(w, r, newFunctionResponse(f))
}

func (rs *FunctionResource) create(w http.ResponseWriter, r *http.Request) {
	data := &newFunctionRequest{}
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
	functionName := chi.URLParam(r, "functionName")
	if functionName == "" {
		_ = render.Render(w, r, ErrNotFound)
		return
	}

	f, err := rs.Store.Get(functionName)
	if err == models.ErrFunctionNotfound {
		_ = render.Render(w, r, ErrNotFound)
		return
	} else if err != nil {
		_ = render.Render(w, r, ErrInternalServerError)
		return
	}

	err = rs.Store.Delete(*f)
	if err != nil {
		_ = render.Render(w, r, ErrInternalServerError)
		return
	}

	render.Respond(w, r, http.NoBody)
}

func (rs *FunctionResource) list(w http.ResponseWriter, r *http.Request) {
	fns, err := rs.Store.GetAll()
	if err == models.ErrFunctionNotfound {
		_ = render.Render(w, r, ErrNotFound)
		return
	} else if err != nil {
		_ = render.Render(w, r, ErrInternalServerError)
		return
	}

	if err := render.RenderList(w, r, newFunctionListResponse(fns)); err != nil {
		_ = render.Render(w, r, ErrRender(err))
		return
	}
}
