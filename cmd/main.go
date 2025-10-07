package main

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"time"

	httpImpl "reporteador/pkg/server/http"
	"reporteador/pkg/server/test/service"

	"github.com/go-chi/chi/v5"
	"go.uber.org/zap"
)

func main() {
	logger, _ := zap.NewProduction()
	defer logger.Sync()

	app, err := service.NewApplication()
	if err != nil {
		logger.Fatal("Error al inicializar la app", zap.Error(err))
	}

	httpServer := httpImpl.NewHttpServer(*app)

	router := chi.NewRouter()

	router.Post("/pdf", httpServer.GeneratePDF)
	router.Get("/pdf/view", httpServer.ViewPDF)
	router.Post("/pdf/ventas", httpServer.GenerateReporteVentas)
	router.Post("/pdf/inventario", httpServer.GenerateReporteInventario)
	router.Post("/pdf/usuarios-activos", httpServer.GenerateReporteUsuariosActivos)

	addr := ":8080"
	srv := &http.Server{
		Addr:    addr,
		Handler: router,
	}

	go func() {
		logger.Info("Servidor iniciado en " + addr)
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			logger.Fatal("Error en el servidor HTTP", zap.Error(err))
		}
	}()

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt)
	<-stop

	logger.Info("Apagando servidor...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		logger.Fatal("Error durante shutdown del servidor", zap.Error(err))
	}

	logger.Info("Servidor cerrado correctamente")
}
