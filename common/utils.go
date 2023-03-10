package common

func CheckResultStatus(result *Result) bool {
	if result == nil || result.UID == "" ||
		result.UID == InvalidUID ||
		result.Status != Success ||
		result.Label == InvalidLabel {
		return false
	}
	return true
}
