# Cache Ex01: In-Memory Cache

Objetivo:

- construir una cache segura para concurrencia
- practicar `sync.RWMutex`

Enunciado:

Completa `MemoryCache`.

Necesita:

- `Get(key string) (string, bool)`
- `Set(key, value string)`

Reglas:

- usa `map[string]string`
- protege acceso con `sync.RWMutex`
- no metas TTL todavia

Criterio de terminado:

- lectura concurrente segura
- escritura segura
- codigo corto y claro
