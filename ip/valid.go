package ip

import "net"

// IsValidIPv4 checks if the provided string is a valid IPv4 address.
func IsValidIPv4(ip string) bool {
	parsedIP := net.ParseIP(ip)
	if parsedIP == nil {
		return false
	}

	// Check if it is an IPv4 address
	return parsedIP.To4() != nil
}
