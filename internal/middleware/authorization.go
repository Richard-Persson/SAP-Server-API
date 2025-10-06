
package middleware


import(

	"errors"
	"net/http"

)


var UnAuthorizedError = errors.New("Invalid username")


func Authorization(next http.Handler) http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request){




	})


}
