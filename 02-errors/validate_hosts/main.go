// main.go
package main

import (
	"fmt"
)

func main() {
	testCases := []struct {
		ip   string
		desc string
	}{
		{"192.168.1.1", "Валидный IPv4"},
		{"2001:db8::1", "Валидный IPv6"},
		{"::ffff:192.168.1.1", "IPv4-mapped IPv6 (мусор для строгих политик)"},
		{"999.0.0.1", "Невалидный синтаксис"},
		{"", "Пустая строка"},
	}

	for _, tc := range testCases {
		h := Host{IP: tc.ip}
		fmt.Printf("\nТест: %s (%q)\n", tc.desc, tc.ip)

		// Строгий IPv4
		if _, err := h.ParseStrictIPv4(); err != nil {
			fmt.Printf("  -> Strict IPv4: ОТКЛОНЕН (%v)\n", err)
		} else {
			fmt.Printf("  -> Strict IPv4: ПРИНЯТ\n")
		}

		// Строгий IPv6
		if _, err := h.ParseStrictIPv6(); err != nil {
			fmt.Printf("  -> Strict IPv6: ОТКЛОНЕН (%v)\n", err)
		} else {
			fmt.Printf("  -> Strict IPv6: ПРИНЯТ\n")
		}
	}
}
