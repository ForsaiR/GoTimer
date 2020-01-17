package service

import (
	"sync"
	"time"
)

var instance *counter
var once sync.Once
var counterGoRoutine bool

//Счетчик
type counter struct {
	count      int
	rebootFlag bool
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

//Функция перезапуска таймера
func ResetTimer() bool {
	counterGoRoutine = true
	return counterGoRoutine
}

//Функция для установки счетчика в значение count
func setCount(count int) {
	(*getInstance()).count = count
}

//Функция для изменения значения rebootFlag
func setRebootFlag(flag bool) {
	(*getInstance()).rebootFlag = flag
}

//Запуск счетчика
func StartCounter() {
	if !counterGoRoutine {
		go func() {
			for i := 0; ; i += 1 {
				if (*getInstance()).rebootFlag {
					i = 0
					setRebootFlag(false)
				}
				setCount(i)
				time.Sleep(3000000000)
			}
		}()
	}
	counterGoRoutine = true
}
