package service

import (
	"sync"
	"time"
)

var instance *counter
var ticker *time.Ticker
var once sync.Once
var counterStatus bool

//Счетчик
type counter struct {
	count      int
	timeStamp time.Time
}

//Функция для получения указателя на структуру
func getInstance() *counter {
	once.Do(func() {
		instance = &counter{}
	})
	return instance
}

//Функция для получения значения счетчика
func GetCount() int {
	return (*getInstance()).count
}

//Функция для установки счетчика в значение count
func setCount(count int) {
	(*getInstance()).count = count
}

//Функция для получения временной метки
func GetTimeStamp() time.Time {
	return (*getInstance()).timeStamp
}

//Функция для установки временной метки
func setTimeStampNow() {
	(*getInstance()).timeStamp = time.Now()
}

//Функция для установки временной метки
func setTimeStamp(time time.Time) {
	(*getInstance()).timeStamp = time
}

func GetDataFromCounter() (int, time.Time)  {
	return (*getInstance()).count, (*getInstance()).timeStamp
}

//Таймер, увеличивает значение счетчика на 1
func timer() {
	timer := time.NewTicker(1 * time.Second)
	setTimeStampNow()
	count := 0
	ticker = timer
	for range timer.C {
		count += 1
		setCount(count)
		setTimeStampNow()
	}
}

//Запуск счетчика
func StartCounter() {
	if !counterStatus {
		go timer()
	}
	counterStatus = true
}

//Выключение счетчика
func StopCounter() {
	if ticker != nil {
		ticker.Stop()
		ticker = nil
	}
	counterStatus = false
}

//Перезапуск счетчика
func ResetCounter() {
	StopCounter()
	setCount(0)
	StartCounter()
}

//Получение состояния счетчика
func CounterStatus() bool {
	return counterStatus
}