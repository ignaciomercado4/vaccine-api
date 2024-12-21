## API REST para Gestión de Registros de Vacunación
### Requisitos Generales
Base de datos: PostgreSQL.
Lenguaje de programación: Go.
Modelos
El sistema debe manejar tres modelos principales:

1. Usuario (User)
id (integer): Identificador único del usuario.
name (string): Nombre del usuario (puede ser nulo).
email (string): Correo electrónico del usuario, utilizado para iniciar sesión.
password (string): Contraseña asociada al usuario.
2. Droga (Drug)
id (integer): Identificador único de la droga.
name (string): Nombre de la droga.
approved (boolean): Indica si la droga está aprobada para su uso.
min_dose (integer): Dosis mínima permitida.
max_dose (integer): Dosis máxima permitida.
available_at (datetime): Fecha desde la cual la droga está disponible para su uso.
3. Vacunación (Vaccination)
id (integer): Identificador único de la vacunación.
name (string): Nombre de la persona que recibirá la vacunación.
drug_id (integer): ID de la droga utilizada para la vacunación.
dose (integer): Dosis administrada en la vacunación.
date (datetime): Fecha y hora de la vacunación.
API REST
Se deben crear las siguientes rutas para gestionar los recursos de la API:

### Rutas de Autenticación (Auth):
- POST /signup: Crea un nuevo usuario proporcionando los campos name, email y password.
- POST /login: Inicia sesión usando email y password. Si las credenciales son correctas, devuelve un token JWT con un vencimiento definido en una variable de entorno.

Rutas de Drogas (Drugs):
Todas estas rutas requieren validación de un token JWT válido.

- POST /drugs: Crea una nueva droga, proporcionando los campos name, approved, min_dose, max_dose y available_at.
- PUT /drugs/:id: Actualiza los detalles de una droga existente.
- GET /drugs: Obtiene la lista de todas las drogas disponibles.
- DELETE /drugs/:id: Elimina una droga.

Rutas de Vacunación (Vaccination):
Todas estas rutas requieren validación de un token JWT válido.

- POST /vaccination: Registra una nueva vacunación, proporcionando los campos name, drug_id, dose y date. Se debe verificar que la dosis esté dentro del rango permitido y que la fecha de vacunación sea posterior a la fecha en que la droga esté disponible.
- PUT /vaccination/:id: Actualiza los detalles de una vacunación.
- GET /vaccination: Obtiene la lista de todas las vacunaciones registradas.
- DELETE /vaccination/:id: Elimina una vacunación.
