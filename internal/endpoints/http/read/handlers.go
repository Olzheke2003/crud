package read

import (
	"crud/internal/services"
	"encoding/json"
	"fmt"
	"net/http"
)

func Handler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			http.Error(w, "Invalid method", http.StatusMethodNotAllowed)
			return
		}
		id := r.URL.Path[len("/read/"):]

		user, err := services.Repository().ReadUser(id)
		if err != nil {
			http.Error(w, fmt.Sprintf("failed to read user %v", err), http.StatusInternalServerError)
			return
		}
		err = json.NewEncoder(w).Encode(user)
		if err != nil {
			http.Error(w, "failed to return user", http.StatusInternalServerError)
			return
		}

	}
}
