package jwt

import (
	"encoding/json"
	"errors"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
	"sync"

	jose "gopkg.in/square/go-jose.v2"
)

var (
	ErrInvalidContentType = errors.New("should have a JSON content type for JWKS endpoint")
	ErrInvalidAlgorithm   = errors.New("algorithm is invalid")
)

type JWKS struct {
	Keys []jose.JSONWebKey `json:"keys"`
}

type JWKClient struct {
	keyCacher KeyCacher
	mu        sync.Mutex
	jwkUrl    string
}

// fetch does fetch a URL and stores in a tmp file
func fetch(url string) error {
	resp, err := http.Get(url)

	f, err := ioutil.TempFile("/tmp", "jwk.*.json")
	if err != nil {
		return err
	}

	// Create the file
	out, err := os.Create(f.Name())
	if err != nil {
		return err
	}
	defer out.Close()

	// Write the body to file
	_, err = io.Copy(out, resp.Body)
	return err
}

// NewJWKClientWithCache creates a new JWKClient instance
// Passing nil to keyCacher will create a persistent key cacher
func NewJWKClientWithCache(url string, keyCacher KeyCacher) *JWKClient {
	if keyCacher == nil {
		keyCacher = newMemoryPersistentKeyCacher()
	}

	return &JWKClient{
		keyCacher: keyCacher,
		jwkUrl:    url,
	}
}

// GetKey returns the key associated with the provided ID.
func (j *JWKClient) GetKey(ID string) (jose.JSONWebKey, error) {
	j.mu.Lock()
	defer j.mu.Unlock()

	searchedKey, err := j.keyCacher.Get(ID)
	if err != nil {
		keys, err := j.downloadKeys()
		if err != nil {
			return jose.JSONWebKey{}, err
		}
		addedKey, err := j.keyCacher.Add(ID, keys)
		if err != nil {
			return jose.JSONWebKey{}, err
		}
		return *addedKey, nil
	}

	return *searchedKey, nil
}

func (j *JWKClient) downloadKeys() ([]jose.JSONWebKey, error) {
	resp, err := http.Get(j.jwkUrl)
	if err != nil {
		return []jose.JSONWebKey{}, err
	}
	defer resp.Body.Close()

	if contentH := resp.Header.Get("Content-Type"); !strings.HasPrefix(contentH, "application/json") {
		return []jose.JSONWebKey{}, ErrInvalidContentType
	}

	var jwks = JWKS{}
	err = json.NewDecoder(resp.Body).Decode(&jwks)

	if err != nil {
		return []jose.JSONWebKey{}, err
	}

	if len(jwks.Keys) < 1 {
		return []jose.JSONWebKey{}, ErrNoKeyFound
	}

	return jwks.Keys, nil
}
