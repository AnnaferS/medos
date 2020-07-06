package main

import "github.com/labstack/echo"

type Users struct {
	GUID string `bson:"GUID"`
	RefTok string `bson:"RefreshToken"`
}

func main() {

	e := echo.New()

	h := &handler{}
	//Выдает пару токенов
	//localhost:8080/createTok?GUID=/*yourGUID*/
	e.POST("/createTok", h.login)

	//Удаляем конкретный Refresh токен из БД
	//localhost:8080/deleteToken?RefreshToken=/*Refresh токен, который хотим удалить*/
	e.DELETE("/deleteToken", h.deleteTok)

	//Удаляем все Refresh токены пользователя (с указанным GUID) из БД
	//localhost:8080/deleteAllTokens?GUID=/*yourGUID*/
	e.DELETE("/deleteAllTokens", h.deleteAllRTok)

	//Обновляем пару токенов
	//localhost:8080/refreshTokenPair?GUID=/*yourGUID*/&RefreshToken=/*yourRefresh*/
	e.PUT("/refreshTokenPair", h.refreshTokens)

	e.Logger.Fatal(e.Start(":8080"))

}