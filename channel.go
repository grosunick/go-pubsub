package pubsub

// Максимальное количество не доставленных в канале сообщений
const DEFAULT_CHANNEL_SIZE = 50

// Структура, описывающиая канал
type Channel struct {
	// Канал, куда будут оправляться сообщения, которые необходимо передать подписчикам subscribers
	channel chan string
	// Подписчики канала
	subscribers *ChannelListeners
	// Брокер подписки, где данный канал будет храниться
	broker IBroker
}

// Фабрика создания нового канала
func NewChannel(props *ChannelProps, broker IBroker) *Channel {
	var ch chan string

	if props.MaxQueueSize == 0 {
		ch = make(chan string, DEFAULT_CHANNEL_SIZE)
	} else if props.MaxQueueSize > 0 {
		ch = make(chan string, props.MaxQueueSize)
	}

	channel := &Channel{
		channel: ch,
		subscribers: NewChannelListenersList(props),
		broker: broker,
	}

	return channel
}

// Возвращает структуру, описывающую слушателей канала
func (this *Channel) GetSubscribers() *ChannelListeners {
	return this.subscribers
}

// Рассылает сообдщение message слушателям, подписанным на канал
func (this *Channel) Publish(message string) {
	this.channel <- message
}

// Обработчик, реализующий прослушивание канала сообщений
func (this *Channel) Observe() {
	defer close(this.channel)

	for {
		select {
		case message := <-this.channel:
			// условаие выхода
			if message == "system:exit" {
				return
			}

			// отправляем сообщение подписчикам
			for listener, _ := range this.subscribers.listeners {
				listener.Send(message)
			}
		}
	}
}

// Остановить отправку сообщений подписчикам, очистить ресурсы
func (this *Channel) Stop() {
	this.channel <- "system:exit"
}
