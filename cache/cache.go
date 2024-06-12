package cache

import (
	"errors"
	"sync"
	"time"

	"github.com/Ashu23042000/logger/logger"
	"github.com/Ashu23042000/proxy-server/constant"
	"github.com/Ashu23042000/proxy-server/model"
)

type ICache interface {
	InsertOne(model.Request) error
	FindOne(string) (model.Request, error)
	FindAll() ([]model.Request, error)
}

type cacheNode struct {
	data      model.Request
	next      *cacheNode
	createdAt time.Time
}

type Cache struct {
	log     logger.ILogger
	head    *cacheNode
	size    uint
	rwMutex *sync.RWMutex
}

func New(log logger.ILogger, size uint) ICache {
	return &Cache{
		log:  log,
		size: size,
	}
}

func (c *Cache) InsertOne(request model.Request) error {

	newNode := &cacheNode{data: request, createdAt: time.Now()}

	if c.head == nil {
		c.head = newNode
		return nil
	}

	// TODO: Implement logic to remove less used data from cache if size is full
	if c.size == constant.CACHE_MAX_SIZE {
		diff := time.Now().Sub(c.head.createdAt)
		c.log.Debugf("diff: %v", diff)
	}

	temp := c.head

	for temp.next != nil {
		temp = temp.next
	}

	temp.next = newNode

	return nil
}

func (c *Cache) FindOne(url string) (model.Request, error) {

	var result model.Request

	if c.head == nil {
		return model.Request{}, errors.New("cache is empty")
	}

	temp := c.head

	for temp.next != nil {
		if temp.data.Url == url {
			result = temp.data
			break
		}
		temp = temp.next
	}

	return result, nil

}

func (c *Cache) FindAll() ([]model.Request, error) {

	var result []model.Request

	if c.head == nil {
		return result, nil
	}

	temp := c.head

	for temp.next != nil {
		result = append(result, temp.data)
		temp = temp.next
	}

	return result, nil
}
