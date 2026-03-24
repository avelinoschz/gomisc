# Mock 01: Catalog Endpoint with Cache

Duracion objetivo:

- 45 a 60 minutos

Problema:

Implementa un servicio HTTP con endpoint:

- `GET /catalog?sku=...`

Comportamiento:

- valida `sku`
- primero revisa cache
- si no existe, consulta una fuente primaria simulada
- si el producto existe, responde JSON y guarda el resultado en cache
- si no existe, responde `404`

Restricciones:

- usa solo standard library
- no uses frameworks
- manten una separacion razonable entre handler, service y cache

Checklist de evaluacion:

- codigo compila
- handler claro
- errores y status codes correctos
- cache hit/miss claro
- tests basicos presentes

Discusion de diseno:

- como cambiaria esto si corriera en Cloud Run
- que observabilidad agregarias
- donde pondrias timeouts y por que
