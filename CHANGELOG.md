# Changelog

Todos los cambios de este proyecto se documentarán en este archivo.

## \[0.1.0] - 2025-06-06

### Added

* Estructura inicial del proyecto con Fiber y Go.

* Conexión a Firestore para almacenamiento de datos.

* Modelos `User` y `Task` con campos para:

  * `nombre`
  * `apellidos`
  * `email`
  * `contrasena`
  * `fecha_nacimiento`
  * `pregunta_secreta`
  * `respuesta_secreta`
  * `titulo`
  * `descripcion`
  * `fecha_inicio`
  * `deadline`
  * `usuario_id`

* Autenticación JWT con tokens de 10 minutos de validez.

* Endpoints para usuarios:

  * `POST /api/register`: Registro de usuarios.

    * Ejemplo de cuerpo:

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
  * `POST /api/login`: Login con JWT.
  * `POST /api/recover-password`: Recuperación de contraseña.
  * `GET /api/users/me`: Obtener datos del usuario.
  * `PUT /api/users/me`: Actualizar usuario.
  * `DELETE /api/users/me`: Eliminar usuario.

* Endpoints para tareas:

  * `POST /api/tasks`: Crear tarea.
  * `GET /api/tasks`: Listar tareas.
  * `PUT /api/tasks/:id`: Actualizar tarea.
  * `DELETE /api/tasks/:id`: Eliminar tarea.

* Middleware de autenticación JWT para rutas protegidas.

* Pruebas unitarias iniciales para el endpoint de registro (`test/user_test.go`).

* Documentación inicial con `README.md` y `CHANGELOG.md`.

### Fixed

* Corrección de imports no utilizados en `handlers/user.go` y `handlers/task.go`.
* Eliminación de texto no válido ("Kernighan–Ritchie") en `handlers/task.go`.
