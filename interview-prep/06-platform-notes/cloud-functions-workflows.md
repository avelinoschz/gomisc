# Cloud Functions and Workflows

## Resumen corto

- Cloud Functions encaja bien en eventos pequenos y bien delimitados.
- Workflows coordina pasos entre servicios, especialmente cuando hay orquestacion, retries o llamados a multiples APIs.

## Preguntas tipicas

### Cuando no usarias Cloud Functions

Cuando la logica ya no es pequena, cuando el flujo necesita demasiado estado o cuando el servicio encaja mejor como API larga y estable.

### Cuando usarias Workflows

Para orquestar pasos entre servicios GCP o APIs externas sin meter esa coordinacion dentro de una sola aplicacion.

### Como lo conectarias con una plataforma interna

Podrias usar Workflows para automatizar onboarding, aprobaciones o secuencias operativas; por ejemplo, validar entrada, crear recursos y notificar resultados.
