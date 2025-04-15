package main

import (
	"context"
	"fmt"
	"gews/model"
	site "gews/model/enum"
	"gews/service/tupiService"
)

type App struct {
	Tupi tupiService.Client
	ctx  context.Context
}

func NewApp() *App {
	return &App{
		Tupi: tupiService.New(),
	}
}

func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
}

func (a *App) Greet(name string) string {
	return fmt.Sprintf("Hello %s, It's show time!", name)
}

func (a *App) Teste() model.News {
	a.Tupi.News.Get()
	return model.News{
		Headline: "Teste",
	}
}

func (a *App) GetNews() []model.NewsSite {
	var newsSite []model.NewsSite
	news, err := a.Tupi.News.Get()
	if err != nil {
		return nil
	}
	newsSite = append(newsSite, model.NewsSite{
		SiteName: "Tupi",
		Site: site.Tupi,
		News:            news,
		Visible:         true})
	return newsSite
}
