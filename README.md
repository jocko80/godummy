# godummy

Песочница для изучения Go. Каждый пример — самодостаточный модуль с тестами и пояснениями.

## Структура

```
godummy/
├── 01-basics/ # Синтаксис, типы, управление потоком
├── 02-errors/ # Обработка ошибок, %w, errors.Is/As
├── 03-structs/ # Методы, ресиверы, embedding
├── 04-interfaces/ # Интерфейсы, полиморфизм, моки
└── 05-concurrency/ # Горутины, каналы, sync
...

```

## Как запускать примеры

Каждая папка — самостоятельный пример:

```bash
cd 02-errors/validate_hosts
go run .
go test -v -race
```
## Прогресс
- 02-errors/validate_hosts — ресиверы, %w, errors.Is, table-driven тесты
- 02-errors/validate_hosts — net/netip, Is4In6(), строгие политики
- 03-structs/embedding — встраивание типов
- 04-interfaces/polymorphism — моки для тестов
- 05-concurrency/safe-counter — sync.Mutex, race detector

## Инсайты
Короткие заметки, которые хочется запомнить:
- Пустая строка ≠ невалидный IP. Это разные типы ошибок (ErrEmptyIP vs ErrInvalidIP). Явная проверка в Parse() сохраняет семантику.
- net/netip > net для нового кода. netip.Addr — value type, без аллокаций, с нативными Is4(), Is6(), Is4In6().
- Parse, don't validate. Возвращай распарсенный объект, а не bool. Избегаешь повторного парсинга.
::ffff:192.168.1.1 — это IPv4-mapped IPv6. netip.Is4() вернёт false, Is4In6() — true. Ловушка для строгих политик.
---

### Практический workflow

**Когда изучаешь новую тему:**
```bash
cd ~/projects/godummy
mkdir -p 03-structs/embedding
cd 03-structs/embedding

# пишешь main.go, main_test.go, README.md

go run .
go test -v
git add .
git commit -m "learn: embedding types in structs"
```

**Когда хочешь поэкспериментировать:**
```bash
git checkout -b experiment/unsafe-ipv6
# ломаешь код, пробуешь идеи
git checkout main   # откат, если не получилось
```

### Раз в неделю:
- Обновляешь чекбокс в корневом README.md (прогресс)
- Дописываешь инсайт, если понял что-то важное

---

## Итог
1. Git — обязательно. Это дисциплина и безопасность экспериментов.
2. Корневой README — карта обучения с чекбоксом прогресса.
3. README в каждом примере — что изучаем, ключевые инсайты, как запускать.
4. Коммить часто с осмысленными сообщениями (learn:, exp:, fix:).
