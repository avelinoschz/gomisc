# Go Core Ex03: In-Memory Service

Objetivo:

- introducir una interfaz pequena con sentido
- separar `service` y `store`
- practicar errores de negocio

Enunciado:

Implementa un mini flujo de onboarding de aplicaciones.

Necesitas:

- `AppStore` con `Save` y `GetByName`
- `InMemoryAppStore`
- `AppService.Register`

Reglas:

- `Name` y `OwnerTeam` son obligatorios
- no se permiten nombres duplicados
- `Register` debe delegar persistencia al store

Meta:

- que el codigo te recuerde a una version pequena de backend real
