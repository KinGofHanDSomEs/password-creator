# Генератор паролей

## Клонирование проекта
```
git clone git@github.com:KinGofHanDSomEs/password-creator.git
```

## Изменение конфигурации
В файле config.yaml содержится конфигурация:
+ count - количество паролей
+ width - длина паролей
+ letters - буквы
+ numbers - цифры
+ specialCharacters - специальные символы
+ isRepeatedCharacters - повторение символов (true - разрешено, false - запрещено)

## Запуск 
```
go run main.go
```
