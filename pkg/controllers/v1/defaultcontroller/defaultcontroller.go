package defaultcontroller

import (
	"log"
	"net/http"

	"github.com/bungysheep/news-api/pkg/controllers/v1/basecontroller"
)

// DefaultController type
type DefaultController struct {
	basecontroller.BaseResource
}

// NewDefaultController - Creates default controller
func NewDefaultController() *DefaultController {
	return &DefaultController{}
}

// GetHome - Return home
func (defCtl *DefaultController) GetHome(w http.ResponseWriter, r *http.Request) {
	log.Printf("Retrieving home.\n")

	defCtl.WriteResponse(w, http.StatusOK, true, nil, "News Api v1")
}
