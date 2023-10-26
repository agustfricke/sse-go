package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/agustfricke/sse-go/handlers/events"
)

func InitRoutes(r *http.ServeMux) {
  handlerEvents := events.NewHandlerEvent()

  r.HandleFunc("/noti", handlerEvents.Handler)
  r.HandleFunc("/test1", HandlerTest1(handlerEvents))
  r.HandleFunc("/test2", HandlerTest2(handlerEvents))
  r.Handle("/", http.FileServer(http.Dir("./public")))

}

func HandlerTest1(notifier *events.HandlerEvent) http.HandlerFunc{
  return func(w http.ResponseWriter, r *http.Request) {
    var data = map[string]any{}
    json.NewDecoder(r.Body).Decode(&data)
    notifier.Brodcast(events.EventMessage{
      EventName: "fii",
      Data: data,
    })
  }
}

func HandlerTest2(notifier *events.HandlerEvent) http.HandlerFunc{
  return func(w http.ResponseWriter, r *http.Request) {
    var data = map[string]any{}
    json.NewDecoder(r.Body).Decode(&data)
    notifier.Brodcast(events.EventMessage{
      EventName: "foo",
      Data: data,
    })
  }
}
