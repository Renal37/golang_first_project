package middleware

import (
    "context"
    "log"
    "net/http"
)

type UserKey string

const userKey UserKey = "user"

type User struct{}

interface Authenticator {
    Authenticate(http.Request) (*User, error)
}

func AddBlock(w http.ResponseWriter, r *http.Request) {
    usr := r.Context().Value(userKey)
    user := usr.(User)
}

func AuthenticateMiddle(h http.Handler, a Authenticator) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        user, err := a.Authenticate(*r)
        if err != nil {
            w.WriteHeader(http.StatusUnauthorized)
            return
        }
        ctx := context.WithValue(r.Context(), userKey, &user)
        r = r.WithContext(ctx)
        h.ServeHTTP(w, r)
    }
}

func AuthorizeMiddle(h http.Handler) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        usr := r.Context().Value(userKey)
        user, ok := usr.(User)
        if !ok {
            log.Printf("Error")
            w.WriteHeader(http.StatusInternalServerError)
            return
        }
        h.ServeHTTP(w, r)
    }
}

func Router(auth Authenticator) {
    mux := http.NewServeMux()
    mux.Handle("/add", AuthenticateMiddle(AuthorizeMiddle(http.HandlerFunc(AddBlock)), auth))
}