package pubsub

// Настройки канала
type ChannelProps struct {
	// Максимальный размер очереди не отправленных сообщений в канале
	MaxQueueSize uint8
}
