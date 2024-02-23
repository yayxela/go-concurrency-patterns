package task

import (
	"net/http"
	"webinar/microservice/task-service/api"
	"webinar/microservice/task-service/utils"
)

func (c *Client) Get(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}

	_, err := c.jwt.ValidateToken(r.Header.Get("Authenticate"))
	if err != nil {
		utils.ServeError(w, err, http.StatusBadRequest)
		return
	}

	fail := 0
	success := 0
	for data := range MakeRequest("https://httpbin.org/get", 10) {
		if data.Err != nil {
			fail++
		}
		if data.Response != nil {
			success++
		}
	}

	if err != nil {
		utils.ServeError(w, err, http.StatusBadRequest)
		return
	}
	utils.ServeResponse(w, api.GetResponse{
		Success: success,
		Fail:    fail,
	})
}
