package modifiers

import "strings"

func lowerCase(content string) string {
    return strings.ToLower(content)
}

func titleCase(content string) string {
    return strings.Title(content)
}

func ApplyModifier(variable, value string) string {
    switch modifier := GetModifier(variable); modifier {
        case "lowerCase":
            return lowerCase(value)
        case "titleCase":
            return titleCase(value)
        default:
            return value
    }
}

func VariableHasModifier(variable string) bool {
    return strings.Contains(variable, " ")
}

func GetKey(variable string) string {
    return strings.Split(variable, " ")[0]
}

func GetModifier(variable string) string {
    return strings.Split(variable, " ")[1]
}
