package address

import "strings"

// ValidTypes contains all address types recognized by the app
var ValidTypes []string = []string{"rua", "avenida", "estrada", "rodovia"}

// GetAddressType returns the given address type
func GetAddressType(addr string) string {
	addr = strings.TrimSpace(addr)
	addr = strings.ToLower(addr)

	for _, addrType := range ValidTypes {
		if strings.Contains(addr, addrType) {
			return addrType
		}
	}

	return "tipo invalido"
}
