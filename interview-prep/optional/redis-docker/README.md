# Optional: Redis with Docker

Este bloque es opcional y queda fuera del camino critico.

Objetivo:

- levantar Redis local
- practicar `GET` y `SET`
- conectar una version pequena de los ejercicios de cache

Ruta sugerida:

1. levantar Redis con Docker Compose
2. validar conexion con `redis-cli`
3. agregar una version minima de cache usando `go-redis/v9`

Si lo hacemos:

- reutilizamos algun ejercicio de `04-cache-concurrency`
- no reemplazamos la cache en memoria; solo la comparamos
