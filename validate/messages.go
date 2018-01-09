package validate

import "fmt"

var (
	defaultMessage = "значение не прошло проверку по правилу %s"
	messages       = map[string]string{
		"email":     "значение должно быть корректным email адресом%s",
		"eq":        "значение должно быть равно %s",
		"gt":        "значение должно быть больше %s",
		"gte":       "значение должно быть не меньше %s",
		"len":       "длина значения должна быть равна %s",
		"lt":        "значение должно быть меньше %s",
		"lte":       "значение не должно превышать %s",
		"max":       "значение не должно превышать %s",
		"min":       "значение должно быть не меньше %s",
		"ne":        "значение не должно быть равно %s",
		"required":  "обязательное поле%s",
		tagPresense: "обязательное поле%s",
	}
)

func TranslateErrorMessage(tag, param string) string {
	mes, ok := messages[tag]
	if ok {
		return fmt.Sprintf(mes, param)
	}
	return fmt.Sprintf(defaultMessage, tag)
}
