# GO CONCURRENCY PATTERNS

## Overall
Здесь представленны паттерны синхронизаций каналов в Golang

## Navigation

1. [Sync](1-sync/README.md)
2. [Generator](2-generator/README.md)
3. [Fan-in](3-fan-in/README.md)
4. [Restoring sequence](4-restoring-sequencing/README.md)
5. [Selector](5-selector/README.md)
6. [Timeout](6-timeout/README.md)
7. [Quit channel](7-quit-channel/README.md)
8. [Replication search](8-replication-search/README.md)
9. [Worker pool](9-worker-pool/README.md)
10. [Context](10-context-use/README.md)
11. [Publish/Subscribe](11-publish-subscribe/README.md)
12. [Pipeline](12-pipeline/README.md)

## Recommendation
Использовать каналы нужно когда:
- Множество входящих данных
- Множество выходящих данных
- Необходимость в таймауте
- Чувствительность к ошибкам

## Bonus
Пример микросервисного приложения использующего механизм синхронизации для отправки http запросов.