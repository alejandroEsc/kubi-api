package server

import (
    "fmt"
)

func FmtAddress(address string , port int ) string {
    return fmt.Sprintf("%s:%d", address, port)
}