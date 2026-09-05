package chapter07

import "fmt"

type User struct {
	UserId string
	Inbox  []string
}

func NewUser(userId string) *User {
	if userId == "" {
		panic("user_id required")
	}
	return &User{UserId: userId}
}

type ChatMessage struct {
	SenderId string
	Message  string
}

type ChatRoom struct {
	Name    string
	members map[string]*User
	log     []ChatMessage
}

func NewChatRoom(name string) *ChatRoom {
	if name == "" {
		panic("name required")
	}
	return &ChatRoom{Name: name, members: map[string]*User{}}
}

func (r *ChatRoom) Join(user *User) {
	if user == nil {
		panic("user is nil")
	}
	r.members[user.UserId] = user
}

func (r *ChatRoom) Leave(userId string) { delete(r.members, userId) }

func (r *ChatRoom) Send(senderId, message string) {
	if _, ok := r.members[senderId]; !ok {
		panic("sender not in room")
	}
	if message == "" {
		panic("message required")
	}
	r.log = append(r.log, ChatMessage{SenderId: senderId, Message: message})
	for uid, user := range r.members {
		if uid != senderId {
			user.Inbox = append(user.Inbox, fmt.Sprintf("[%s] %s: %s", r.Name, senderId, message))
		}
	}
}

func (r *ChatRoom) History() []ChatMessage {
	return append([]ChatMessage{}, r.log...)
}

type ChatServer struct {
	rooms map[string]*ChatRoom
}

func NewChatServer() *ChatServer { return &ChatServer{rooms: map[string]*ChatRoom{}} }

func (s *ChatServer) CreateRoom(name string) *ChatRoom {
	if _, ok := s.rooms[name]; ok {
		panic("room exists")
	}
	room := NewChatRoom(name)
	s.rooms[name] = room
	return room
}

func (s *ChatServer) GetRoom(name string) *ChatRoom {
	room, ok := s.rooms[name]
	if !ok {
		panic("room not found")
	}
	return room
}

func RunQ707() {
	server := NewChatServer()
	room := server.CreateRoom("general")
	alice := NewUser("alice")
	bob := NewUser("bob")
	room.Join(alice)
	room.Join(bob)
	room.Send("alice", "hello")
	fmt.Println(bob.Inbox)
}
