package master

import (
	error2 "github.com/galo/moloon/internal/api/error"

	"net/http"

	"github.com/galo/moloon/internal/database"
	"github.com/galo/moloon/internal/disco"
	"github.com/galo/moloon/pkg/models"
	"github.com/go-chi/chi"
	"github.com/go-chi/render"
)

// Resource the controller to manage agents and functions
type Resource struct {
	store            database.FunctionStore
	discoveryService disco.DiscoveryService
}

// NewResource creates the controller to manage agents and functions
func NewResource(store database.FunctionStore, d disco.DiscoveryService) *Resource {

	return &Resource{store, d}
}

type agentResponse struct {
	*models.Agent
}

func (rd *agentResponse) Render(w http.ResponseWriter, r *http.Request) error {
	// Pre-processing before a response is marshalled and sent across the wire
	return nil
}

func newAgentResponse(a *models.Agent) *agentResponse {
	resp := &agentResponse{a}
	return resp
}

func newAgentListResponse(agents []*models.Agent) []render.Renderer {
	list := []render.Renderer{}
	for _, a := range agents {
		list = append(list, newAgentResponse(a))
	}
	return list
}

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

func (rs *Resource) router() *chi.Mux {
	//auth, err := jwt.NewTokenAuth()
	//if err != nil {
	//	logging.Logger.Panic(err)
	//}

	r := chi.NewRouter()
	//r.Use(auth.Verifier())
	//r.Use(jwt.Authenticator)

	//r.Use(rs.VersionCtx)
	//r.Use(rs.NamespaceCtx)

	r.Get("/", rs.listAgents)
	// r.Post("/agents", rs.createAgent)
	//r.Delete("/", rs.deleteAgent)

	return r
}

func (rs *Resource) createAgent(w http.ResponseWriter, r *http.Request) {

}

func (rs *Resource) listAgents(w http.ResponseWriter, r *http.Request) {
	//List all agents
	agents, err := rs.discoveryService.GetAll()
	if err != nil {
		_ = render.Render(w, r, error2.ErrInternalServerError)
		return
	}

	if err := render.RenderList(w, r, newAgentListResponse(agents)); err != nil {
		_ = render.Render(w, r, error2.ErrRender(err))
		return
	}

}

func (rs *Resource) get(w http.ResponseWriter, r *http.Request) {

}

func (rs *Resource) delete(w http.ResponseWriter, r *http.Request) {

}
