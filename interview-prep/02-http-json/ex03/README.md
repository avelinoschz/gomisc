# HTTP Ex03: Consistent Error Handling

Objetivo:

- evitar repetir demasiado codigo en handlers
- crear helpers pequenos para respuestas JSON

Enunciado:

Completa:

- `writeJSON`
- `writeError`
- el handler `GET /health`

Respuesta de `GET /health`:

```json
{
  "status": "ok",
  "service": "catalog-api"
}
```

Cuando ocurra un error de serializacion:

- responder `500`
- devolver un JSON de error consistente
