# Go Core Ex01: Parse and Validate

Objetivo:

- practicar parsing de texto a structs
- dividir validaciones en funciones pequenas
- escribir errores utiles

Enunciado:

Implementa `ParseServices`.

Formato de entrada:

- una linea por servicio
- cada linea viene como `name,team,replicas`

Ejemplo:

```text
catalog,platform,3
checkout,commerce,2
```

Reglas:

- `name` y `team` son obligatorios
- `replicas` debe ser entero positivo
- si una linea es invalida, regresa error y detente
- ignora lineas vacias

Criterio de terminado:

- parsing claro
- errores con contexto suficiente
- sin paquetes externos
