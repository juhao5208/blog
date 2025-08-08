package config

/*import (
	"blog/common/file"
	"fmt"
	"github.com/go-playground/validator/v10"
	"gopkg.in/yaml.v2"
	"os"
	"sync"
)

var (
	validate  = validator.New()              // tag认证器
	configMap = make(map[string]interface{}) // 注册进来的配置文件，k：文件路径，v：配置文件结构体
	mu        sync.RWMutex                   // 配置文件读写锁
)

// RegisterConfig 注册配置文件
func RegisterConfig(path string, config interface{}) {
	mu.Lock()
	defer mu.Unlock()
	configMap[path] = config
}

// LoadConfig 加载配置
func LoadConfig() error {
	if len(configMap) == 0 {
		fmt.Println("not found config")
		return nil
	}

	for path, config := range configMap {
		// 文件是否存在
		if !file.FileExists(path) {
			return fmt.Errorf("config file %s not exist", path)
		}
		// 解析
		data, err := os.ReadFile(path)
		if err != nil {
			return fmt.Errorf("config file %s read error: %v", path, err)
		}
		err = yaml.Unmarshal(data, config)
		if err != nil {
			return fmt.Errorf("config file %s unmarshal error: %v", path, err)
		}
		configMap[path] = config
	}
	return nil
}
*/
