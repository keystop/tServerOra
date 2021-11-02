package signal

import (
	"context"
	"os"
	"os/signal"
	"syscall"
)

// слушаем в горутине(отдельном потоке) сигнал операционки на выход Ctrl+C
// При получении сигнала завершаем контекст через перданную в него cancel функцию
func HandleQuit(cancel context.CancelFunc) {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	<-c
	cancel()
}
