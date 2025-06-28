package api

import "net/http"

func HandlerHello(cfg *ApiConfig) {
	type response struct {
		MSG string `json:"msg"`
	}

	respondWithJSON(cfg.Resp, http.StatusOK, response{
		MSG: "hello there bob!",
	})
}

