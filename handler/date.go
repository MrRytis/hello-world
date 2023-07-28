package handler

import (
	"github.com/MrRytis/hello-world/utils"
	"net/http"
	"time"
)

type DateTimeResponse struct {
	DateTime string `json:"datetime"`
}

func DateTime(w http.ResponseWriter, r *http.Request) {
	response := DateTimeResponse{
		DateTime: time.Now().Format(time.RFC3339),
	}

	utils.JSON(w, http.StatusOK, response)
}
