package mlibs

import (
	"context"
	"github.com/pkg/errors"
	"golang.org/x/sync/errgroup"
	//	"net/http"
	"os"
	"os/signal"
	"syscall"

	//	"mlibs/handlers"
	"mlibs/transport"
)

type App struct {
	ctx     context.Context
	cancel  func()
	Name    string
	servers []transport.Server
}

func NewApp(opts ...Option) *App {
	ctx, cancel := context.WithCancel(context.Background())
	app := App{
		ctx:    ctx,
		cancel: cancel,
	}
	for _, o := range opts {
		o(&app)
	}

	return &app
}

func (a *App) Run() error {
	// all server go to run
	errgo, ctx := errgroup.WithContext(a.ctx)

	for _, tsrv := range a.servers {
		srv := tsrv
		errgo.Go(func() error {
			return srv.Start()
		})
		errgo.Go(func() error {
			<-ctx.Done()
			return srv.Stop()
		})
	}

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
		return errors.Wrap(err, "app run stoped")
	}

	return nil
}

func (a *App) Stop() error {

	if a.cancel != nil {
		a.cancel()
	}

	return nil
}
