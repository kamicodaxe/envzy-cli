package dotenv

import (
	"strings"
)

// ParseDotenvFile parses the dotenv file content into variables and comments.
func ParseDotenvFile(content string) []Variable {
	var variables []Variable
	currentComment := ""

	lines := strings.Split(content, "\n")
	for _, line := range lines {
		line = strings.TrimSpace(line)
		if line == "" {
			// Skip empty lines
			continue
		}
		if strings.HasPrefix(line, "#") {
			// Found a comment line
			if currentComment != "" {
				// Append to the existing comment
				currentComment += "\n" + line
			} else {
				currentComment = line + "\n"
			}
		} else if strings.Contains(line, "=") {
			// Found a variable line
			parts := strings.SplitN(line, "=", 2)
			key := parts[0]
			value := parts[1]

			// Check for inline comments
			commentIndex := strings.Index(value, "#")
			if commentIndex != -1 {
				// Extract the comment and remove it from the value
				comment := strings.TrimSpace(value[commentIndex:])
				value = strings.TrimSpace(value[:commentIndex])
				if currentComment != "" {
					// Append to the existing comment
					currentComment += "\n" + comment
				} else {
					currentComment = comment + "\n"
				}
			}

			// Handle values enclosed in double quotes
			if strings.HasPrefix(value, "\"") && strings.HasSuffix(value, "\"") {
				value = strings.Trim(value, "\"")
			}

			// Create a Variable struct and append to the list
			variable := Variable{
				Key:     key,
				Value:   value,
				Comment: currentComment,
			}
			variables = append(variables, variable)

			// Reset the current comment
			currentComment = ""
		}
	}

	return variables
}
