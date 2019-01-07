package handler

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"plugin"

	"github.com/muranoya/mock-server/src/config"
	"github.com/muranoya/mock-server/src/util"

	"github.com/dghubble/trie"
)

type requestContainer struct {
	Path         string
	Plugin       *plugin.Plugin
	AllowMethods []string
}

// HTTPHandler is the custom http handler container
type HTTPHandler struct {
	routing trie.Trier
}

// NewHTTPHandler returnes http handler
func NewHTTPHandler(endpoints []config.EndpointConfig) (*HTTPHandler, error) {
	router := trie.NewRuneTrie()

	for _, ep := range endpoints {
		reqCon := requestContainer{
			Path:         ep.Path,
			AllowMethods: ep.AllowMethod,
		}

		p, err := plugin.Open(ep.Plugin)
		if err != nil {
			return nil, err
		}
		reqCon.Plugin = p

		router.Put(ep.Path, &reqCon)
	}

	return &HTTPHandler{
		routing: router,
	}, nil
}

// ServeHTTP is handling HTTP request
func (h *HTTPHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	con, ok := h.routing.Get(r.RequestURI).(*requestContainer)
	if !ok {
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		fmt.Fprintln(os.Stderr, "trie has invalid container")
		return
	}

	if !util.Containes(con.AllowMethods, r.Method) {
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		fmt.Fprintf(os.Stderr, "method(%v) not allowed on %v, allowed methos are %v\n",
			r.Method, con.Path, con.AllowMethods)
		return
	}

	decodeSym, err := con.Plugin.Lookup("Decode")
	if err != nil {
		http.Error(w, http.StatusText(http.StatusNoContent), http.StatusNoContent)
		fmt.Fprintf(os.Stderr, "%v plugin cannot lookup decode function\n", con.Path)
		return
	}

	decodeFunc, ok := decodeSym.(func(path, body string, header http.Header) ([]byte, error))
	if !ok {
		http.Error(w, http.StatusText(http.StatusNoContent), http.StatusNoContent)
		fmt.Fprintf(os.Stderr, "%v plugin cannot convert to Decode function\n", con.Path)
		return
	}

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusNoContent), http.StatusNoContent)
		fmt.Fprintln(os.Stderr, "body read failed")
		return
	}

	bytes, err := decodeFunc(r.RequestURI, string(body), r.Header)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusNoContent), http.StatusNoContent)
		fmt.Fprintf(os.Stderr, "%v plugin cannot convert to decode function", con.Path)
		return
	}

	w.Write(bytes)
}
