package server

import (
	"fmt"
	"net/http"
	"time"

	log "github.com/leodotcloud/log"
)

// Server structure is used to the store backend information
type Server struct {
}

// ListenAndServe is used to setup ping and reload handlers and
// start listening on the specified port
func (s *Server) ListenAndServe(listen string) error {
	go func(id string) {
		num := 0
		for {
			time.Sleep(time.Millisecond * 1000)
			log.Infof("%v :: Number: %v", id, num)
			num++
			time.Sleep(time.Millisecond * 1000)
			log.Errorf("%v :: Number: %v", id, num)
			num++
			time.Sleep(time.Millisecond * 1000)
			log.Debugf("%v :: Number: %v", id, num)
			num++
		}
	}("server")

	http.HandleFunc("/ping", s.ping)
	http.HandleFunc("/v1/loglevel", s.loglevel)
	log.Infof("Listening on %s", listen)
	err := http.ListenAndServe(listen, nil)
	if err != nil {
		log.Errorf("got error while ListenAndServe: %v", err)
	}
	return err
}

func (s *Server) ping(rw http.ResponseWriter, req *http.Request) {
	log.Debugf("Received ping request")
	rw.Write([]byte("OK"))
}

func (s *Server) loglevel(rw http.ResponseWriter, req *http.Request) {
	// curl -X POST -d "level=debug" localhost:12345/v1/loglevel
	log.Debugf("Received loglevel request")
	if req.Method == http.MethodGet {
		level := log.GetLevel().String()
		rw.Write([]byte(fmt.Sprintf("%s\n", level)))
	}

	if req.Method == http.MethodPost {
		if err := req.ParseForm(); err != nil {
			rw.WriteHeader(http.StatusInternalServerError)
			rw.Write([]byte(fmt.Sprintf("Failed to parse form: %v\n", err)))
		}
		level, err := log.ParseLevel(req.Form.Get("level"))
		if err != nil {
			rw.WriteHeader(http.StatusInternalServerError)
			rw.Write([]byte(fmt.Sprintf("Failed to parse loglevel: %v\n", err)))
		} else {
			log.SetLevel(level)
			rw.Write([]byte("OK\n"))
		}
	}
}
