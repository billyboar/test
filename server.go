package developertest

import (
	"context"
	"net/http"
	"strconv"

	"github.com/billyboar/developer-test-1/externalservice"
	"github.com/labstack/echo"
)

var externalService = externalservice.ClientService{}

func newServer() *echo.Echo {
	e := echo.New()
	e.POST("/api/posts/:id", post)
	e.GET("/api/posts/:id", get)
	return e
}

func StartServer(e *echo.Echo) {
	e.Logger.Fatal(e.Start(":8080"))
}

func StopServer(e *echo.Echo) error {
	return e.Shutdown(context.Background())
}

func post(ctx echo.Context) error {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		return err
	}
	title := ctx.FormValue("title")
	description := ctx.FormValue("description")
	post := externalservice.Post{
		Title:       title,
		Description: description,
	}

	responsePost, err := externalService.POST(id, &post)
	if err != nil {
		return err
	}

	return ctx.JSON(http.StatusOK, responsePost)
}

func get(ctx echo.Context) error {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		return err
	}
	if _, err := externalService.GET(id); err != nil {
		errResponse := externalservice.Error{Code: http.StatusBadRequest, Message: err.Error()}
		return ctx.JSON(http.StatusBadRequest, errResponse)
	}
	return nil
}
