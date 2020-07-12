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

	storyCtrl "socmed/src/modules/story/controller"
	storyRepo "socmed/src/modules/story/repository"
	storyUc "socmed/src/modules/story/usecase"
)

func main() {
	app := iris.New()

	app.Logger().SetLevel("debug")

	views := iris.HTML("./web/views", ".html").Layout("layout.html").Reload(true)

	app.RegisterView(views)
	app.StaticWeb("/public", "./web/public")

	db, err := config.GetMongoDB()

	if err != nil {
		os.Exit(2)
	}
	//story
	storyRepository := storyRepo.NewStoryRepositoryMongo(db, "stories")
	storyUsecase := storyUc.NewStoryUsecase(storyRepository)

	//profile
	profileRepository := repository.NewProfileRepositoryMongo(db, "profile")
	profileUsecase := usecase.NewProfileUsecase(profileRepository)

	sessionManager := sessions.New(sessions.Config{
		Cookie:  "cookiename",
		Expires: time.Minute * 10,
	})

	//profile route
	profile := mvc.New(app.Party("/profile"))
	profile.Register(profileUsecase, sessionManager.Start)

	profile.Handle(new(controller.ProfileController))

	//story route
	story := mvc.New(app)
	story.Register(storyUsecase)

	story.Handle(new(storyCtrl.StoryController))

	app.Run(iris.Addr(":3000"))
}
