# Validate Hosts

## Что изучаем

- Обработка ошибок через `%w` и `errors.Is`
- Ресиверы по значению vs указателю (Боднер, гл. 7)
- Table-driven тесты с `t.Run`
- `net/netip` вместо устаревшего `net.IP`
- Строгие политики валидации (IPv4 vs IPv6 vs mapped)

## Ключевые инсайты

### 1. Ресивер по значению
```go
func (h Host) Parse() (netip.Addr, error) { ... }
```
Метод не мутирует `Host`, значит — value receiver. Копирование дешёвое (одна строка).

### 2. Единая точка истины
`Validate()` вызывает `Parse()` под капотом. Логика валидации — в одном месте.

### 3. Ловушка IPv4-mapped
`::ffff:192.168.1.1` — валидный IP, но не чистый IPv4.
- `addr.Is4() → false`
- `addr.Is4In6() → true`
Для строгих политик (BNG, Accel-PPP) нужно явно отвергать такие адреса.

## Запуск
```bash
go run .
go test -v -race
go test -cover
```
## Связанные темы
- 03-structs/embedding — как встроить Host в другие структуры
- 04-interfaces/ — как сделать интерфейс Validator для Host
