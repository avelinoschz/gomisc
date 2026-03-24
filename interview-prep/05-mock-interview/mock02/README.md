# Mock 02: Add a New Requirement

Duracion objetivo:

- 45 a 60 minutos

Problema:

Toma el mock anterior y agrega una regla nueva:

- soportar `GET /catalog?sku=...&include=meta`

Cuando `include=meta`, la respuesta debe incluir:

- `source`: `cache` o `store`
- `served_at`: timestamp en formato RFC3339

Restricciones:

- no rompas el formato base del producto
- intenta extender la solucion sin reescribir todo

Checklist de evaluacion:

- cambio pequeno pero limpio
- no duplicar demasiada logica
- tests adaptados a la nueva respuesta

Discusion de diseno:

- como harrias versionado de API si esto creciera
- como evitar acoplar demasiado el handler al formato de respuesta
