package main

import (
	"context"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"net/http"
)

//Маршрут 2
//Обновляет пару токенов
func (h *handler) refreshTokens(c echo.Context) error {

	guidU := c.FormValue("GUID")
	refreshTok := c.FormValue("RefreshToken")

	token, err := jwt.Parse(refreshTok, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return []byte("secret"), nil
	})

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		if int(claims["sub"].(float64)) == 1 {

			if mongoFind(refreshTok) != 0 {

				// Create client
				client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://127.0.0.1:27017"))
				if err != nil {
					log.Fatal(err)
				}

				// Create connect
				err = client.Connect(context.TODO())
				if err != nil {
					log.Fatal(err)
				}

				// Check the connection
				err = client.Ping(context.TODO(), nil)
				if err != nil {
					log.Fatal(err)
				}

				collection := client.Database("Users").Collection("All")

				filter := bson.D{{"RefreshToken", refreshTok }}

				//Удаляем токен
				deleteResult, err := collection.DeleteOne(context.TODO(), filter)
				if err != nil {
					log.Fatal(err)
				}
				fmt.Printf("Deleted %v documents in the trainers collection\n", deleteResult.DeletedCount)

				newTokenPair, err := generateTokenPair(guidU)
				if err != nil {
					return err
				}

				return c.JSON(http.StatusOK, newTokenPair)
			}
		}
		return echo.ErrUnauthorized
	}
	return err
}

func mongoFind(rt string) int {

	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://127.0.0.1:27017"))
	if err != nil {
		log.Fatal(err)
	}

	// Create connect
	err = client.Connect(context.TODO())
	if err != nil {
		log.Fatal(err)
	}

	collection := client.Database("Users").Collection("All")

	filter := bson.D{{"RefreshToken", rt}}

	var result Users
	var a int
	a = 1

	err = collection.FindOne(context.TODO(), filter).Decode(&result)
	if err != nil {
		a = 0
		log.Fatal(err)
	}

	fmt.Printf("Found a single document: %+v\n", result)

	return a
}