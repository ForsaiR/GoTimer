package main

import (
	"sync"
	"time"
)

var ticker *time.Ticker
var countStatus bool
var instance *counter
var once sync.Once


//Счетчик
type counter struct {
	count      *int
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
func (c *counter) getCount() int {
	return *c.count
}

//Функция для установки счетчика в значение count
func (c *counter) setCount(count int) {
	c.count = &count
}

//Функция для получения временной метки
func (c *counter) getTimeStamp() time.Time {
	return c.timeStamp
}

//Функция для установки временной метки (текущее время)
func (c *counter) setTimeStampNow() {
	c.timeStamp = time.Now()
}

//Функция для установки временной метки
func (c *counter) setTimeStamp(time time.Time) {
	c.timeStamp = time
}

//Получение данных из счетчика
func (c *counter) getDataFromCounter() (int, time.Time)  {
	return c.getCount(), c.getTimeStamp()
}

//Таймер, увеличивает значение счетчика на 1
func timer() {
	timer := time.NewTicker(1 * time.Second)
	(*getInstance()).setTimeStampNow()
	count := 0
	ticker = timer
	for range timer.C {
		count += 1
		(*getInstance()).setCount(count)
		(*getInstance()).setTimeStampNow()
	}
}

//Запуск счетчика
func startCounter() {
	if !countStatus {
		go timer()
	}
	countStatus = true
}

//Выключение счетчика
func stopCounter() {
	if ticker != nil {
		ticker.Stop()
		ticker = nil
	}

	countStatus = false
}

//Перезапуск счетчика
func resetCounter() {
	stopCounter()
	(*getInstance()).setCount(0)
	startCounter()
}

//Получение состояния счетчика
func counterStatus() bool {
	return countStatus
}