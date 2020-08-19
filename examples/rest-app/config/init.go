package config

import (
	"fmt"
	"time"
)

func init() {
	fmt.Println("App Go started", time.Now().Format("02-01-2006"))
}
