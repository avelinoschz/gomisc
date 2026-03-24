# Observability and Troubleshooting

## Resumen corto

En un rol de platform/backend importa mucho poder aislar si el problema es de codigo, configuracion, red, dependencia externa o despliegue.

## Cosas que deberias mencionar con naturalidad

- logs estructurados
- metricas basicas de latencia, error rate y throughput
- request IDs o trace IDs
- health checks
- timeouts y retries con criterio

## Preguntas tipicas

### Como investigarias un servicio lento en produccion

Empezaria revisando metricas y logs para acotar si es general o puntual, luego veria dependencias externas, latencia por endpoint y cambios recientes de despliegue o configuracion.

### Que agregarias a un servicio nuevo

- logs estructurados
- metricas de requests
- health endpoint
- timeouts en clientes externos
- dashboards y alertas basicas
