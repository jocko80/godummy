package main

import (
	"errors"
	"fmt"
)

// ==========================================
// 1. ГЛОБАЛЬНЫЕ ОШИБКИ (Sentinel Errors)
// ==========================================
var (
	ErrInvalidVLAN = errors.New("invalid VLAN")
)

// ==========================================
// 2. МОДУЛЬ СЕССИЙ И VLAN (Биллинг / Accel-PPP)
// ==========================================

type Session struct {
	ID   string
	vlan int // Приватное поле: защита от случайного изменения снаружи
}

func NewSession(id string) *Session {
	return &Session{ID: id}
}

// SetVLAN гарантирует, что в системе не появится невалидный VLAN
func (s *Session) SetVLAN(vlan int) error {
	if vlan < 1 || vlan > 4094 {
		return fmt.Errorf("vlan=%d: %w", vlan, ErrInvalidVLAN)
	}
	s.vlan = vlan
	return nil
}

// VLAN геттер для безопасного чтения
func (s *Session) VLAN() int {
	return s.vlan
}

// handleAuth пример контроллера с мягкой обработкой ошибок биллинга
func HandleAuth(userID string, rawVlan int) (*Session, error) {
	sess := NewSession(userID)

	if err := sess.SetVLAN(rawVlan); err != nil {
		return nil, err
	}

	return sess, nil
}
