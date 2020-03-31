package app

import (
	"net/http"
	"time"

	jwtmiddleware "github.com/auth0/go-jwt-middleware"
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

	// Finally, write the token to the browser window
	w.Write([]byte(tokenString))
})

//JwtMiddleware for validating jwt tokens
var JwtMiddleware = jwtmiddleware.New(jwtmiddleware.Options{
	ValidationKeyGetter: func(token *jwt.Token) (interface{}, error) {
		return mySigningKey, nil
	},
	SigningMethod: jwt.SigningMethodHS256,
})
