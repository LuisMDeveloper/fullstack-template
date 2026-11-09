# fullstack-template

Backend en Go (`api/`) y frontend en Vue (`web/`). El frontend se compila y se embebe en el binario de Go.

## Requisitos del entorno

Herramientas que deben estar instaladas en tu máquina. Las versiones son con las que se desarrolló el proyecto.

| Herramienta | Versión | Para qué | Obligatoria |
| --- | --- | --- | --- |
| [Go](https://go.dev/dl/) | 1.27.3 | Compilar el backend | Sí |
| [Node.js](https://nodejs.org/) | 26.13.0 | Compilar el frontend | Sí |
| npm | 11.20.0 | Instalar dependencias de `web/` (viene con Node) | Sí |
| [GNU Make](https://www.gnu.org/software/make/) | 4.4.1 | Comandos de la raíz (`make build`, `make run`) | Sí |
| [pnpm](https://pnpm.io/) | 12.18.0 | Instalar las herramientas globales de commits | Solo para commitear |
| [Lefthook](https://github.com/evilmartians/lefthook) | 2.2.1 | Git hooks (escrito en Go) | Solo para commitear |
| [commitlint](https://commitlint.js.org/) | 21.2.4 | Valida los mensajes de commit | Solo para commitear |

Las dependencias de `api/` (`go.mod`) y de `web/` (`package.json`) los instalan Go y npm solos.

### Herramientas globales para commits

commitlint **no** es dependencia del proyecto: se usa instalado de forma global con pnpm.

```sh
# Lefthook: instálalo según https://lefthook.dev/installation/ (gestor de paquetes o `go install`)
pnpm add -g @commitlint/cli @commitlint/config-conventional
lefthook install                                         # una vez por clon, activa los hooks
```

`$(go env GOPATH)/bin` y el directorio de binarios globales de pnpm deben estar en tu `PATH`.

## Uso

Desde la raíz:

```sh
make build   # npm ci + build de web/ -> api/internal/webui/dist, luego compila Go en bin/api
make run     # build y ejecuta el servidor en http://localhost:8080
make clean   # borra bin/ y la salida del frontend
```

Desarrollo del frontend con recarga en caliente (el proxy envía `/api` a `:8080`):

```sh
cd api && go run ./cmd/api   # necesita haber corrido `make build` al menos una vez
cd web && npm run dev
```

## API

| Método | Ruta | Respuesta |
| --- | --- | --- |
| GET | `/api/health` | `{"status":"ok"}` |
| GET | `/api/hello` | `{"message":"Hello, World!"}` |
| GET | `/api/info` | JSON estático de prueba |

Cualquier otra ruta bajo `/api/` responde 404 en JSON. El resto de rutas sirve el frontend.

## Commits

Se sigue [Conventional Commits](https://www.conventionalcommits.org/): `tipo(scope): descripción`, por ejemplo `feat(web): add login page`.

Tipos: `feat`, `fix`, `docs`, `style`, `refactor`, `perf`, `test`, `build`, `ci`, `chore`.

Lefthook ejecuta estos hooks (ver `lefthook.yml`):

- `pre-commit`: `gofmt`, `go vet`, `oxfmt --check` y `oxlint`.
- `commit-msg`: commitlint con `commitlint.config.mjs`.

## Notas

- `go build` falla en un clon nuevo hasta que se compile el frontend, porque `api/internal/webui/dist` no está en git y `go:embed` lo necesita. `make build` lo resuelve.
- Los tests: `cd api && go test ./...` y `cd web && npm run test:unit`.
