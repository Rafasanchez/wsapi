# WSApi

This is a Golang API project built with Gorm and Gin that connects to a MySQL database. It provides CRUD operations for various tables including Moneda, Pais, Empresa, Departamento, Municipio, Cliente, Proveedor, Producto, Caja, CajaMov, ListaPrecioh, ListaPreciod, FormaPago, Sucursal, Empleado, Ordenh, and Ordend. The API also includes JWT authentication for security.

## Project Structure

```
WSApi
├── cmd
│   └── main.go
├── pkg
│   ├── api
│   │   ├── handlers
│   │   │   ├── moneda.go
│   │   │   ├── pais.go
│   │   │   ├── empresa.go
│   │   │   ├── departamento.go
│   │   │   ├── municipio.go
│   │   │   ├── cliente.go
│   │   │   ├── proveedor.go
│   │   │   ├── producto.go
│   │   │   ├── caja.go
│   │   │   ├── cajamov.go
│   │   │   ├── listaprecioh.go
│   │   │   ├── listapreciod.go
│   │   │   ├── formapago.go
│   │   │   ├── sucursal.go
│   │   │   ├── empleado.go
│   │   │   ├── ordenh.go
│   │   │   └── ordend.go
│   │   ├── middleware
│   │   │   └── jwt.go
│   │   └── routes
│   │       └── routes.go
│   ├── config
│   │   └── db.go
│   └── models
│       ├── moneda.go
│       ├── pais.go
│       ├── empresa.go
│       ├── departamento.go
│       ├── municipio.go
│       ├── cliente.go
│       ├── proveedor.go
│       ├── producto.go
│       ├── caja.go
│       ├── cajamov.go
│       ├── listaprecioh.go
│       ├── listapreciod.go
│       ├── formapago.go
│       ├── sucursal.go
│       ├── empleado.go
│       ├── ordenh.go
│       └── ordend.go
├── .env
├── .gitignore
├── go.mod
├── go.sum
└── README.md
```

## Getting Started

To get started with the WSApi project, follow these steps:

1. Clone the repository: `git clone <repository-url>`
2. Navigate to the project directory: `cd WSApi`
3. Set up the MySQL database connection details in the `.env` file.
4. Install the dependencies: `go mod download`
5. Start the server: `go run cmd/main.go`

## API Endpoints

The following API endpoints are available:

- **Moneda**
  - `GET /monedas`: Get all monedas
  - `GET /monedas/:id`: Get a moneda by ID
  - `POST /monedas`: Create a new moneda
  - `PUT /monedas/:id`: Update a moneda
  - `DELETE /monedas/:id`: Delete a moneda

- **Pais**
  - `GET /paises`: Get all paises
  - `GET /paises/:id`: Get a pais by ID
  - `POST /paises`: Create a new pais
  - `PUT /paises/:id`: Update a pais
  - `DELETE /paises/:id`: Delete a pais

- **Empresa**
  - `GET /empresas`: Get all empresas
  - `GET /empresas/:id`: Get an empresa by ID
  - `POST /empresas`: Create a new empresa
  - `PUT /empresas/:id`: Update an empresa
  - `DELETE /empresas/:id`: Delete an empresa

- **Departamento**
  - `GET /departamentos`: Get all departamentos
  - `GET /departamentos/:id`: Get a departamento by ID
  - `POST /departamentos`: Create a new departamento
  - `PUT /departamentos/:id`: Update a departamento
  - `DELETE /departamentos/:id`: Delete a departamento

- **Municipio**
  - `GET /municipios`: Get all municipios
  - `GET /municipios/:id`: Get a municipio by ID
  - `POST /municipios`: Create a new municipio
  - `PUT /municipios/:id`: Update a municipio
  - `DELETE /municipios/:id`: Delete a municipio

- **Cliente**
  - `GET /clientes`: Get all clientes
  - `GET /clientes/:id`: Get a cliente by ID
  - `POST /clientes`: Create a new cliente
  - `PUT /clientes/:id`: Update a cliente
  - `DELETE /clientes/:id`: Delete a cliente

