package backend

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// Route represents a single outbound route.
type Route struct {
	ID          string `json:"id"`
	Destination string `json:"destination"`
}

// Routes represents a collection of outbound routes.
type Routes struct {
	Routes []Route `json:"routes"`
}

// ManageOutboundRouting handles requests to manage outbound routing.
func ManageOutboundRouting(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		// Return the list of outbound routes
		routes := Routes{
			Routes: []Route{
				{ID: "1", Destination: "http://localhost:8080"},
				{ID: "2", Destination: "http://localhost:8081"},
				{ID: "3", Destination: "http://localhost:8082"},
			},
		}

		jsonBytes, err := json.Marshal(routes)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.Write(jsonBytes)

	case http.MethodPost:
		// Add a new outbound route
		var route Route
		err := json.NewDecoder(r.Body).Decode(&route)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		// TODO: Add the new route to the list of outbound routes

		w.WriteHeader(http.StatusCreated)

	case http.MethodDelete:
		// Delete an outbound route
		id := r.URL.Query().Get("id")

		// TODO: Delete the outbound route with the given ID

		w.WriteHeader(http.StatusNoContent)

	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func main() {
	// Example usage of ManageOutboundRouting
	http.HandleFunc("/outbound-routing", ManageOutboundRouting)

	fmt.Println("Listening on port 8080...")
	http.ListenAndServe(":8080", nil)
}
