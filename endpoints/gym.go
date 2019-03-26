package endpoints

import (
	"context"
	"encoding/json"
	"fitup/models"
	"net/http"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
)

// GetGymsHandler endoint
func GetAllGymsHandler(db *mongo.Database, response http.ResponseWriter, request *http.Request) {
	response.Header().Set("content-type", "application/json")
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	result, err := models.Get(ctx, db)
	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(`{ "message": "` + err.Error() + `" }`))
		return
	}
	json.NewEncoder(response).Encode(result)
}
