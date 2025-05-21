package utils

import (
	"encoding/gob"
	"fmt"
	"os"
	"path/filepath"
)

// Cache 结构体定义
type Cache struct {
	BasePath string
}

// NewCache 创建一个新的缓存实例，并确保目标目录存在
func NewCache(basePath string) (*Cache, error) {
	// 确保路径存在
	if err := os.MkdirAll(basePath, 0755); err != nil {
		return nil, fmt.Errorf("无法创建缓存目录: %v", err)
	}
	return &Cache{
		BasePath: basePath,
	}, nil
}

// Set 写入键值对到文件中
func (c *Cache) Set(key string, data interface{}) error {
	filePath := filepath.Join(c.BasePath, key)
	file, err := os.Create(filePath)
	if err != nil {
		return fmt.Errorf("无法创建文件 %s: %v", key, err)
	}
	defer file.Close()

	encoder := gob.NewEncoder(file)
	if err := encoder.Encode(data); err != nil {
		return fmt.Errorf("编码数据失败: %v", err)
	}
	return nil
}

// Get 根据 key 读取数据
func (c *Cache) Get(key string, data interface{}) error {
	filePath := filepath.Join(c.BasePath, key)
	file, err := os.Open(filePath)
	if err != nil {
		return fmt.Errorf("无法打开文件 %s: %v", key, err)
	}
	defer file.Close()

	decoder := gob.NewDecoder(file)
	if err := decoder.Decode(data); err != nil {
		return fmt.Errorf("解码数据失败: %v", err)
	}
	return nil
}

// Clean 删除指定 key 的缓存文件
func (c *Cache) Clean(key string) error {
	filePath := filepath.Join(c.BasePath, key)
	if err := os.Remove(filePath); err != nil {
		if os.IsNotExist(err) {
			return fmt.Errorf("文件不存在: %s", key)
		}
		return fmt.Errorf("删除文件失败 %s: %v", key, err)
	}
	return nil
}

// CleanAll 清除所有缓存文件
func (c *Cache) CleanAll() error {
	files, err := os.ReadDir(c.BasePath)
	if err != nil {
		return fmt.Errorf("读取缓存目录失败: %v", err)
	}

	for _, file := range files {
		filePath := filepath.Join(c.BasePath, file.Name())
		if err := os.Remove(filePath); err != nil {
			return fmt.Errorf("删除文件失败 %s: %v", filePath, err)
		}
	}
	return nil
}
