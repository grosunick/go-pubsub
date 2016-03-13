package pubsub

import "fmt"

// Слушатель, отправляющий сообщение в консоль
type ListenerEcho struct {
	// Количество отправленных сообщений
	sendMessaged int
}

func (this *ListenerEcho) Send(message interface{}) {
	fmt.Println(message.(string))
}
