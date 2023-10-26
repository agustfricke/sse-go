package events

import (
	"fmt"
	"net/http"
	"sync"
)

type EventMessage struct {
  EventName string
  Data any
}

type HandlerEvent struct {
  m sync.Mutex
  clientes map[string]*cliente
}

func NewHandlerEvent() *HandlerEvent{
  return &HandlerEvent{
    clientes: make(map[string]*cliente),
    // todo
  }
}

func (h *HandlerEvent) Handler (w http.ResponseWriter, r *http.Request) {
  w.Header().Set("Content-Type", "text/event-stream")
  w.Header().Set("Cache-Control", "no-cache")
  w.Header().Set("Connection", "keep-alive")

  id := r.URL.Query().Get("id")
  if id == "" {
    fmt.Println("No hay id")
    return
  }
  
  flusher, ok := w.(http.Flusher)
  if !ok {
    w.WriteHeader(http.StatusBadRequest)
    return 
  }

  c := newClient(id)
  h.register(c)
  fmt.Println("Connected:", id)
  c.OnLine(r.Context(), w, flusher)
  fmt.Println("Desconnected:", id)
  h.remove(id)
}

func (h *HandlerEvent) register(c *cliente) {
  h.m.Lock()
  defer h.m.Unlock()
  h.clientes[c.ID] = c

}

func (h *HandlerEvent) remove (id string) {
  h.m.Lock()
  delete(h.clientes, id)
  defer h.m.Unlock()
}

func (h *HandlerEvent) Brodcast (m EventMessage) {
  h.m.Lock()
  defer h.m.Unlock()
  for _, c := range h.clientes {
    c.sendMessage <- m
  }
}
