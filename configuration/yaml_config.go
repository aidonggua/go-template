package configuration

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"gopkg.in/yaml.v3"
	"os"
)

var Configs map[string]any

func Init(profile Profile, workDir string) {
	yamlConfig, err := os.ReadFile(fmt.Sprintf("%s/application.yaml", workDir))
	if err != nil {
		panic(fmt.Errorf("读取配置文件错误"))
	}

	configs := make(map[string]any)
	err = yaml.Unmarshal(yamlConfig, &configs)
	if err != nil {
		panic(fmt.Errorf("解析配置文件错误"))
	}

	if profile == ProfileDev || profile == ProfileProd || profile == ProfileTest {
		yamlConfig, err := os.ReadFile(fmt.Sprintf("%s/application-%s.yaml", workDir, profile))
		err = yaml.Unmarshal(yamlConfig, &configs)
		if err != nil {
			panic(fmt.Errorf("解析配置文件错误"))
		}
	} else if profile == ProfileDefault {
	} else {
		panic(fmt.Errorf("无效的profile: %s", profile))
	}

	Configs = flattenConfig(configs)
	log.Debug(Configs)
}

func flattenConfig(configs map[string]any) map[string]any {
	flattenConfigs := make(map[string]any)
	for key, value := range configs {
		switch value.(type) {
		case map[string]any:
			flattenConfig := flattenConfig(value.(map[string]any))
			for k, v := range flattenConfig {
				flattenConfigs[fmt.Sprintf("%s.%s", key, k)] = v
			}
		default:
			flattenConfigs[key] = value
		}
	}
	return flattenConfigs
}
