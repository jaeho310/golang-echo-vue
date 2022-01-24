package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// /static/html에 spa 프레임워크가 만들어준 하나의 html을 넣어주고
	// "/" 요청에 해당파일을 내려줍니다.
	e.File("/", "static/index.html")

	// static 파일을 찾는 루트경로를 static으로 설정합니다.
	e.Static("/", "static")

	// vue router를 사용하려는 요청이 백엔드에 넘어왔다면 404에러처리하면 안됩니다.
	// 백엔드에서 404에러가 발생한경우 프론트엔드로 넘깁니다.
	echo.NotFoundHandler = func(c echo.Context) error {
		return c.File("static/index.html")
	}
	e.Logger.Fatal(e.Start(":8083"))
}
