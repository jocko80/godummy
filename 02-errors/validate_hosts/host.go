package main

import (
	"errors"
	"fmt"
	"net/netip" // Современный, быстрый и безопасный пакет для IP (Go 1.18+)
)

// ==========================================
// 1. ГЛОБАЛЬНЫЕ ОШИБКИ (Sentinel Errors)
// ==========================================
var (
	ErrEmptyIP   = errors.New("empty IP address")
	ErrInvalidIP = errors.New("invalid IP address")
)

// ==========================================
// 2. МОДУЛЬ ВАЛИДАЦИИ IP (Сетевые политики)
// ==========================================

type Host struct {
	IP string
}

// Parse: единая точка истины для синтаксического разбора
func (h Host) Parse() (netip.Addr, error) {
	if h.IP == "" {
		return netip.Addr{}, fmt.Errorf("address %q: %w", h.IP, ErrEmptyIP)
	}

	addr, err := netip.ParseAddr(h.IP)
	if err != nil {
		return netip.Addr{}, fmt.Errorf("address %q: %w: %w", h.IP, ErrInvalidIP, err)
	}

	return addr, nil
}

// Validate: удобная обёртка, если нужен только факт валидности (DRY)
func (h Host) Validate() error {
	_, err := h.Parse()
	return err
}

// ParseStrictIPv4: принимает ТОЛЬКО чистый IPv4.
// Отвергает IPv6 и гибридные ::ffff:192.168.1.1
func (h Host) ParseStrictIPv4() (netip.Addr, error) {
	addr, err := h.Parse()
	if err != nil {
		return netip.Addr{}, err
	}

	if !addr.Is4() {
		return netip.Addr{}, fmt.Errorf("address %q: %w (требуется чистый IPv4)", h.IP, ErrInvalidIP)
	}

	return addr, nil
}

// ParseStrictIPv6: принимает ТОЛЬКО чистый IPv6.
// Отвергает IPv4 и гибридные ::ffff:192.168.1.1
func (h Host) ParseStrictIPv6() (netip.Addr, error) {
	addr, err := h.Parse()
	if err != nil {
		return netip.Addr{}, err
	}

	// addr.Is6() вернёт true и для ::ffff:192.168.1.1, поэтому добавляем проверку Is4In6()
	if !addr.Is6() || addr.Is4In6() {
		return netip.Addr{}, fmt.Errorf("address %q: %w (требуется чистый IPv6)", h.IP, ErrInvalidIP)
	}

	return addr, nil
}
