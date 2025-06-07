# Task Manager API

Una API REST desarrollada en Go usando el framework Fiber, con autenticación JWT y almacenamiento en Firestore. Permite gestionar usuarios y tareas (CRUD) con recuperación de contraseña mediante pregunta secreta.

## Características

### Gestión de usuarios:
- Registro, login, actualización y eliminación de usuarios.
- Recuperación de contraseña usando pregunta y respuesta secreta.

### Gestión de tareas:
- Crear, listar, actualizar y eliminar tareas asociadas a un usuario.

### Autenticación:
- Tokens JWT con validez de 10 minutos.

### Base de datos:
- Firestore para el almacenamiento.

### Estructura modular:
- Código organizado con buenas prácticas (basado en el ejemplo del profesor Emmanuel), separando modelos, handlers, rutas, middleware y utilidades.

---

## Estructura del proyecto

task-manager-api/
├── main.go              # Punto de entrada de la aplicación
├── go.mod               # Módulo Go y dependencias
├── config/              # Configuración de conexión a Firestore
├── models/              # Estructuras de datos (User, Task)
├── handlers/            # Lógica para los endpoints
├── routes/              # Definición de rutas de la API
├── middleware/          # Middleware para autenticación JWT
├── utils/               # Funciones auxiliares (JWT)
├── test/                # Prueba automatizada
├── README.md            # Documentación del proyecto
├── CHANGELOG.md         # Registro de cambios
└── serviceAccountKey.json # Credenciales de Firestore
└── .gitignore           # Archivos a ignorar

---

## Requisitos

- **Go**: Versión 1.20 o superior.
- **Firestore**: Cuenta de Firebase con un proyecto configurado y archivo de credenciales (`serviceAccountKey.json`).
- **Postman** o **curl**: Para probar los endpoints.

---

## Instalación

1. **Clonar el repositorio:**

   ```bash
   git clone <URL_DEL_REPOSITORIO>
   cd task-manager-api
   ```

2. **Instalar dependencias:**

   ```bash
   go mod tidy
   ```

3. **Configurar Firestore:**

   - Descarga el archivo de credenciales (`serviceAccountKey.json`) desde tu proyecto de Firebase.
   - Colócalo en la raíz del proyecto.

4. **Ejecutar el proyecto:**

   ```bash
   go run main.go
   ```

   El servidor se iniciará en [http://localhost:3000](http://localhost:3000).

---

## Endpoints

### Usuarios

#### POST `/api/register`
Registra un nuevo usuario.

**Ejemplo de cuerpo:**

```json
{
  "nombre": "Ariadna",
  "apellidos": "López",
  "email": "ariadna.lopez@example.com",
  "contrasena": "ariadna123",
  "fecha_nacimiento": "2000-05-15T00:00:00Z",
  "pregunta_secreta": "¿Cuál es tu comida favorita?",
  "respuesta_secreta": "Tacos"
}
```

#### POST `/api/login`
Autentica un usuario y devuelve un token JWT.

**Ejemplo de cuerpo:**

```json
{
  "email": "ariadna.lopez@example.com",
  "contrasena": "ariadna123"
}
```

#### POST `/api/recover-password`
Recupera la contraseña usando pregunta secreta.

**Ejemplo de cuerpo:**

```json
{
  "email": "ariadna.lopez@example.com",
  "pregunta_secreta": "¿Cuál es tu comida favorita?",
  "respuesta_secreta": "Tacos",
  "nueva_contrasena": "nuevaContrasena123"
}
```

#### GET `/api/users/me`
Obtiene los datos del usuario autenticado (requiere token).

#### PUT `/api/users/me`
Actualiza los datos del usuario (requiere token).

#### DELETE `/api/users/me`
Elimina el usuario (requiere token).

---

### Tareas

#### POST `/api/tasks`
Crea una nueva tarea (requiere token).

**Ejemplo de cuerpo:**

```json
{
  "titulo": "Tarea 1",
  "descripcion": "Hacer la tarea de desarrollo web",
  "fecha_inicio": "2025-06-06T10:00:00Z",
  "deadline": "2025-06-07T23:59:59Z"
}
```

#### GET `/api/tasks`
Lista las tareas del usuario autenticado (requiere token).

#### PUT `/api/tasks/:id`
Actualiza una tarea (requiere token).

#### DELETE `/api/tasks/:id`
Elimina una tarea (requiere token).

---

## Notas:

- Las rutas protegidas requieren el encabezado `Authorization: Bearer <TOKEN_JWT>`.
- Los tokens JWT expiran cada 10 minutos.

---

## Probar la API

### Con Postman

1. Crea una colección en Postman (por ejemplo, "Task Manager API").
2. Configura las solicitudes para cada endpoint (ver ejemplos arriba).
3. Añade el encabezado `Content-Type: application/json` para solicitudes POST y PUT.
4. Para rutas protegidas, incluye el token JWT en `Authorization: Bearer <TOKEN_JWT>`.
5. Guarda las solicitudes en la colección para reutilizarlas.

---

## Pruebas automatizadas

Ejecuta las pruebas unitarias incluidas:

```bash
go test ./test
```

---

## Configuración de Firestore

- Asegúrate de que las colecciones `users` y `tasks` estén habilitadas en Firestore (se crean automáticamente al guardar el primer documento).
- Verifica que el archivo `serviceAccountKey.json` esté en la raíz del proyecto.

---

## Dependencias

- `github.com/gofiber/fiber/v2`: Framework web.
- `cloud.google.com/go/firestore`: Cliente de Firestore.
- `firebase.google.com/go/v4`: SDK de Firebase.
- `github.com/golang-jwt/jwt/v5`: Generación y validación de JWT.
- `golang.org/x/crypto`: Encriptación de contraseñas.

---

## Notas de seguridad

- **Clave JWT**: La clave secreta para JWT está en `utils/jwt.go`. En producción, usa una variable de entorno.
- **Credenciales de Firestore**: No incluyas `serviceAccountKey.json` en el control de versiones (agrega al `.gitignore`).

---