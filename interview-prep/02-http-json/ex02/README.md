# HTTP Ex02: POST + Decode + Validate

Goal:

- practice JSON decoding
- validate input
- respond with `201`, `400`, and `409`

Prompt:

Complete `POST /catalog`.

Expected body:

```json
{
  "sku": "DRILL-002",
  "name": "Drill",
  "price": 80
}
```

Rules:

- `sku` and `name` are required
- `price` must be greater than 0
- if the `sku` already exists, respond with `409`
