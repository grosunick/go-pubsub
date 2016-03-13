package pubsub

// Константы ошибок
const ERROR_LISTENER_NOT_EXISTS = 21 // подписчик не существует
const ERROR_CHANNEL_NOT_EXISTS = 22  // канал не существует

// Период времени по прошествии которого происходит проверка и удаление старых каналов без подписчиков
const CHANNELS_WITHOUT_SUBSCRIBERS_CHECK_PERIOD = 600
