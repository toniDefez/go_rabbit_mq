# 🐇 RabbitMQ Direct Exchange Example (Go)

> Basado en el Tutorial 4 oficial de RabbitMQ.

---

## 🎯 Objetivo

- Un **producer** envía mensajes con diferentes `routing keys`: `info`, `warn`, `error`.
- Dos **consumers** se suscriben a distintas `binding keys`:
  - **C₁** escucha solo `error`.
  - **C₂** escucha `error` y `warn`.

---

## Cómo ejecutarlo

### 1. Ejecuta el consumer C₁ (escucha `error`):

# Terminal 1
go run consumer.go error


### 2. Ejecuta el consumer C₂ (escucha `warn` y `error`):

# Terminal 2
go run consumer.go warn error


### 3. Publica mensajes:
# Terminal 3
go run producer.go error
go run producer.go warn
go run producer.go info


---

## ✅ Resultado esperado

- El mensaje con `error` lo reciben **C₁** y **C₂**
- El mensaje con `warn` lo recibe solo **C₂**
- El mensaje con `info` no lo recibe nadie (ninguna cola está enlazada con `info`)

