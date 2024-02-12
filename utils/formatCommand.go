package utils

import (
	"errors"
	"github.com/slainsama/msgr_server/globals"
	"regexp"
	"strconv"
	"strings"
)

func FormatCommand(command string, taskId string, scriptName string, args []string) (formatCommand []string, err error) {
	commandFields := strings.Fields(command)
	pattern := `\{[^{}]+\}`
	re := regexp.MustCompile(pattern)
	for i, field := range commandFields {
		if re.MatchString(field) {
			switch field {
			case "{taskId}":
				commandFields[i] = taskId
			case "{server}":
				commandFields[i] = globals.UnmarshaledConfig.SERVER.Host
			case "{scriptName}":
				commandFields[i] = scriptName
			default:
				// 如果是 "{argN}" 格式的占位符
				if argIndex, err := strconv.Atoi(field[5 : len(field)-1]); err == nil && argIndex <= len(args) {
					commandFields[i] = args[argIndex-1]
				} else {
					return nil, errors.New("errFormat")
				}
			}
		}
	}
	return commandFields, nil
}
