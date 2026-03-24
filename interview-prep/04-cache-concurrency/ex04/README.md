# Concurrency Ex04: Fan-In Basico

Objetivo:

- coordinar dos consultas concurrentes
- usar channels sin complicar de mas la solucion

Enunciado:

Tienes dos funciones proveedoras:

- una trae precio
- otra trae stock

Completa `LoadSnapshot` para que:

- ejecute ambas al mismo tiempo
- espere ambos resultados
- combine todo en un `ProductSnapshot`
- si alguna falla, regrese error

Restriccion:

- no uses `sync.WaitGroup`
- resuelvelo con channels porque ese es el musculo que queremos practicar aqui
