package delivery

import (
	"encoding/json"
	"net/http"
	"smtp/app/sender/dto"
	"smtp/app/sender/services"
	"smtp/pkg/httpext"
)

type Controller struct {
	service services.Sender
}

func NewController(service services.Sender) *Controller {
	return &Controller{
		service: service,
	}
}

func (c *Controller) Send(w http.ResponseWriter, r *http.Request) {
	var payload dto.SendMailPayload

	if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
		httpext.AbortJSON(w, err.Error(), http.StatusBadRequest)
	}

	err := c.service.Send(payload)

	if err != nil {
		httpext.AbortJSON(w, err.Error(), http.StatusNonAuthoritativeInfo)
	}

	httpext.EmptyResponse(w, http.StatusOK)
}
