# HTTP Ex01: GET + Query Params + JSON

Objetivo:

- implementar un handler `GET`
- leer query params
- responder JSON

Enunciado:

Completa el endpoint `GET /catalog?sku=...`.

Comportamiento esperado:

- si falta `sku`, responder `400`
- si no existe, responder `404`
- si existe, responder `200` con JSON del producto

Notas:

- usa el catalogo en memoria ya incluido
- no necesitas base de datos
- manten la respuesta simple y clara
