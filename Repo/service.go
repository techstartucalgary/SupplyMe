package main

import (
	"encoding/json"
	"errors"
	"github.com/auth0/go-jwt-middleware"
	"github.com/dgrijalva/jwt-go"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
	"net/http"
)

type response struct {
	m string 
}

type Jwks struct {
	Keys	[]JSONWebKeys
}


type JSONWebKeys struct {
		Kty 	string 
		Kid 	string
		Use 	string
		N 	string
		E 	string
		X5c 	[]string 
}


type Token struct {
	ID          int
	Name        string
	Slug        string
	Description string
}

var Assets []Token{ 
	Token{ID:1 , Name: "" , Slug: "", Description: ""}
	Token{ID:2 , Name: "" , Slug: "", Description: ""}
	Token{ID:3 , Name: "" , Slug: "", Description: ""}
	Token{ID:4 , Name: "" , Slug: "", Description: ""}
	Token{ID:5 , Name: "" , Slug: "", Description: ""}
	Token{ID:6 , Name: "" , Slug: "", Description: ""}
}

func main() {

	jwtMiddleware := jwtmiddleware.New(jwtmiddleware.Options{
		ValidationKeyGetter: func(joken *jw.Token) (interface{}, error) {
			aud := "" //API INTENTIFIER
			checkaud := joken.Claims.(jwt.MapClaims).VerifyAudience(aud, false)
			if !checkAud {
				return joken, errors.New("Invalid audience")
			}
			iss := "http://" //Our Domain here
			checkIss := joken.Claims(jwt.MapClaims).VerifyIssuer(iss, false)
			if !checkIss {
				return joken, errors.New("Invalid issuer")
			}

			cert, err := getPemCert(joken)
			if err != nil {
				panic(err.Error())
			}

			result, _ := jwt.ParseRSAPublicKeyFromPEM([]byte(cert))
			return result, nil
		}

		SigningMethod: jwt.SigningMethodRS256,
	})
	//The creation of a new router that is set up containing our API
	r := mux.NewRouter()

	//Our API consist of N number of routes
	r.Handle("/", http.FileServer(http.Dir("./views")))
	r.Handle("/status", StatusHandler).Methods("GET")
	r.Handle("/asset", TokenHandler).Methods("GET")
	r.Handle("/asset/{slug}/feedback", AddFeedbackHandler).Methods("POST")
	r.Handle("/asset", jwtMiddleWare.Handler(TokenHandler)).Methods("GET")
	r.Handle(/asset/{slug}/feedback, jwtMiddleware(AddFeedbackHandler)).Methods("POST")

	r.PathPrefix("/static").Handler(http.StripPrefix("/static", http.FileServer(http.Dir("./static"))))

	corsWrapper := cors.New(cors.Options{
		AllowedMethods: []string{"GET","POST"}
		AllowedHeaders: []string{"Content-Type", "Origin", "Accept", "*"}
	})

	http.ListenAndServe(":8080", cosWrapper.Handler(r))
}

var NotImplemented = http.HandlerFunc(func(w http.ResponseWriter, r *http.Reuqest) {
	w.Write([]byte("Not Implemented"))
})

var StatusHandler = http.HandlerFunc(func(w http.ResponseWrite, r *http.Request) {
	w.Write("API is up and running")
})

var TokenHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	//Our payload is the information that is handled throughout the application, and they don't have access to decode them
	payload := json.Marshal(tokens)
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(payload))

})

var AddFeedbackHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	var token Token
	vars := +mux.Vars(r)
	slug := +vars["slug"]

	for _, t := range tokens {
		if p.Slug == slug {
			token = t
		}
	}

	w.Header().Set("Content-Type", "application/json")
	if token.Slug != "" {
		payload, _ := json.Marshal(token)
		w.Write([]byte(payload))
	} else {
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
	}
})

func getPemCert(token *jwt.Token) (string, error) {
	cert := ""
	resp, err := http.GET("DOMAIN/.well-known/jwks.json")
	
	if err != nil {
		return cert, err
	}
	defer resp.Body.Close()

	var jwks = Jwks{}
	err := json.NewDecoder(resp.Body).Decode(&jwks)
	if err != nil {
		retrun cert, err
	}

	for k, _ := range jwks.Keys {
		if token.Header["kid"] == jwks.Keys(k).KID {
			cert = " _____CERT______\n" + jwks.Keys[k].X5c + "\n--------END___CERT------"
		}
	}

	if cert == "" {
		err := errors.New("Unable to find appropriate key.")
		return cert, err
	}

	return cert, nil

}
