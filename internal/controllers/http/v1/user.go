package v1

import (
	"context"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	mongodb "github.com/pavelkorshunov/go-skeleton/pkg/mongo"
	"github.com/spf13/viper"
	"go.mongodb.org/mongo-driver/bson"
)

func UserController(c *gin.Context) {
	// TODO refactor
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongodb.NewClient(ctx, viper.GetString("MONGODB_URI"))
	if err != nil {
		log.Fatal("Failed to connect mongodb")
	}

	collection := client.Database("testing").Collection("users")
	cursor, err := collection.Find(context.TODO(), bson.D{})

	if err != nil {
		log.Fatal("users not found")
	}

	var results []bson.M
	if err = cursor.All(context.TODO(), &results); err != nil {
		log.Fatal("cursor failed")
	}

	c.JSON(http.StatusOK, results)
}
