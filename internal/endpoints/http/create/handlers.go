package create

import (
	"crud/internal/domain/models"
	"crud/internal/services"
	"encoding/json"
	"fmt"
	"net/http"
)

func Handler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, "Invalid method", http.StatusMethodNotAllowed)
			return
		}
		//парсим параметры запроса
		var params Params
		if err := json.NewDecoder(r.Body).Decode(&params); err != nil {
			http.Error(w, "Invalid request params", http.StatusBadRequest)
			return
		}

		//создаем пользователя
		var user models.User
		user.UserProperties = params.UserProperties

		err := services.Repository().CreateUser(user)

		if err != nil {
			http.Error(w, fmt.Sprintf("failed to create user %v", err), http.StatusInternalServerError)
			return

		}

		//пишем в ответ что пользователь создан
		resp := Response{
			Msg: "User created",
		}
		w.WriteHeader(http.StatusCreated)
		_, err = w.Write([]byte(resp.Msg))
		if err != nil {
			http.Error(w, fmt.Sprintf("failed to write response %v", err), http.StatusInternalServerError)
			return
		}
	}
}
