package cache

import (
	"context"
	"encoding/json"
	"time"
)

// CacheService Redis缓存服务接口
type CacheService interface {
	// 基础缓存操作
	Set(ctx context.Context, key string, value interface{}, expiration time.Duration) error
	Get(ctx context.Context, key string, dest interface{}) error
	Delete(ctx context.Context, keys ...string) error
	Exists(ctx context.Context, key string) (bool, error)
	SetExpire(ctx context.Context, key string, expiration time.Duration) error

	// Hash操作 (适用于用户会话、配置等)
	HSet(ctx context.Context, key, field string, value interface{}) error
	HGet(ctx context.Context, key, field string, dest interface{}) error
	HDelete(ctx context.Context, key string, fields ...string) error
	HExists(ctx context.Context, key, field string) (bool, error)

	// Set操作 (适用于权限集合等)
	SAdd(ctx context.Context, key string, members ...interface{}) error
	SRemove(ctx context.Context, key string, members ...interface{}) error
	SMembers(ctx context.Context, key string) ([]string, error)
}

// cacheService Redis缓存服务实现
type cacheService struct {
	client Client
}

// NewCacheService 创建新的缓存服务实例
func NewCacheService(client Client) CacheService {
	return &cacheService{
		client: client,
	}
}

// Set 设置缓存值
func (s *cacheService) Set(ctx context.Context, key string, value interface{}, expiration time.Duration) error {
	data, err := json.Marshal(value)
	if err != nil {
		return err
	}
	return s.client.Set(ctx, key, data, expiration).Err()
}

// Get 获取缓存值
func (s *cacheService) Get(ctx context.Context, key string, dest interface{}) error {
	data, err := s.client.Get(ctx, key).Result()
	if err != nil {
		return err
	}
	return json.Unmarshal([]byte(data), dest)
}

// Delete 删除缓存键
func (s *cacheService) Delete(ctx context.Context, keys ...string) error {
	return s.client.Del(ctx, keys...).Err()
}

// Exists 检查键是否存在
func (s *cacheService) Exists(ctx context.Context, key string) (bool, error) {
	count, err := s.client.Exists(ctx, key).Result()
	return count > 0, err
}

// SetExpire 设置键过期时间
func (s *cacheService) SetExpire(ctx context.Context, key string, expiration time.Duration) error {
	return s.client.Expire(ctx, key, expiration).Err()
}

// HSet 设置哈希字段值
func (s *cacheService) HSet(ctx context.Context, key, field string, value interface{}) error {
	data, err := json.Marshal(value)
	if err != nil {
		return err
	}
	return s.client.HSet(ctx, key, field, data).Err()
}

// HGet 获取哈希字段值
func (s *cacheService) HGet(ctx context.Context, key, field string, dest interface{}) error {
	data, err := s.client.HGet(ctx, key, field).Result()
	if err != nil {
		return err
	}
	return json.Unmarshal([]byte(data), dest)
}

// HDelete 删除哈希字段
func (s *cacheService) HDelete(ctx context.Context, key string, fields ...string) error {
	return s.client.HDel(ctx, key, fields...).Err()
}

// HExists 检查哈希字段是否存在
func (s *cacheService) HExists(ctx context.Context, key, field string) (bool, error) {
	return s.client.HExists(ctx, key, field).Result()
}

// SAdd 添加集合成员
func (s *cacheService) SAdd(ctx context.Context, key string, members ...interface{}) error {
	return s.client.SAdd(ctx, key, members...).Err()
}

// SRemove 删除集合成员
func (s *cacheService) SRemove(ctx context.Context, key string, members ...interface{}) error {
	return s.client.SRem(ctx, key, members...).Err()
}

// SMembers 获取集合所有成员
func (s *cacheService) SMembers(ctx context.Context, key string) ([]string, error) {
	return s.client.SMembers(ctx, key).Result()
}
