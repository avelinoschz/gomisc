# Go Core Ex02: Composition and Errors

Objetivo:

- practicar composicion simple
- usar errores para cortar flujo invalido
- preparar la mente para separar `service` y `store`

Enunciado:

Completa `DeploymentPlanner`.

Debe recibir:

- ambiente
- nombre de servicio
- replicas deseadas

Debe validar:

- ambiente permitido: `dev`, `stage`, `prod`
- servicio no vacio
- replicas positivas
- en `prod`, replicas minimas de 2

Salida:

- un `DeploymentPlan` listo para imprimirse o usarse despues
