package app

import (
	"context"
	"fmt"
	"lesson/internal/report/controller"
	pb "lesson/pkg/report"
	"log"
	"net"
	"os/signal"
	"syscall"

	"golang.org/x/sync/errgroup"
	"google.golang.org/grpc"
)

func Start(ctx context.Context, addr string) {
	ctx, stop := signal.NotifyContext(ctx, syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	g, ctx := errgroup.WithContext(ctx)
	lis, err := net.Listen("tcp", addr)
	if err != nil {
		log.Println(err)
	}
	s := grpc.NewServer()
	pb.RegisterReportServer(s, new(controller.Controller))
	g.Go(func() error {
		log.Printf("старт сервиса (%s)\n", addr)
		if err = s.Serve(lis); err != nil {
			log.Println(fmt.Errorf("ошибка сервиса (%s): %w", addr,
				err))
		}
		return nil
	})
	g.Go(func() error {
		<-ctx.Done()
		s.GracefulStop()
		log.Printf("завершение работы сервиса (%s)\n", addr)
		return nil
	})
	if err = g.Wait(); err != nil {
		log.Println(fmt.Errorf("ошибка сервиса (%s): %w", addr, err))
	}
}
