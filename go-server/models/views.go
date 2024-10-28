package models

import "fmt"

func GenerateViewPath(path string) string {
	return fmt.Sprintf("./views/%s.json", path)
}
