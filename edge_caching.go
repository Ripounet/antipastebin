package antipastebin

import (
	"appengine"
	"net/http"
)
// Neither of these combinations work!!
// This might be so because
// "the edge cache is only available for Google Apps domains, no effect on appspot.com" 
// (See https://groups.google.com/d/msg/google-appengine/8QgEUBOiNFw/m4O5quSO8q0J )
//
func init() {
	http.HandleFunc("/resource2", resource2)
	http.HandleFunc("/resource3", resource3)
	http.HandleFunc("/resource4", resource4)
	http.HandleFunc("/resource5", resource5)
	http.HandleFunc("/resource6", resource6)
	http.HandleFunc("/resource7", resource7)
}

func resource2(w http.ResponseWriter, r *http.Request) {
	c := appengine.NewContext(r)
	c.Infof("%v", r)
	
	w.WriteHeader(http.StatusOK)  // 200
	w.Write([]byte("This is resource 2."))
}

func resource3(w http.ResponseWriter, r *http.Request) {
	c := appengine.NewContext(r)
	c.Infof("%v", r)
	
	w.Header()["Cache-control"] = []string{ "private, max-age=30" }
	w.WriteHeader(http.StatusOK)  // 200
	w.Write([]byte("This is resource 3."))
}

func resource4(w http.ResponseWriter, r *http.Request) {
	c := appengine.NewContext(r)
	c.Infof("%v", r)
	
	w.Header()["Cache-control"] = []string{ "public, max-age=30" }
	w.WriteHeader(http.StatusOK)  // 200
	w.Write([]byte("This is resource 4."))
}

func resource5(w http.ResponseWriter, r *http.Request) {
	c := appengine.NewContext(r)
	c.Infof("%v", r)
	
	w.Header()["Cache-control"] = []string{ "public, max-age=300" }
	w.WriteHeader(http.StatusOK)  // 200
	w.Write([]byte("This is resource 5."))
}

func resource6(w http.ResponseWriter, r *http.Request) {
	c := appengine.NewContext(r)
	c.Infof("%v", r)
	
	w.Header()["Cache-control"] = []string{ "public, max-age=300" }
	w.Header()["Pragma"] = []string{ "public, max-age=300" }
	w.WriteHeader(http.StatusOK)  // 200
	w.Write([]byte("This is resource 6."))
}

func resource7(w http.ResponseWriter, r *http.Request) {
	c := appengine.NewContext(r)
	c.Infof("%v", r)
	
	w.Header()["Cache-control"] = []string{ "public, max-age=300" }
	w.Header()["Pragma"] = []string{ "public" }
	w.WriteHeader(http.StatusOK)  // 200
	w.Write([]byte("This is resource 7."))
}