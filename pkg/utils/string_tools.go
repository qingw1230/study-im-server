package utils

// IsContain 检查 list 中是否包含 target
func IsContain(target string, list []string) bool {
	for _, l := range list {
		if target == l {
			return true
		}
	}
	return false
}
