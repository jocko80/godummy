package main

import (
	"errors"
	"testing"
)

func TestParseStrictIPv4(t *testing.T) {
	tests := []struct {
		name    string
		ip      string
		wantIP  string
		wantErr error
	}{
		{
			name:    "valid IPv4",
			ip:      "192.168.1.1",
			wantIP:  "192.168.1.1",
			wantErr: nil,
		},
		{
			name:    "valid IPv4 another",
			ip:      "10.0.0.5",
			wantIP:  "10.0.0.5",
			wantErr: nil,
		},
		{
			name:    "invalid IPv4 (out of range)",
			ip:      "999.10.0.1",
			wantIP:  "",
			wantErr: ErrInvalidIP,
		},
		{
			name:    "empty IP",
			ip:      "",
			wantIP:  "",
			wantErr: ErrEmptyIP,
		},
		{
			name:    "IPv6 address (should be rejected by strict policy)",
			ip:      "2001:db8::68",
			wantIP:  "",
			wantErr: ErrInvalidIP,
		},
		{
			name:    "IPv4-mapped IPv6 (::ffff:...) — must be rejected",
			ip:      "::ffff:192.168.1.1",
			wantIP:  "",
			wantErr: ErrInvalidIP,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := Host{IP: tt.ip}
			ip, err := h.ParseStrictIPv4()

			if tt.wantErr == nil {
				// Успех: ошибки нет, вернулся именно тот IP, который ждали
				if err != nil {
					t.Errorf("ParseStrictIPv4(%q) unexpected error = %v; want no error", tt.ip, err)
				}
				if ip.String() != tt.wantIP {
					t.Errorf("ParseStrictIPv4(%q) IP = %v; want %v", tt.ip, ip.String(), tt.wantIP)
				}
			} else {
				// Ошибка: тип совпадает, адрес нулевой
				if !errors.Is(err, tt.wantErr) {
					t.Errorf("ParseStrictIPv4(%q) error = %v; want error type %v", tt.ip, err, tt.wantErr)
				}
				if ip.IsValid() {
					t.Errorf("ParseStrictIPv4(%q) при ошибке вернул валидный IP %v; want нулевой netip.Addr{}", tt.ip, ip)
				}
			}
		})
	}
}

func TestParseStrictIPv6(t *testing.T) {
	tests := []struct {
		name    string
		ip      string
		wantIP  string
		wantErr error
	}{
		{
			name:    "valid IPv6",
			ip:      "2001:db8::1",
			wantIP:  "2001:db8::1",
			wantErr: nil,
		},
		{
			name:    "valid IPv6 expanded (parser normalizes)",
			ip:      "2001:0db8:0000:0000:0000:0000:0000:0001",
			wantIP:  "2001:db8::1",
			wantErr: nil,
		},
		{
			name:    "IPv4 rejected",
			ip:      "192.168.1.1",
			wantIP:  "",
			wantErr: ErrInvalidIP,
		},
		{
			name:    "IPv4-mapped IPv6 — must be rejected",
			ip:      "::ffff:192.168.1.1",
			wantIP:  "",
			wantErr: ErrInvalidIP,
		},
		{
			name:    "invalid syntax",
			ip:      "999.0.0.1",
			wantIP:  "",
			wantErr: ErrInvalidIP,
		},
		{
			name:    "empty IP",
			ip:      "",
			wantIP:  "",
			wantErr: ErrEmptyIP,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := Host{IP: tt.ip}
			ip, err := h.ParseStrictIPv6()

			if tt.wantErr == nil {
				// Успех: ошибки нет, адрес нормализован до канонической формы
				if err != nil {
					t.Errorf("ParseStrictIPv6(%q) unexpected error = %v; want no error", tt.ip, err)
				}
				if ip.String() != tt.wantIP {
					t.Errorf("ParseStrictIPv6(%q) IP = %v; want %v", tt.ip, ip.String(), tt.wantIP)
				}
			} else {
				// Ошибка: тип совпадает, адрес нулевой
				if !errors.Is(err, tt.wantErr) {
					t.Errorf("ParseStrictIPv6(%q) error = %v; want error type %v", tt.ip, err, tt.wantErr)
				}
				if ip.IsValid() {
					t.Errorf("ParseStrictIPv6(%q) при ошибке вернул валидный IP; want нулевой netip.Addr{}", tt.ip)
				}
			}
		})
	}
}

func TestValidate(t *testing.T) {
	if err := (Host{IP: "192.168.1.1"}).Validate(); err != nil {
		t.Errorf("Validate(valid) = %v; want nil", err)
	}
	if err := (Host{IP: "not-an-ip"}).Validate(); !errors.Is(err, ErrInvalidIP) {
		t.Errorf("Validate(invalid) = %v; want ErrInvalidIP", err)
	}
	if err := (Host{IP: ""}).Validate(); !errors.Is(err, ErrEmptyIP) {
		t.Errorf("Validate(empty) = %v; want ErrEmptyIP", err)
	}
}
