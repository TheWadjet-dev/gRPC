# gRPC - Ejemplo en Go

Este proyecto implementa un servicio gRPC en Go que responde a solicitudes con un mensaje "¡Hola, [nombre]!".

## **Requisitos**

- Go instalado (v1.16+)
- Plugin `protoc` y su extensión para Go:
  ```bash
  go install google.golang.org/protobuf/cmd/protoc-gen-go@latest