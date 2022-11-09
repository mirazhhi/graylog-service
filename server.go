package web

import (
    "net/http"
    "time"
    "context"
)

type Server struct {
    httpServer *http.Server
}

func (s *Server) Run(port string) {
    s.httpServer = &http.Server{
        Addr:           ":" + port,
//         Handler:        myHandler,
        ReadTimeout:    10 * time.Second,
        WriteTimeout:   10 * time.Second,
        MaxHeaderBytes: 1 << 20,
    }

    s.httpServer.ListenAndServe()
}


func (s *Server) Shutdown(context context.Context) error {
    return s.httpServer.Shutdown(context)
}