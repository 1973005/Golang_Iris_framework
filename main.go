package main

import (
	"os"
	"time"

	"github.com/kataras/iris"
	"github.com/kataras/iris/mvc"
	"github.com/kataras/iris/sessions"

	"socmed/src/config"
	"socmed/src/modules/profile/controller"
	"socmed/src/modules/profile/repository"
	"socmed/src/modules/profile/usecase"
)

func main() {
	app := iris.New()

	app.Logger().SetLevel("debug")

	views := iris.HTML("./web/views", ".html").Layout("layout.html").Reload(true)

	app.RegisterView(views)
	app.StaticWeb("/public", "./web/public")

	app.Get("/", func(ctx iris.Context) {
		ctx.ViewData("Message", "Golang Socmed")
		ctx.View("index.html")
	})

	db, err := config.GetMongoDB()

	if err != nil {
		os.Exit(2)
	}
	profileRepository := repository.NewProfileRepositoryMongo(db, "profile")
	profileUsecase := usecase.NewProfileUsecase(profileRepository)

	sessionManager := sessions.New(sessions.Config{
		Cookie:  "cookiename",
		Expires: time.Minute * 10,
	})

	profile := mvc.New(app.Party("/profile"))
	profile.Register(profileUsecase, sessionManager.Start)

	profile.Handle(new(controller.ProfileController))

	app.Run(iris.Addr(":3000"))
}
