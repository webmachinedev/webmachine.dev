package api

import (
	"encoding/json"
	"net/http"

	"github.com/webmachinedev/go-clients/db/functions"
	"github.com/webmachinedev/go-clients/db/types"
)

type Request struct {
	Package string `json:"pkg"`
	Function string `json:"func"`
	Arguments []string `json:"args"`
}

type Response struct {
	Result interface{} `json:"res"`
	Error string `json:"err"`
}

// Handler transforms html form post data into function calls.
func Handler(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Access-Control-Allow-Credentials", "true")
	w.Header().Add("Access-Control-Allow-Origin", r.Header.Get("origin"))
	w.Header().Add("Access-Control-Allow-Methods", "POST")
	w.Header().Add("Access-Control-Allow-Headers", "X-CSRF-Token, X-Requested-With, Accept, Accept-Version, Content-Length, Content-MD5, Content-Type, Date, X-Api-Version")

	w.Header().Set("Content-Type", "application/json")

	var req Request
	err := json.NewDecoder(r.Body).Decode(&req)
    if err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

	var res Response

	switch req.Package {
	case "github.com/webmachinedev/go-clients/db/types":
		switch req.Function {
		case "Index":
			types.Index()
		}
	case "github.com/webmachinedev/go-clients/db/functions":
		switch req.Function {
		case "Index":
			functions.Index()
		}
	}

	json.NewEncoder(w).Encode(res)
}

