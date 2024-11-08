package gin

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/rebirthmonkey/ops/pkg/log"
	"golang.org/x/sync/errgroup"
	"net/http"
)

type Server struct {
	*Config

	*http.Server
	*gin.Engine
}

func New() (*Server, error) {
	opts := NewOptions()
	config := NewConfig()

	if err := opts.ApplyTo(config); err != nil {
		log.Errorln("New ApplyTo Config Error: ", err)
		return nil, err
	}

	serverInstance, err := config.New()
	if err != nil {
		log.Errorln("New New Error: ", err)
		return nil, err
	}

	return serverInstance, nil
}

func (s *Server) Run() error {
	log.Infoln("[GinServer] Run")

	s.Server = &http.Server{
		Addr:    s.Address,
		Handler: s,
	}

	var eg errgroup.Group

	eg.Go(func() error {
		log.Infoln("[GinServer] Listen on HTTP: ", s.Address)

		s.Engine.Run("0.0.0.0:8888")
		if err := s.Server.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Errorln("failed to start Gin HTTP server: ", err.Error())
			return err
		}
		return nil
	})

	if err := eg.Wait(); err != nil {
		log.Infoln(err.Error())
	}
	return nil
}
