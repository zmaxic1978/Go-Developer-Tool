package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	// создаем файл в текущей директории
	nf, err := os.OpenFile("logfile", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Panicln("не могу создать файл для логов:", err.Error())
	}
	defer nf.Close()

	mux := http.NewServeMux()
	mux.HandleFunc("/hello", hello)
	log.SetOutput(nf) // <--- указываем куда выводить логи
	l := log.New(nf, "", 0)
	logHandler := logMiddleware(l)
	httpServer := &http.Server{Addr: ":50000", Handler: logHandler(authHandler(mux))} // <--- если хотим сначала логировать, а потом проверять авторизацию
	if err := httpServer.ListenAndServe(); err != nil {
		log.Fatalln(fmt.Errorf("не удалось запустить сервер: %w", err))
	}
}

func hello(res http.ResponseWriter, req *http.Request) {
	msg := "Hello, Go!"
	log.Println("resp:", msg)
	fmt.Fprint(res, msg)
}

func authHandler(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		xId := r.Header.Get("x-my-app-id")
		if xId != "my_secret" {
			http.Error(w, "пользователь не авторизован", http.StatusUnauthorized)
			return
		}
		h.ServeHTTP(w, r)
	})
}

func logMiddleware(l *log.Logger) func(h http.Handler) http.Handler {
	return func(h http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			log.Println("url:", r.URL)
			h.ServeHTTP(w, r)
		})
	}
}
