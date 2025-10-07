
package auth


import (
	"errors"
	"net/http"

	"github.com/Richard-Persson/SAP-Server-API/internal/dto"

	log "github.com/sirupsen/logrus"
)



var UnAuthorizedError = errors.New("Invalid username or token")

func Authorization(next http.Handler) http.Handler{
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request){
		
		var username string = r.URL.Query().Get("username")
		var token = r.Header.Get("Authorization")
		err := error

		if username == "" || token = "" {

			return http.Handle(err)

		}
		
	})
}
