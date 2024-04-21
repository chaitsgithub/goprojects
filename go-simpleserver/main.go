package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", handlerfunc)
	fmt.Println("Server listening on port 8080")
	http.ListenAndServe(":8080", nil)
}

func handlerfunc(w http.ResponseWriter, r *http.Request) {
	// Get all cookies
	cookies := r.Cookies()
	fmt.Println("Getting Cookies!")
	// Loop through cookies
	for _, cookie := range cookies {
		name := cookie.Name
		value := cookie.Value
		// Access other cookie attributes like Expires, Path, etc. (optional)
		fmt.Fprintf(w, "Cookie name: %s, value: %s\n", name, value)
	}

	cookie := http.Cookie{
		Name:     "exampleCookie",
		Value:    "Hello world!",
		Path:     "/",
		MaxAge:   3600,
		HttpOnly: true,
		Secure:   true,
		SameSite: http.SameSiteLaxMode,
	}
	http.SetCookie(w, &cookie)

	fmt.Fprintf(w, "Hello World")
}
