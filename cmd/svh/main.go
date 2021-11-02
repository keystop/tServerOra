package main

import (
	"context"
	"fmt"

	"tServerOra/internal/defoptions"
	"tServerOra/internal/repository"
	"tServerOra/internal/server"
	"tServerOra/internal/signal"
)

// Main.
func main() {
	// Создаем глобальный контекст выполнения
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	opt := defoptions.NewDefOptions()
	sr, err := repository.NewServerRepo(ctx, opt.DBConnString())
	if err != nil {
		fmt.Println("Ошибка при подключении к БД: ", err)
		return
	}
	defer sr.Close()
	s := new(server.Server)
	go signal.HandleQuit(cancel)
	go s.Start(ctx, sr, opt)
	<-ctx.Done()
}