- **Proveedor**
  - `GET /proveedores`: Get all proveedores
  - `GET /proveedores/:id`: Get a proveedor by ID
  - `POST /proveedores`: Create a new proveedor
  - `PUT /proveedores/:id`: Update a proveedor
  - `DELETE /proveedores/:id`: Delete a proveedor

- **Producto**
  - `GET /productos`: Get all productos
  - `GET /productos/:id`: Get a producto by ID
  - `POST /productos`: Create a new producto
  - `PUT /productos/:id`: Update a producto
  - `DELETE /productos/:id`: Delete a producto

- **Caja**
  - `GET /cajas`: Get all cajas
  - `GET /cajas/:id`: Get a caja by ID
  - `POST /cajas`: Create a new caja
  - `PUT /cajas/:id`: Update a caja
  - `DELETE /cajas/:id`: Delete a caja

- **CajaMov**
  - `GET /cajamovs`: Get all cajamovs
  - `GET /cajamovs/:id`: Get a cajamov by ID
  - `POST /cajamovs`: Create a new cajamov
  - `PUT /cajamovs/:id`: Update a cajamov
  - `DELETE /cajamovs/:id`: Delete a cajamov

- **ListaPrecioh**
  - `GET /listapreciohs`: Get all listapreciohs
  - `GET /listapreciohs/:id`: Get a listaprecioh by ID
  - `POST /listapreciohs`: Create a new listaprecioh
  - `PUT /listapreciohs/:id`: Update a listaprecioh
  - `DELETE /listapreciohs/:id`: Delete a listaprecioh

- **ListaPreciod**
  - `GET /listapreciods`: Get all listapreciods
  - `GET /listapreciods/:id`: Get a listapreciod by ID
  - `POST /listapreciods`: Create a new listapreciod
  - `PUT /listapreciods/:id`: Update a listapreciod
  - `DELETE /listapreciods/:id`: Delete a listapreciod

- **FormaPago**
  - `GET /formaspagos`: Get all formaspagos
  - `GET /formaspagos/:id`: Get a formapago by ID
  - `POST /formaspagos`: Create a new formapago
  - `PUT /formaspagos/:id`: Update a formapago
  - `DELETE /formaspagos/:id`: Delete a formapago

- **Sucursal**
  - `GET /sucursales`: Get all sucursales
  - `GET /sucursales/:id`: Get a sucursal by ID
  - `POST /sucursales`: Create a new sucursal
  - `PUT /sucursales/:id`: Update a sucursal
  - `DELETE /sucursales/:id`: Delete a sucursal

- **Empleado**
  - `GET /empleados`: Get all empleados
  - `GET /empleados/:id`: Get an empleado by ID
  - `POST /empleados`: Create a new empleado
  - `PUT /empleados/:id`: Update an empleado
  - `DELETE /empleados/:id`: Delete an empleado

- **Ordenh**
  - `GET /ordenhs`: Get all ordenhs
  - `GET /ordenhs/:id`: Get an ordenh by ID
  - `POST /ordenhs`: Create a new ordenh
  - `PUT /ordenhs/:id`: Update an ordenh
  - `DELETE /ordenhs/:id`: Delete an ordenh

- **Ordend**
  - `GET /ordends`: Get all ordends
  - `GET /ordends/:id`: Get an ordend by ID
  - `POST /ordends`: Create a new ordend
  - `PUT /ordends/:id`: Update an ordend
  - `DELETE /ordends/:id`: Delete an ordend

## Authentication

The API endpoints are protected with JWT authentication. To access the protected endpoints, include the JWT token in the `Authorization` header of the request.

## Dependencies

The project uses the following dependencies:

- Gorm: A Go ORM library for database access
- Gin: A web framework for building APIs in Go

## Contributing

Contributions are welcome! If you have any suggestions, bug reports, or feature requests, please open an issue or submit a pull request.

## License

This project is licensed under the [MIT License](LICENSE).