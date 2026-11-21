package main

import (
	"errors"
	"testing"
)

func TestSetVLAN(t *testing.T) {
	tests := []struct {
		name    string
		vlan    int
		wantErr error // nil = ожидаем успех
	}{
		{
			name:    "lower boundary (1)",
			vlan:    1,
			wantErr: nil,
		},
		{
			name:    "upper boundary (4094)",
			vlan:    4094,
			wantErr: nil,
		},
		{
			name:    "just past upper boundary (4095)",
			vlan:    4095,
			wantErr: ErrInvalidVLAN,
		},
		{
			name:    "zero",
			vlan:    0,
			wantErr: ErrInvalidVLAN,
		},
		{
			name:    "negative",
			vlan:    -1,
			wantErr: ErrInvalidVLAN,
		},
		{
			name:    "wild garbage (9999)",
			vlan:    9999,
			wantErr: ErrInvalidVLAN,
		},
		{
			name:    "typical production value (100)",
			vlan:    100,
			wantErr: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := NewSession("user-42")
			err := s.SetVLAN(tt.vlan)

			if tt.wantErr == nil {
				// Успех: ошибки нет, значение записалось
				if err != nil {
					t.Fatalf("SetVLAN(%d) unexpected error = %v; want no error", tt.vlan, err)
				}
				if got := s.VLAN(); got != tt.vlan {
					t.Errorf("SetVLAN(%d): VLAN() = %d; want %d", tt.vlan, got, tt.vlan)
				}
			} else {
				// Ошибка: тип совпадает, значение НЕ изменилось
				if !errors.Is(err, tt.wantErr) {
					t.Errorf("SetVLAN(%d) error = %v; want error type %v", tt.vlan, err, tt.wantErr)
				}
				if s.VLAN() != 0 {
					t.Errorf("SetVLAN(%d): failed set changed state: VLAN() = %d; want 0 (unchanged)", tt.vlan, s.VLAN())
				}
			}
		})
	}
}

// Проверяем инвариант: неудачный сеттер не портит уже установленное значение.
func TestSetVLANKeepsOldValueOnFailure(t *testing.T) {
	s := NewSession("user-42")

	if err := s.SetVLAN(100); err != nil {
		t.Fatalf("SetVLAN(100) unexpected error = %v", err)
	}

	if err := s.SetVLAN(9999); err == nil {
		t.Fatal("SetVLAN(9999) = nil error; want ErrInvalidVLAN")
	}

	if got := s.VLAN(); got != 100 {
		t.Errorf("after failed SetVLAN(9999): VLAN() = %d; want 100 (old value preserved)", got)
	}
}
