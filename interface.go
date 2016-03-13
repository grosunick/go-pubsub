package pubsub

// Интерфейс слушателя событий
type IListener interface {
	// Доставляет сообщение, поступившее из канала конкретному слушателю
	Send(message interface{})
}

// Интерфейс канала
type IChannel interface {
	// Рассылает сообщение message слушателям, подписанным на канал
	Publish(message string)
	// Останавливает прослушивание канала
	Stop()
	// Возвращает структуру, описывающую слушателей канала
	GetSubscribers() *ChannelListeners
}

// интерфейс брокера очередей
type IBroker interface {
	// Создает канал с названием name
	CreateChannel(name string, props *ChannelProps) IChannel
	// Возвращает объект канала по его имени. Если канала до этого не существовало, он будет создан
	GetChannel(name string) (IChannel, bool)
	// Удаляет канал по его имени.
	DeleteChannel(name string) (bool, error)
	// Run broker
	Run()
}