# Terraform, OPA, Rego, CEL

## Resumen corto

- Terraform sirve para declarar infraestructura de forma repetible.
- Lo importante para entrevista no es recitar sintaxis, sino entender modulos, variables, estado y drift.
- OPA/Rego y CEL ayudan a expresar reglas o politicas para validar configuraciones y despliegues.

## Preguntas tipicas

### Que es drift

Es la diferencia entre lo declarado y lo que realmente existe en infraestructura.

### Para que sirven los modulos

Para reutilizar patrones y reducir copia/pega en IaC.

### Donde entraria policy-as-code

Antes o durante el despliegue, para bloquear configuraciones que no cumplen estandares de seguridad, naming, red o compliance.

### Como responder si te preguntan por OPA/Rego o CEL y no eres experto

"No diria que soy especialista, pero entiendo el objetivo: expresar reglas declarativas para validar configuraciones. Me siento comodo aprendiendo la sintaxis concreta mientras mantengo claro el flujo donde esas politicas se evalúan."
