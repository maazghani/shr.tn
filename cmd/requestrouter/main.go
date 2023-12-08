package requestrouter

import (
	"fmt"
	"net/http"
	"strings"
)

type RequestRouter struct {
	Username   string
	Repository string
	Function   string
}

func ParseURL(r *http.Request) RequestRouter {
	path := strings.TrimSpace(r.URL.Path)
	parts := strings.SplitN(path, "/", 3)
	router := RequestRouter{}

	if len(parts) >= 2 {
		router.Username = parts[1]
		if len(parts) == 3 {
			router.Repository = parts[2]
		}
	}

	if router.Repository == "" {
		router.Function = "ExecuteFuzzySearch"
	} else {
		router.Function = "ManageOutboundRouting"
	}

	return router
}

func HandleRequest(w http.ResponseWriter, router RequestRouter) {
	w.Header().Set("X-Function-Invocation", router.Function)

	switch router.Function {
	case "ExecuteFuzzySearch":
		// Invoke OpenFaaS Gateway with appropriate URL and body
		fmt.Println("Invoking ExecuteFuzzySearch...")
	case "ManageOutboundRouting":
		// Redirect user to GitHub repository URL
		fmt.Println("Redirecting user to", "https://github.com/"+router.Username+"/"+router.Repository)
	default:
		fmt.Println("Unsupported function:", router.Function)
	}
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		router := ParseURL(r)
		HandleRequest(w, router)
	})

	fmt.Println("Server listening on port 8080...")
	http.ListenAndServe(":8080", nil)
}
