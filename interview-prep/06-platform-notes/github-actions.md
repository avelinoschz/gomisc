# GitHub Actions

## Resumen corto

Un pipeline basico para este tipo de rol deberia validar rapido, fallar pronto y separar build/test/deploy con pasos claros.

## Pipeline base que deberias poder explicar

- checkout del repo
- setup de Go
- cache de modulos
- `go test ./...`
- build
- analisis adicional opcional
- deploy condicionado por branch o entorno

## Preguntas tipicas

### Que cuidarias en CI

- feedback rapido
- jobs pequenos y legibles
- secretos bien manejados
- promover artefactos inmutables

### Como meterias AI tools sin perder calidad

Usaria Copilot o herramientas similares para acelerar implementacion, pero mantendria tests, code review y politicas de pipeline como red de seguridad.
