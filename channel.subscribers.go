package pubsub

import (
	"github.com/grosunick/go-common/core"
	"time"
)

// Стуктура, описывающая слушателей канала
type ChannelListeners struct {
	// Список подписчиков на канал
	listeners map[IListener]bool
	// Время последнего удаления подписчика
	lastSubscriberDeleteTime int64
}

// Создает объект списка подписчиков на канал
func NewChannelListenersList(channelProps *ChannelProps) *ChannelListeners {
	return &ChannelListeners{
		make(map[IListener]bool),
		0,
	}
}

// Добавляет подписчика
func (this *ChannelListeners) Add(listener IListener) bool {
	this.listeners[listener] = true
	return true
}

// Удаляет подписчика
func (this *ChannelListeners) Remove(listener IListener) (bool, error) {
	if _, ok := this.listeners[listener]; !ok {
		return false, &core.Error{ERROR_LISTENER_NOT_EXISTS, "listener not exists"}
	}
	delete(this.listeners, listener)
	this.lastSubscriberDeleteTime = time.Now().Unix()

	return true, nil
}

// Возвращает количество подписчиков
func (this *ChannelListeners) Len() int {
	return len(this.listeners)
}

// Возвращает признак того, что канал можно удалить по причине не использования
func (this *ChannelListeners) CanRemoveChannel(brokerProps *BrokerProps) bool {
	if brokerProps.LifeTimeWithoutSubscribers == 0 {
		return false
	}

	if len(this.listeners) > 0 {
		return false
	}

	if this.lastSubscriberDeleteTime <= 0 {
		this.lastSubscriberDeleteTime = time.Now().Unix()
	}

	if (time.Now().Unix() - this.lastSubscriberDeleteTime) > int64(brokerProps.LifeTimeWithoutSubscribers) {
		return true
	}

	return false
}
