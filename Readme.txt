# Tienda Web Fullstack

Bienvenido a **Tienda Web Fullstack**, un proyecto de e-commerce que combina **Backend en Node.js + Express** con **Frontend en Go**. Esta aplicación permite:

- Gestionar productos y stock.
- Administrar órdenes y su estado (pendiente/entregada).
- Registrar usuarios y diferenciar roles (admin y usuario normal).
- Realizar compras desde el frontend con carrito de compras.

---

## Tecnologías

**Backend:**

- Node.js
- Express
- MongoDB con Mongoose
- JWT para autenticación
- Bcrypt para encriptar contraseñas
- Pug para templates

**Frontend:**

- Go con net/http
- MongoDB
- Templates HTML
- Manejo de carrito de compras

---

## Funcionalidades

### Backend

- Login y logout de usuarios con roles.
- CRUD de productos (solo admin).
- Gestión de órdenes: crear, ver, marcar como entregada.
- Panel de administración para revisar productos y órdenes.
- Autenticación con JWT y cookies.

### Frontend

- Mostrar productos desde la base de datos.
- Agregar productos al carrito.
- Comprar productos y generar órdenes automáticamente.
- Vistas de confirmación de pedido y carrito.

---

## Instalación

### Backend

1. Clonar repositorio.
2. Instalar dependencias:

```bash
npm install
