package hub

import (
	"github.com/gorilla/websocket"
)

type User struct {
	conn *websocket.Conn
}

type Admin struct {
	conn *websocket.Conn
}

type Hub struct {
	users []User
	admin Admin
}

func (hub *Hub) subscribeUser(conn *websocket.Conn)  {
	hub.users = append(hub.users, User{conn: conn})
}
