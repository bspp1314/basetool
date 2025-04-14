package basefile

import (
	"os"
	"strings"
)

func StringListToFileLine(datas []string, file string) error {
	content := strings.Join(datas, "\n")
	return os.WriteFile(file, []byte(content), 0644)
}

func ReadFileToStringList(file string) ([]string, error) {
	// 读取文件内容
	content, err := os.ReadFile(file)
	if err != nil {
		return nil, err
	}

	// 将内容按行分割成字符串列表
	lines := strings.Split(string(content), "\n")

	// 移除空行
	var result []string
	for _, line := range lines {
		result = append(result, line)

	}

	return result, nil
}

func ReadFileToString(file string) (string, error) {
	content, err := os.ReadFile(file)
	if err != nil {
		return "", err
	}

	return string(content), nil
}

func ReadFileToMap(file string) (map[string]bool, error) {
	content, err := os.ReadFile(file)
	if err != nil {
		return nil, err
	}

	// 将内容按行分割成字符串列表
	lines := strings.Split(string(content), "\n")

	// 移除空行
	var result map[string]bool
	for _, line := range lines {
		result[line] = true
	}

	return result, nil
}
