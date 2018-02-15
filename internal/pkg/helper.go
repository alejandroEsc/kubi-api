package internalpkg

import (
	"fmt"
)

// FmtAddress allows you formatting addresses with address and port
func FmtAddress(address string, port int) string {
	return fmt.Sprintf("%s:%d", address, port)
}
