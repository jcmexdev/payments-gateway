# Etapa de construcción
FROM golang:1.22-alpine as builder

# Establece el directorio de trabajo dentro del contenedor
WORKDIR /go/src/app

# Copia el código fuente de tu aplicación al contenedor
COPY . .

# Descarga las dependencias del módulo
RUN go mod download

# Compila la aplicación estáticamente
RUN CGO_ENABLED=0 GOOS=linux go build -o payments-app ./payments/cmd/main.go

# Etapa de producción
FROM alpine:latest

# Establece el directorio de trabajo dentro del contenedor
WORKDIR /app

# Copia el binario compilado desde la etapa de construcción al contenedor de producción
COPY --from=builder /go/src/app/payments-app .

# Expone el puerto en el que se ejecuta tu aplicación (ajusta según sea necesario)
EXPOSE 8080

# Comando por defecto para ejecutar tu aplicación
CMD ["./payments-app"]
