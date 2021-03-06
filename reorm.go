package reorm

import (
	"fmt"

	"github.com/go-redis/redis"
	"github.com/saromanov/re-orm/internal/config"
	"github.com/saromanov/re-orm/internal/storage"
)

// ReOrm provides implementation of the Redis ORM
type ReOrm struct {
	client *redis.Client
	config *config.Config
}

// New initialize Redis Orm
func New(c *Config) (*ReOrm, error) {
	if c == nil {
		return nil, fmt.Errorf("config is not defined")
	}
	if c.Address == "" {
		return nil, fmt.Errorf("address to redis is not defined")
	}
	client := redis.NewClient(&redis.Options{
		Addr:     c.Address,
		Password: c.Password,
		DB:       0,
	})
	return &ReOrm{
		client: client,
		config: &config.Config{
			KeyPrefix: c.KeyPrefix,
			SetType:   config.SetType(c.SetType),
		},
	}, nil
}

// Save provides saving of the data. Also, it returns stored id
func (r *ReOrm) Save(data ...interface{}) error {
	for _, d := range data {
		if _, err := storage.Save(r.client, d); err != nil {
			return fmt.Errorf("unable to save data: %v", err)
		}
	}
	return nil
}

// Get provides getting of the data by search request
func (r *ReOrm) Get(resp, data interface{}) error {
	return storage.Get(r.client, resp, data)
}

// GetValueByField provides getting of the data by the field
// from the struct
func (r *ReOrm) GetValueByField(name, field string, req interface{}) (interface{}, error) {
	return storage.GetValueByField(r.client, name, field, req)
}

// Last provides getting of the last element if data is duplicated
func (r *ReOrm) Last(resp, data interface{}) error {
	return storage.Last(r.client, resp, data)
}

// GetByID provides getting of the data by id
func (r *ReOrm) GetByID(name string, ID, data interface{}) error {
	return storage.GetByID(r.client, name, ID, data)
}

// DeleteByID provides removing of the data by id
func (r *ReOrm) DeleteByID(req interface{}, ID interface{}) error {
	return storage.DeleteByID(r.client, req, ID)
}

// Find provides finding of the data by non-default values
// at the req. For example:
// r.Find(&Car{Name: "BMW"})
// And it should find by the field "Name"
func (r *ReOrm) Find(req interface{}) ([]interface{}, error) {
	return storage.Find(r.client, req)
}

// Update provides updating of the data by request
func (r *ReOrm) Update(id, req interface{}) error {
	return storage.Update(r.client, id, req)
}
