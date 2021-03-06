package middleware

import (
	"encoding/json"
	"log"
	"net/http"
	"strings"

	"github.com/arthurkay/lucid/utils"
)

// CheckAuth takes in an handler and returns a handler, otherwise proceeds with processig
func CheckAuth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		bearerToken := r.Header.Get("Authorization")
		token := strings.Split(bearerToken, " ")

		// Make sure authorization is set in http headers
		if len(token) > 1 {
			ok, err := utils.VerifyToken(token[1])
			if err != nil {
				log.Printf("%v", err)
			}
			if ok {
				next.ServeHTTP(w, r)
			} else {
				resp := map[string]string{
					"status":  "failed",
					"message": "Unauthorized access denied",
				}
				data, er := json.Marshal(resp)
				if er != nil {
					log.Printf("%v", er)
				}
				w.Write(data)
			}
		} else {
			resp := map[string]string{
				"status":  "failed",
				"message": "No authorization parameters set",
			}
			data, er := json.Marshal(resp)
			if er != nil {
				log.Printf("%v", er)
			}
			w.Write(data)
		}
	})
}
