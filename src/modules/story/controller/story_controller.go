package controller

import (
	"github.com/kataras/iris"
	"github.com/kataras/iris/mvc"

	"socmed/src/modules/story/usecase"
)

type StoryController struct {
	Ctx iris.Context

	StoryUsecase usecase.StoryUsecase
}

//GetReturn localhost:3000
func (c *StoryController) Get() mvc.Result {
	stories, err := c.StoryUsecase.GetAll()

	if err != nil {
		return mvc.View{
			Name: "index.html",
			Data: iris.Map{"Title": "Stories"},
		}
	}

	return mvc.View{
		Name: "index.html",
		Data: iris.Map{"Title": "Stories", "Stories": stories},
	}
}
