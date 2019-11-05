package functions

import (
	"context"
	"github.com/galo/pym/api"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/galo/moloon/auth/jwt"
	"github.com/galo/moloon/logging"
	"github.com/galo/moloon/models"
	"github.com/go-chi/chi"
	"github.com/go-chi/render"
	"github.com/google/uuid"
)

// LinksResource implements links controller handler.
type FunctionsResource struct {
}

// NewLinksResource creates and returns a links resource.
func NewFunctionsResource() *FunctionsResource {
	return &FunctionsResource{}
}

func (rs *FunctionsResource) router() *chi.Mux {
	auth, err := jwt.NewTokenAuth()
	if err != nil {
		logging.Logger.Panic(err)
	}

	r := chi.NewRouter()
	r.Use(auth.Verifier())

	r.Use(jwt.Authenticator)
	//r.Use(rs.LinkCtx)
	r.Put("/", rs.createLink)

	r.Route("/{linkID}", func(r chi.Router) {
		r.Use(LinkIDCtx)
		r.Put("/content", rs.addDoc)

		//r.Get("/content", rs.getDoc)
	})

	return r
}

func (rs *FunctionsResource) addDoc(w http.ResponseWriter, r *http.Request) {
	//Based on https://zupzup.org/go-http-file-upload-download/

	profile := r.PostFormValue("profile")
	file, header, err := r.FormFile("file")
	if err != nil {
		logging.GetLogEntry(r).Error(err)
		_ = render.Render(w, r, ErrBadRequest)
		return
	}
	defer file.Close()

	name := strings.Split(header.Filename, ".")
	logging.GetLogEntry(r).Debug("File name:", name[0])

	// Get file hhandler
	id, err := uuid.NewRandom()
	if err != nil {
		logging.GetLogEntry(r).Error("Error generating UUID", err)
		_ = render.Render(w, r, ErrInternalServerError)
		return
	}

	fileName := id.String()
	newPath := filepath.Join("/tmp", fileName)
	logging.GetLogEntry(r).Debug("Profile: %s, File: %s\n", profile, newPath)

	newFile, err := os.Create(newPath)
	if err != nil {
		logging.GetLogEntry(r).Error("Error creating tmp file", err)
		_ = render.Render(w, r, ErrInternalServerError)
		return
	}
	defer newFile.Close()

	// Read the file and write into disk
	fileBytes, err := io.Copy(newFile, file)
	if err != nil {
		logging.GetLogEntry(r).Error("Error copying file", err)
		_ = render.Render(w, r, ErrInternalServerError)
		return
	}

	// Empty file uploaded
	if fileBytes == 0 {
		logging.GetLogEntry(r).Warn("Empty file uploaded!")
		_ = render.Render(w, r, ErrBadRequest)
		return
	}

	claims := jwt.ClaimsFromCtx(r.Context())

	// Return the document
	d := models.Document{Name: "mydoc",
		DocumentId: "123",
		FolderId:   "123",
		CreatedAt:  api.MakeTimestamp(),
		ModifiedAt: api.MakeTimestamp(),
		Owner:      claims.Sub,
		ModifiedBy: claims.Sub,
		UploadedBy: claims.Sub}

	res := newDocumentResponse(d)

	render.Respond(w, r, res)
}

func (rs *FunctionsResource) createLink(w http.ResponseWriter, r *http.Request) {
	//TODO: Return a link

	l := models.Links{LinkURL: "https://localhost:3000/api/links/aabbccff12bbccc/content"}

	res := newLinksResponse(l)

	render.Respond(w, r, res)
}

// DocumentResponse is the response payload for the Document data model.
//
// In the DocumentResponse object, first a Render() is called on itself,
// then the next field, and so on, all the way down the tree.
// Render is called in top-down order, like a http handler middleware chain.
type DocumentResponse struct {
	models.Document `json:"document,omitempty"`
}

// LinksResponse is the response payload for the Document data model.
//
// In the DocumentResponse object, first a Render() is called on itself,
// then the next field, and so on, all the way down the tree.
// Render is called in top-down order, like a http handler middleware chain.
type LinksResponse struct {
	models.Links `json:"link,omitempty"`
}

func newDocumentResponse(d models.Document) *DocumentResponse {
	resp := &DocumentResponse{Document: d}
	return resp
}

func newLinksResponse(l models.Links) *LinksResponse {
	resp := &LinksResponse{Links: l}
	return resp
}

//func (rs *LinksResource) LinkCtx(next http.Handler) http.Handler {
//	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
//		claims := jwt.ClaimsFromCtx(r.Context())
//		ctx := context.WithValue(r.Context(), ctxProfile, claims)
//		next.ServeHTTP(w, r.WithContext(ctx))
//	})
//}

// LinkCtx middleware is used to load an LinkID  from
// the URL parameters passed through as the request. In case
// the LinkID could not be found, we stop here and return a 404.
func LinkIDCtx(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		linkID := chi.URLParam(r, "linkID")
		if linkID == "" {
			_ = render.Render(w, r, ErrNotFound)
			return
		}
		ctx := context.WithValue(r.Context(), "linkID", linkID)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
