package reorm

import (
	"fmt"

	"github.com/go-redis/redis"
	"github.com/saromanov/re-orm/internal/storage"
)

// ReOrm provides implementation of the Redis ORM
type ReOrm struct {
	client *redis.Client
}

// New initialize Redis Orm
func New(c *Config) *ReOrm {
	client := redis.NewClient(&redis.Options{
		Addr:     c.Address,
		Password: c.Password,
		DB:       0,
	})
	return &ReOrm{
		client: client,
	}
}

// Save provides saving of the data. Also, it returns stored id
func (r *ReOrm) Save(data ...interface{}) (string, error) {
	for _, d := range data {
		if _, err := storage.Save(r.client, d); err != nil {
			return "", fmt.Errorf("unable to save data: %v", err)
		}
	}
	return "", nil
}

// GetByID provides getting of the data by id
func (r *ReOrm) GetByID(ID, data interface{}) error {
	return storage.Get(r.client, ID, data)
}

// DeleteByID provides removing of the data by id
func (r *ReOrm) DeleteByID(ID interface{}) error {
	return storage.DeleteByID(r.client, ID)
}

// Find provides finding of the data by non-default values
// at the req. For example:
// r.Find(&Car{Name: "BMW"}, &resp)
// And it should find by the field "Name"
func (r *ReOrm) Find(req, data interface{}) error {
	return storage.Find(r.client, req, data)
}
