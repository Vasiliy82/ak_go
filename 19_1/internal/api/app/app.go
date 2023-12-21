package app

import (
	"context"
	"fmt"
	"lesson/internal/api/controller"
	"lesson/internal/api/repository"
	"lesson/internal/api/usecase"
	pb "lesson/pkg/report"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"golang.org/x/sync/errgroup"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func Start(ctx context.Context, addr, popularityAddr string) {
	ctx, stop := signal.NotifyContext(ctx, syscall.SIGINT, syscall.SIGTERM)
	defer stop()
	l := log.New(os.Stdout, "", 0)
	reportApi, err := initPopularityRepository(popularityAddr)
	if err != nil {
		return
	}
	popularityRepository := repository.NewPopularity(l, reportApi)
	hbUsecase := usecase.NewHonorBoard(popularityRepository)
	c := controller.NewController(hbUsecase, l)
	mux := route(c)
	logHandler := logMiddleware(l)
	httpServer := &http.Server{
		Addr: addr,
		Handler: authHandler(
			logHandler(mux),
		),
	}
	g, ctx := errgroup.WithContext(ctx)
	g.Go(func() error {
		log.Printf("старт сервиса (%s)\n", addr)
		return httpServer.ListenAndServe()
	})
	g.Go(func() error {
		<-ctx.Done()
		err := httpServer.Shutdown(context.Background())
		if err != nil {
			return err
		}
		log.Printf("завершение работы сервиса (%s)\n", addr)
		return nil
	})
	if err := g.Wait(); err != nil {
		if err == http.ErrServerClosed {
			return
		}
		log.Println(fmt.Errorf("ошибка сервиса (%s): %w", addr, err))
	}
}
func initPopularityRepository(addr string) (pb.ReportClient, error) {
	opts := []grpc.DialOption{
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	}
	conn, err := grpc.Dial(addr, opts...)
	if err != nil {
		log.Println("Не удалось подключиться (%s): %v", addr, err)
		return nil, err
	}
	return pb.NewReportClient(conn), err
}
