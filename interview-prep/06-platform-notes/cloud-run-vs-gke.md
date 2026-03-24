# Cloud Run vs GKE

## Resumen corto

- Cloud Run sirve muy bien para servicios HTTP stateless con despliegue rapido y poca friccion operativa.
- GKE tiene mas poder y mas complejidad; lo eliges cuando necesitas control fino de networking, sidecars, scheduling o cargas mas especializadas.
- En una plataforma interna, Cloud Run suele acelerar onboarding; GKE suele aparecer cuando ya hay necesidades mas avanzadas o estandarizacion de Kubernetes.

## Preguntas tipicas

### Cuando elegirias Cloud Run sobre GKE

Cuando el servicio es stateless, tiene trafico HTTP claro, no necesita control profundo del cluster y queremos minimizar operacion.

### Cuando elegirias GKE

Cuando necesitamos Kubernetes nativo, sidecars, DaemonSets, politicas complejas, workloads de larga vida o configuraciones avanzadas de red.

### Riesgos o tradeoffs de Cloud Run

- menos control fino del runtime
- cold starts segun el patron de trafico
- limites y modelo mas opinionado

### Que dirias en entrevista

"Para una plataforma de migracion, Cloud Run ayuda a mover equipos rapido porque reduce la carga operativa. Si la aplicacion necesita capacidades mas avanzadas de Kubernetes o mas control de infraestructura, entonces GKE se vuelve mejor opcion."
