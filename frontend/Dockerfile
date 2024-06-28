# Dockerfile para el frontend en Next.js
FROM node:14-alpine AS builder

WORKDIR /frontend

# Instalar dependencias
COPY package.json package-lock.json ./
RUN npm install

# Copiar el código
COPY . .

# Compilar la aplicación para producción
RUN npm run build

# Imagen final mínima
FROM nginx:alpine

# Copiar los archivos de construcción de Next.js al servidor nginx
COPY --from=builder /frontend/out /usr/share/nginx/html

# Configurar el servidor nginx
EXPOSE 80
CMD ["nginx", "-g", "daemon off;"]
