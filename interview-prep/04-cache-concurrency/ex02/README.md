# Cache Ex02: Fetch with Fallback

Objetivo:

- simular un servicio que primero consulta cache y luego una fuente lenta
- practicar flujo `cache hit` y `cache miss`

Enunciado:

Completa `CatalogService.GetProductName`.

Comportamiento esperado:

- si el valor existe en cache, lo regresa de inmediato
- si no existe, usa `ProductFetcher`
- si el fetch funciona, guarda el resultado en cache
- si el fetch falla, regresa error

Meta:

- que el flujo quede muy claro al leerlo
