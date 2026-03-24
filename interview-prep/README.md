# Interview Prep

Esta carpeta contiene una ruta de preparacion intensiva para una prueba tecnica de Go con enfoque backend/platform.

Objetivos:

- recuperar velocidad escribiendo Go de memoria
- practicar HTTP, JSON, tests, cache y concurrencia
- ensayar conversacion tecnica de platform engineering

Reglas de uso:

1. Avanza en orden. Cada bloque asume el anterior.
2. `00-warmup/ex01-copy-by-hand` es el unico ejercicio completamente resuelto.
3. Del resto en adelante, el codigo es starter code o esqueletos pequenos.
4. Intenta resolver primero sin ayuda y luego trae tu solucion para review.
5. Si te atoras, lee el README del ejercicio antes de consultar documentacion externa.

Comandos utiles:

```bash
go run ./interview-prep/00-warmup/ex01-copy-by-hand
go test ./...
```

Ruta sugerida:

- Dia 1: `00-warmup` completo + `01-go-core/ex01`
- Dia 2: resto de `01-go-core` + `02-http-json/ex01`
- Dia 3: resto de `02-http-json` + `03-testing/ex01`
- Dia 4: resto de `03-testing`
- Dia 5: `04-cache-concurrency`
- Dia 6: `05-mock-interview/mock01` + `06-platform-notes`
- Dia 7: `05-mock-interview/mock02` + repaso de puntos debiles

Convenciones:

- Enunciados y notas en espanol.
- Codigo en ingles.
- En backend, intenta mantener una forma consistente:
  - `handler`
  - `service`
  - `store` o `cache`
- Usa standard library por defecto.

Checklist de progreso:

- [ ] Warmup completo
- [ ] Go core completo
- [ ] Primer handler HTTP sin consultar docs
- [ ] Primer test con `httptest`
- [ ] Cache en memoria con `sync.RWMutex`
- [ ] Ejercicio con `context.Context`
- [ ] Mock de 45-60 minutos
- [ ] Repaso de notas de platform/cloud
