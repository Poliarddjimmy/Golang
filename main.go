package main
import (
"log"
"net/http"
)
func home(w http.ResponseWriter, r *http.Request) {
	// Check if the current request URL path exactly matches "/". If it doesn't, use
	// the http.NotFound() function to send a 404 response to the client.
	// Importantly, we then return from the handler. If we don't return the handler
	// would keep executing and also write the "Hello from SnippetBox" message.
	if r.URL.Path != "/" {
	http.NotFound(w, r)
	return
	}
	w.Write([]byte("Hello from Snippetbox"))
}
// Add a showSnippet handler function.
func showSnippet(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Display a specific snippet..."))
}
// Add a createSnippet handler function.
// func createSnippet(w http.ResponseWriter, r *http.Request) {
// 	w.Write([]byte("Create a new snippet..."))
// }


/////Post/////
func createSnippet(w http.ResponseWriter, r *http.Request) {
	// Use r.Method to check whether the request is using POST or not. Note that
	// http.MethodPost is a constant equal to the string "POST".
	if r.Method != http.MethodPost {
		// If it's not, use the w.WriteHeader() method to send a 405 status
		// code and the w.Write() method to write a "Method Not Allowed"
		// response body. We then return from the function so that the
		// subsequent code is not executed.
		w.WriteHeader(405)
		w.Write([]byte("Method Not Allowed"))
		return
	}
	w.Write([]byte("Create a new snippet..."))
}


func errorSnippet(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
	w.Header().Set("Allow", http.MethodPost)
	// Use the http.Error() function to send a 405 status code and "Method Not
	// Allowed" string as the response body.
	http.Error(w, "Method Not Allowed! Ok????", 405)
	return
	}
	w.Write([]byte("Error snippet..."))
}

// func main() {
// 	mux := http.NewServeMux()
// 	mux.HandleFunc("/", home)
// 	mux.HandleFunc("/snippet", showSnippet)
// 	mux.HandleFunc("/snippet/create", createSnippet)
// 	log.Println("Starting server on :4000")
// 	err := http.ListenAndServe(":4000", mux)
// 	log.Fatal(err)
// }

// var DefaultServeMux = NewServeMux()

func main() {
	http.HandleFunc("/", home)
	http.HandleFunc("/snippet", showSnippet)
	http.HandleFunc("/snippet/error", errorSnippet)
	http.HandleFunc("/snippet/create", createSnippet)
	log.Println("Starting server on :4000")
	err := http.ListenAndServe(":4000", nil)
	log.Fatal(err)
}