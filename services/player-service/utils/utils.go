package utils

import (
	"net/http"
	"strconv"

	"github.com/Stardome-Team/Stardome-API/services/player-service/models"
	"github.com/gin-gonic/gin"
)

// ParseQueryToInt ;
func ParseQueryToInt(queries ...string) ([]int, *models.ErrorParsing) {
	var list []int = make([]int, len(queries))

	for i, q := range queries {
		v, err := strconv.Atoi(q)

		if err != nil {
			return []int{}, &models.ErrorParsing{
				Type:       gin.ErrorTypePublic,
				Error:      ErrorInvalidQuery,
				Metadata:   http.StatusText(http.StatusBadRequest),
				StatusCode: http.StatusBadRequest,
			}
		}
		list[i] = v
	}

	return list, nil
}
