package plugin

import "github.com/hashicorp/go-uuid"

func VariableParser(variable string) string{
	switch variable {
	case "{{uuid}}":
		id, _ := uuid.GenerateUUID()
		return id
	}

	return variable
}
