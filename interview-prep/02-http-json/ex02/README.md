# HTTP Ex02: POST + Decode + Validate

Objetivo:

- practicar decode de JSON
- validar input
- responder con `201`, `400` y `409`

Enunciado:

Completa `POST /catalog`.

Body esperado:

```json
{
  "sku": "DRILL-002",
  "name": "Drill",
  "price": 80
}
```

Reglas:

- `sku` y `name` son obligatorios
- `price` debe ser mayor que 0
- si el `sku` ya existe, responder `409`
