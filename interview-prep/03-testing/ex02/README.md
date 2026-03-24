# Testing Ex02: Handler Tests with httptest

Objetivo:

- probar un handler sin levantar un servidor real
- validar status code y body JSON

Enunciado:

Escribe tests para `catalogHandler`.

Casos minimos:

- `400` cuando falta `sku`
- `404` cuando no existe
- `200` con JSON valido cuando existe

Tip:

- usa `httptest.NewRequest`
- usa `httptest.NewRecorder`
