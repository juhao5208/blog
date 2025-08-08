package validate

import (
	"fmt"
	"github.com/go-playground/validator/v10"
	"sync"
)

var tagList []CustomTag
var customValidator *validator.Validate
var validatorrOnce sync.Once

type CustomTag struct {
	Tag string
	Fn  validator.Func
}

// Instance 单例对象
func Instance() *validator.Validate {
	validatorrOnce.Do(func() {
		customValidator = validator.New()
	})
	return customValidator
}

// RegisterCustomTag 注册自定义认证
func RegisterCustomTag(tag string, fn validator.Func) {
	tagList = append(tagList, CustomTag{tag, fn})
}

// InitValidator 初始化认证器
func InitValidator() error {
	for _, item := range tagList {
		err := Instance().RegisterValidation(item.Tag, item.Fn)
		if err != nil {
			return fmt.Errorf("validator register tag %s error: %v", item.Tag, err)
		}
	}
	return nil
}
