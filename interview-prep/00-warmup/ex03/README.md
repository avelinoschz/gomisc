# Warmup Ex03: Pointer Receivers + Interface

Objetivo:

- practicar receivers con puntero
- recordar interfaces pequenas
- separar logica de estado de salida por consola

Enunciado:

Completa el `TaskList` para que:

- `Add` agregue una tarea no vacia
- `Complete` marque una tarea como completada por titulo
- `Stats` devuelva total y completadas
- `ConsolePrinter` imprima el resumen usando la interfaz `StatsReporter`

Reglas:

- si el titulo viene vacio, regresa error
- si `Complete` no encuentra la tarea, regresa error
- evita sobreingenieria

Comando:

```bash
go run ./interview-prep/00-warmup/ex03
```
