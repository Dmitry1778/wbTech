package db

import (
	"context"
	"log"
	"sync"
	"wbTech/internal/domain"
)

type Cache struct {
	mutex   *sync.RWMutex
	data    map[string]*domain.NewOrder
	DBItems *DB
}

func NewInit(db *DB) *Cache {
	csh := Cache{}
	csh.InitCache(db)
	return &csh
}

func (c *Cache) InitCache(db *DB) {
	c.DBItems = db
	db.SetCacheInstance(c)
	c.mutex = &sync.RWMutex{}

	c.data = make(map[string]*domain.NewOrder)

	// Восстанавление кеша из базы данных, если он есть в бд
	c.DBItems.getCacheFromDatabase()
}

func (c *Cache) Set(key string, value *domain.NewOrder) {
	c.mutex.Lock()
	c.data[key] = value
	c.mutex.Unlock()

	err := c.DBItems.PutOrder(context.Background(), domain.NewOrder{})
	if err != nil {
		panic(err.Error())
	}
}

func (c *Cache) Get(key string) (*domain.NewOrder, error) {
	var order *domain.NewOrder
	var err error

	c.mutex.RLock()
	order, found := c.data[key]
	c.mutex.RUnlock()

	if found {
		log.Printf("Order found:%v\n", found)
	} else {
		order, err = c.DBItems.GetOrder(context.Background(), key)
		if err != nil {
			log.Printf("%s: GetOrder(): ошибка получения Order:", err)
			return order, err
		}
		c.Set(key, order)
	}
	return order, nil
}
