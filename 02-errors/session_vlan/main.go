package main

import (
	"errors"
	"fmt"
)

func main() {
	testUser("user-valid", 100)
	testUser("user-bad", 9999)
}

func testUser(userID string, vlan int) {
	// Вызываем публичную функцию
	sess, err := HandleAuth(userID, vlan)

	if err != nil {
		if errors.Is(err, ErrInvalidVLAN) {
			fmt.Printf("WARN: Пользователь %s — недопустимый VLAN: %v. Требуется проверка.\n", userID, err)
			return
		}
		fmt.Printf("ERROR: Пользователь %s: %v\n", userID, err)
		return
	}

	fmt.Printf("OK: Сессия создана, VLAN = %d\n", sess.VLAN())
}
