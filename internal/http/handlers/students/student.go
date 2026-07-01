package students

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"

	"github.com/ankur59/students-api/internal/types"
	response "github.com/ankur59/students-api/internal/utils"
	"github.com/go-playground/validator/v10"
)

func New() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		var student types.Students

		err := json.NewDecoder(r.Body).Decode(&student)

		if err := validator.New().Struct(student); err != nil {
			response.WriteJson(w, http.StatusBadRequest, response.GeneralError(http.StatusBadRequest, "Missing Data", ""))
		}

		if errors.Is(err, io.EOF) {
			response.WriteJson(w, http.StatusBadRequest, response.GeneralError(http.StatusBadRequest, "Data is missing", map[string]string{"status": "error"}))
			return
		}

		response.WriteJson(w, http.StatusCreated, map[string]string{"success": "OK"})

	}
}
