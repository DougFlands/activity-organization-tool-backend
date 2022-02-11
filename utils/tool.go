package utils

import (
	"encoding/json"
	"fmt"
)

func ToolJsonFmt(j interface{}) {
	s, _ := json.MarshalIndent(j, "", " ")
	fmt.Printf("%s\n", s)
}
