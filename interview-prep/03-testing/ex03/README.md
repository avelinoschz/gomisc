# Testing Ex03: Error Cases and Invalid JSON

Goal:

- practice negative cases
- get faster at reviewing decode errors and status codes

Prompt:

Complete `createProductHandler` and then write tests.

Minimum cases:

- `400` for invalid JSON
- `400` for an invalid payload
- `201` for a valid product
- `409` if the SKU already exists
