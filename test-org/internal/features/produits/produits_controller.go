package modules

import (
	"encoding/json"
	"net/http"
)

type ProduitsController struct {
	service ProduitsService
}

func NewProduitsController(service ProduitsService) *ProduitsController {
	return &ProduitsController{
		service: service,
	}
}

func (c *ProduitsController) HandleGet(w http.ResponseWriter, r *http.Request) {
	result := c.service.DoBusinessLogic()
	json.NewEncoder(w).Encode(map[string]string{"message": result})
}
