package main

import (
	"fmt"
	"go_practice/config"
	routes "go_practice/routers"
	"log"
	"net/http"
	"os"

	socketio "github.com/googollee/go-socket.io"
)

func main() {
	db := config.SetupDB()
	// db.AutoMigrate(&models.Task{})

	r := routes.SetupRoutes(db)
	r.Run(":" + os.Getenv("PORT"))

	server := socketio.NewServer(nil)

	server.OnConnect("/", func(s socketio.Conn) error {
		s.SetContext("")
		fmt.Println("connected:", s.ID())
		return nil
	})

	server.OnEvent("/", "notice", func(s socketio.Conn, msg string) {
		fmt.Println("notice:", msg)
		s.Emit("reply", "have "+msg)
	})

	server.OnEvent("/chat", "msg", func(s socketio.Conn, msg string) string {
		s.SetContext(msg)
		return "recv " + msg
	})

	server.OnEvent("/", "bye", func(s socketio.Conn) string {
		last := s.Context().(string)
		s.Emit("bye", last)
		s.Close()
		return last
	})

	server.OnError("/", func(s socketio.Conn, e error) {
		fmt.Println("meet error:", e)
	})

	server.OnDisconnect("/", func(s socketio.Conn, reason string) {
		fmt.Println("closed", reason)
	})

	go server.Serve()
	defer server.Close()

	http.Handle("/socket.io/", server)
	fs := http.FileServer(http.Dir("static"))
	http.Handle("/", fs)

	log.Println("Serving at localhost:8000...")
	log.Fatal(http.ListenAndServe(":"+os.Getenv("PORT"), nil))

}
