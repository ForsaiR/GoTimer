package main

import (
	"fmt"
	"sync"
	"time"
)

var ticker *time.Ticker
var countStatus bool
var instance *Counter
var once sync.Once


//Счетчик
type Counter struct {
	Count      int
	TimeStamp time.Time
}

//Функция для получения указателя на структуру
func getInstance() *Counter {
	once.Do(func() {
		instance = &Counter{}
	})
	return instance
}

//Функция для получения значения счетчика
func (c *Counter) getCount() int {
	return c.Count
}

//Функция для установки счетчика в значение count
func (c *Counter) setCount(count int) {
	c.Count = count
}

//Функция для получения временной метки
func (c *Counter) getTimeStamp() time.Time {
	return c.TimeStamp
}

//Функция для установки временной метки (текущее время)
func (c *Counter) setTimeStampNow() {
	c.TimeStamp = time.Now()
}

//Функция для установки временной метки
func (c *Counter) setTimeStamp(time time.Time) {
	c.TimeStamp = time
}

//Получение данных из счетчика
func (c *Counter) getDataFromCounter() (int, time.Time)  {
	return c.getCount(), c.getTimeStamp()
}

//Таймер, увеличивает значение счетчика на 1 и отправляет значение в хаб
func (h *Hub) timer() {
	timer := time.NewTicker(1 * time.Second)
	(*getInstance()).setTimeStampNow()
	count := 0
	ticker = timer
	fmt.Printf("count: %d\n", count)

	h.broadcast <- count
	for range timer.C {
		count += 1
		(*getInstance()).setCount(count)
		(*getInstance()).setTimeStampNow()
		h.broadcast <- count
		fmt.Printf("count: %d\n", (*getInstance()).Count)
	}
}

//Запуск счетчика c передачей сообщения в сокет
func (h *Hub) startCounter() {
	if !countStatus {
		go h.timer()
	}
	countStatus = true
}

//Выключение счетчика
func (h *Hub) stopCounter() {
	if ticker != nil {
		ticker.Stop()
		ticker = nil
	}

	countStatus = false
}

//Перезапуск счетчика
func (h *Hub) resetCounter() {
	h.stopCounter()
	(*getInstance()).setCount(0)
	h.startCounter()
}

//Получение состояния счетчика
func counterStatus() bool {
	return countStatus
}