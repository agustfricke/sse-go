package events

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

type cliente struct {
  ID string
  sendMessage chan EventMessage
}

func newClient(id string) *cliente {
  return &cliente{
    ID: id,
    sendMessage: make(chan EventMessage),
  }
}

func (c *cliente) OnLine(ctx context.Context, w io.Writer, flush http.Flusher) {
  for {
    select {
    case m := <- c.sendMessage:
      data, err := json.Marshal(m.Data)
      if err != nil {
        log.Print(err)
      }
      const format = "event:%s\ndata:%s\n\n"
      fmt.Fprintf(w, format, m.EventName, string(data))
      flush.Flush()

    case <- ctx.Done():
      return
    }

  }
}

