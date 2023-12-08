package backend

import (
	"fmt"
	"net/http"
)

func ManageOutboundRouting(request *http.Request) (interface{}, error) {
	// Extract username and repository from request headers
	username := request.Header.Get("Username")
	repository := request.Header.Get("Repository")

	// Construct the redirect URL
	redirectURL := fmt.Sprintf("https://github.com/%s/%s", username, repository)

	// Send a redirect response
	return http.Redirect(request.Context().ResponseWriter, request.Context().Request, redirectURL, http.StatusFound)
}

func main() {
	// Register the function with OpenFaaS
	http.HandleFunc("/", ManageOutboundRouting)

	// Start the server
	http.ListenAndServe(":8082", nil)
}
