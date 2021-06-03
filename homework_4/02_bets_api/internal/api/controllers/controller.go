package controllers


// Controller implements handlers for web server requests.
type Controller struct {
	betResponse BetResponse
}

// NewController creates a new instance of Controller
func NewController(betResponse BetResponse) *Controller {
	return &Controller{
		betResponse: betResponse,
	}
}
