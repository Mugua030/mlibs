package mlibs

import (
	"context"
	"github.com/pkg/errors"
	"golang.org/x/sync/errgroup"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"mlibs/handlers"
)

type App struct {
	ctx    context.Context
	cancel func()
}

func NewApp() *App {
	ctx, cancel := context.WithCancel(context.Background())
	return &App{
		ctx:    ctx,
		cancel: cancel,
	}
}

func (a *App) Run() error {
	// http server
	errgo, ctx := errgroup.WithContext(a.ctx)

	errgo.Go(func() error {
		//http.HandleFunc("/hello", handlers.HelloHandler)
		//return http.ListenAndServe(":8888", nil)
		muxSv := http.NewServeMux()
		muxSv.Handle("/hello", http.HandlerFunc(handlers.HelloHandler))

		addr := ":8888"
		srv := http.Server{
			Addr:    addr,
			Handler: muxSv,
		}
		go func() {
			<-ctx.Done()
			srv.Shutdown(context.Background())
		}()

		return srv.ListenAndServe()
	})
	// signal process
	stopCh := make(chan os.Signal, 1)
	errgo.Go(func() error {
		signal.Notify(stopCh, syscall.SIGINT, syscall.SIGTERM)
		for {
			select {
			case <-ctx.Done():
				return ctx.Err()
			case <-stopCh:
				a.Stop()
			}
		}
	})

	if err := errgo.Wait(); err != nil {
		return errors.New(err.Error())
	}

	return nil
}

func (a *App) Stop() error {

	if a.cancel != nil {
		a.cancel()
	}

	return nil
}
