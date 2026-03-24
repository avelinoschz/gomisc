# Concurrency Ex03: Context Cancellation

Objetivo:

- practicar `context.Context`
- recordar como abortar trabajo cuando el caller ya no quiere esperar

Enunciado:

Completa `SlowFetcher.Fetch`.

Comportamiento esperado:

- si el contexto se cancela antes del delay, regresa `context.Canceled` o `context.DeadlineExceeded`
- si alcanza a terminar, regresa el valor

Pista:

- usa `select`
- combina `time.After` con `<-ctx.Done()`
