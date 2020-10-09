package goRotines

import (
	"fmt"

	"o2b.com.br/WhatsAppProcessWorker/domain"

	"o2b.com.br/WhatsAppProcessWorker/domain/entities"
)

// Worker is my worker
type Worker struct {
	Message  *entities.Message
	WorkPile *WorkPile
	Http     *Http
}

func (w *Worker) putOnDonePile() {
	processing := "processing"
	doneMessage := domain.JSONStringfy(&entities.Message{
		ID:   w.Message.ID,
		Body: &processing,
	})
	fmt.Println(doneMessage)
	w.Http.Post(doneMessage)
}

// Process is my worker process
func (w *Worker) Process() {

	w.putOnDonePile()
}

// NewWorker constructor
func NewWorker(_message *entities.Message, _workPile *WorkPile) *Worker {
	return &Worker{
		Message:  _message,
		WorkPile: _workPile,
		Http:     NewHttp("http://127.0.0.1:52387/sync"),
	}
}
