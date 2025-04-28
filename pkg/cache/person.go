package cache

import "context"

func (c *AppCache) SavePerson(ctx context.Context, name string, age uint64) error {
	err := c.Set(ctx, name, age, 0).Err()
	if err != nil {
		return err
	}
	return nil
}

func (c *AppCache) GetPerson(ctx context.Context, name string) (uint64, error) {
	val, err := c.Get(ctx, name).Uint64()
	if err != nil {
		return 0, err
	}
	return val, nil
}
