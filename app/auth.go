package app

import (
	"bankapp/utils"
	"net/http"
	"time"
	"bankapp/jwtmiddleware"
	"github.com/dgrijalva/jwt-go"
)


/* Set up a global string for our secret */
var mySigningKey = []byte("secret")

// GetTokenHandler provides authentication
var GetTokenHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	/* Create the token */
	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)

	/* Set token claims */
	claims["admin"] = true

	//Setting 5 day expiration time
	claims["exp"] = time.Now().Add(time.Hour * 24 * 5).Unix()

	/* Sign the token with our secret */
	tokenString, _ := token.SignedString(mySigningKey)
	message := make(map[string]string)
	message["token_string"]=tokenString
	utils.RespondWithJSON(w, http.StatusFound, message)

})



//ModJWTHandler is to
func ModJWTHandler(h http.Handler) http.Handler {
//JwtMiddleware for validating jwt tokens
var JwtMiddleware = jwtmiddleware.New(jwtmiddleware.Options{
	ValidationKeyGetter: func(token *jwt.Token) (interface{}, error) {
		return mySigningKey, nil
	},
	SigningMethod: jwt.SigningMethodHS256,
})
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Let secure process the request. If it returns an error,
		// that indicates the request should not continue.
		err := JwtMiddleware.CheckJWT(w, r)
		w.(http.Flusher).Flush()

		// If there was an error, do not continue.
		if err != nil {
			utils.RespondWithError(w, http.StatusUnauthorized, err.Error())
			return
		}

		h.ServeHTTP(w, r)
	})
}
