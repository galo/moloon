package controller

import (
	error2 "github.com/galo/moloon/api/error"

	"github.com/galo/moloon/database"
	"github.com/galo/moloon/disco"
	"github.com/galo/moloon/models"
	"github.com/go-chi/chi"
	"github.com/go-chi/render"
	"net/http"
)

type ControllerResource struct {
	store            database.FunctionStore
	discoveryService disco.DiscoveryService
}

func NewControllerResource(store database.FunctionStore) *ControllerResource {

	d := disco.NewDiscoveryService()

	return &ControllerResource{store, d}
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

	r.Get("/agents", rs.listAgents)

	//r.Route("/{functionName}", func(r chi.Router) {
	//	r.Get("/", rs.get)
	//	r.Delete("/", rs.delete)
	//})

	return r
}

func (rs *ControllerResource) create(w http.ResponseWriter, r *http.Request) {

}

func (rs *ControllerResource) listAgents(w http.ResponseWriter, r *http.Request) {
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

func (rs *ControllerResource) get(w http.ResponseWriter, r *http.Request) {

}

func (rs *ControllerResource) delete(w http.ResponseWriter, r *http.Request) {

}
