# Warmup Ex02: Slice + Map Summary

Objetivo:

- practicar transformacion de datos
- recorrer slices
- acumular resultados en un map
- devolver multiples valores

Enunciado:

Implementa `BuildDepartmentSummary`.

Entrada:

- una lista de empleados

Salida:

- un `map[string]int` con total de empleados por departamento
- un entero con el total de empleados activos

Reglas:

- ignora empleados con `Department == ""`
- solo cuenta en `activeTotal` a los que tengan `Active == true`
- el `map` debe incluir todos los departamentos validos aunque solo tengan un empleado

Criterio de terminado:

- compila
- usa un solo loop principal
- no usa paquetes externos

Comando:

```bash
go run ./interview-prep/00-warmup/ex02
```
