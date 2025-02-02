# GraphQL API microservice using GORM and gqlgen

### Overview
This project provides a `GraphQL` API microservice built with `Go`,
utilizing `gqlgen` for code generation,
`genqlient` for integration testing,
and `GORM` for database interactions.
It supports containerization with `Docker`
and deployment to a `Kubernetes` cluster.

## Get started
You can use this project (clone required for 2-5) in the following ways (ordered by difficulty):
### 1. Deployed instance in Kubernetes on GCP.
> Note: possible unavailable (deployment or DB)

Playground: [http://34.147.13.187](http://34.147.13.187)
### 2. Docker with an external database `galaxy`. Config: [compose-galaxy.yaml](https://github.com/kalalannn/go-graphql_galaxy/blob/main/config/compose-galaxy.yaml)
> Note: possible unavailable (DB)
<details>
<summary><a>Expand ... </a></summary>

* Prerequisite: `docker` ideally with [buildx](https://github.com/docker/buildx)
* Start (Up)

```bash
make build_app_image up_galaxy
```
* Finish (Down)

```bash
make down_galaxy
```
</details>

### 3. Locally with an external database `galaxy`. Config: [local-galaxy.yaml](https://github.com/kalalannn/go-graphql_galaxy/blob/main/config/local-galaxy.yaml)
> Note: possible unavailable (DB)
<details>
<summary><a>Expand ... </a></summary>

* Prerequisite: `go >= 1.23`
* Start (Up)
```bash
make mod_download local_run_galaxy
```
* Finish (Down)
```bash
^C
```
</details>

### 4. Docker with an internal database `postgres`. Config: [compose-local.yaml](https://github.com/kalalannn/go-graphql_galaxy/blob/main/config/compose-local.yaml)
> Note: always available, port=8081
<details>
<summary><a>Expand ... </a></summary>

* Prerequisite: `docker` ideally with [buildx](https://github.com/docker/buildx)
* Start (Up)
```bash
make up_local
make migrate_db    # Run once
```

* Finish (Down)
```bash
make down_local
make clean_db      # if needed
```
</details>

### 5. Locally with an internal database `postgres`. Config: [local.yaml](https://github.com/kalalannn/go-graphql_galaxy/blob/main/config/local.yaml)
> Note: always available
<details>
<summary><a>Expand ... </a></summary>

* Prerequisite: `go >= 1.23`, `docker` ideally with [buildx](https://github.com/docker/buildx)
* Start (Up)
```
make up_db_only
make migrate_db    # Run once
make local_run
```
* Finish (Down)
```bash
^C
make down_db_only
make clean_db      # if needed
```
</details>

## API Endpoints
- **`GET /`**: GraphQL Playground
- **`GET /ping`**: Server healthcheck
- **`GET /graphql`**: GraphQL endpoint

## Configuration
The service uses the `APP_CONFIG_PATH` environment variable to locate the configuration file. The configuration file should be structured as follows:
```yaml
env: local                     # Current environment (used for logging, etc.)
server:
  host: 0.0.0.0                # HTTP server host
  port: 8080                   # HTTP server port
  use_playground: true         # use playground endpoint ?
  use_introspection: true      # enable introspection ?
  gql_complexity_limit: 100    # graphql complexity limit
  gql_depth_limit: 10          # max request depth limit >=
database:
  user: postgres               # DB_USER
  password: postgres           # DB_PASSWORD
  host: localhost              # DB_HOST
  port: 5432                   # DB_PORT
  db_name: postgres            # DB_NAME
  sslmode: disable             # SSL policy
  timezone: UTC                # DB timezone
```

## GraphQL
> Note: Check introspection for better experience
### Schemas: [internal/graphql/schemas](https://github.com/kalalannn/go-graphql_galaxy/tree/main/internal/graphql/schemas)

### Entities
<details>
<summary><a>Expand ... </a></summary>

```graphql
type Character {
  id: ID!
  name: String!
  gender: String
  ability: String!
  minimal_distance: Float!
  weight: Float
  born: String!
  in_space_since: String
  beer_consumption: Int!
  knows_the_answer: Boolean!
  nemeses: [Nemesis!]!
}

type Genders {
  male: Int64!
  female: Int64!
  other: Int64!
}

type Nemesis {
  id: ID!
  is_alive: Boolean!
  years: Int
  character: Character!
  secrets: [Secret!]!
}

type AliveNemeses {
  alive: Int64!
  dead: Int64!
}

type Secret {
  id: ID!
  secret_code: Int64!
  nemesis: Nemesis!
}
```
</details>

### Queries
<details>
<summary><a>Expand ... </a></summary>

```graphql
type Query {
  server_time: String!
  health_check: Boolean!
  # Character
  average_age: Float!
  average_weight: Float!
  average_beer_consumption: Float!
  characters_count: Int64!
  characters(
    orderBy: CharacterOrderBy = {field: id, direction: ASC}, # default
    pagination: PaginationInput
  ): [Character!]!
  character(id: ID!): Character
  genders: Genders!
  # Nemesis
  nemeses_count: Int64!
  average_nemeses_years: Float!
  nemeses(
    orderBy: NemesisOrderBy = {field: id, direction: ASC}, # default
    pagination: PaginationInput
  ): [Nemesis!]!
  nemesis(id: ID!): Nemesis
  alive_nemeses: AliveNemeses
  # Secret
  secrets_count: Int64!
  secrets(
    orderBy: SecretOrderBy = {field: id, direction: ASC}, # default
    pagination: PaginationInput
  ): [Secret!]!
  secret(id: ID!): Secret
}
```
</details>


## Project structure
<details>
<summary><a>Expand ... </a></summary>

```bash
├── cmd                                # entrypoints (main)
│   ├── generate                       # GQLGen go code generator
│   └── server                         # GraphQL server main
├── config                             # Server configuration files for environments
│   └── ...
├── docker-compose-(local|galaxy).yaml # Docker-compose for db: local (postgres) and external (galaxy)
├── Dockerfile.(core|base|app)         # Core, Base and App Dockerfiles
├── gqlgen.yml                         # GQLGen configuration
├── genqlient.yaml                     # GenQlient configuration
├── go.mod && go.sum                   # Go modules, versions and checksums
├── internal                           # Internal modules (business logic)
│   ├── app                            # Application module
│   ├── genqlient                      # GenQlient module
│   │   ├── generated                  # generated by GenQlient
│   │   └── operations                 # GenQlient client's queiries (for integration tests)
│   │       └── ...
│   ├── gorm                           # GORM module
│   │   ├── entities                   # Entities module (Character, Nemesis, Secret)
│   │   │   ├── ...
│   │   └── services                   # Services module (entities management, DB operations)
│   │       └── ...
│   ├── gqlcontext                     # gqlcontext processor (for Depth extension and DB preloads)
│   ├── graphql                        # GraphQL module
│   │   ├── generated                  # generated by GQLGen
│   │   ├── models                     # generated Go models from GraphQL schemas
│   │   │   └── ...
│   │   ├── resolvers                  # GraphQL query resolvers
│   │   │   └── ...
│   │   └── schemas                    # GraphQL schemas (for generation)
│   │       └── ...
│   ├── server                         # HTTP Server configuration (routes, handlers, extensions)
│   │   └── ...
│   └── transformers                   # Transform GORM DB Entities to generated GraphQL models
│       └── ...
├── Makefile                           # Makefile for most operations
├── pkg                                # Shared modules (no business logic)
│   └── ...
├── sql                                # Raw PostgreSQL files for psql
│   └── ...
└── tests                              # Integration tests folder
    └── ...
```
</details>

## Data structure
### DB Table: `Character`, GORM model: [Character](https://github.com/kalalannn/go-graphql_galaxy/blob/main/internal/gorm/entities/character.go), GraphQL schema: [Character](https://github.com/kalalannn/go-graphql_galaxy/blob/main/internal/graphql/schemas/character.graphql#L1-L13)

<details>
<summary><a>Expand ... </a></summary>
    
```sql
    Column        |         Type         | NULLable?
------------------+----------------------+----------
 id               | integer              | not null
 name             | text                 | not null
 gender           | text                 |         
 ability          | text                 | not null
 minimal_distance | numeric              | not null
 weight           | numeric              |         
 born             | timestamp without tz | not null
 in_space_since   | timestamp without tz |         
 beer_consumption | integer              | not null
 knows_the_answer | boolean              | not null
----------------------------------------------------
Referenced by:
  TABLE "nemesis" CONSTRAINT "character" FOREIGN KEY (character_id) REFERENCES "character"(id) NOT VALID
```
</details>

### DB Table: `Nemesis`, GORM model: [Nemesis](https://github.com/kalalannn/go-graphql_galaxy/blob/main/internal/gorm/entities/nemesis.go), GraphQL schema: [Nemesis](https://github.com/kalalannn/go-graphql_galaxy/blob/main/internal/graphql/schemas/nemesis.graphql#L1-L7)
<details>
<summary><a>Expand ... </a></summary>
    
```sql
    Column    |  Type   | NULLable?
--------------+---------+----------
 is_alive     | boolean | not null 
 years        | integer |
 id           | integer | not null 
 character_id | integer |
-----------------------------------
Foreign keys:
    "character" FOREIGN KEY (character_id) REFERENCES "character"(id) NOT VALID
Referenced by:
    TABLE "secret" CONSTRAINT "nemesis" FOREIGN KEY (nemesis_id) REFERENCES nemesis(id)
```
</details>

### DB Table: `Secret`, GORM model: [Secret](https://github.com/kalalannn/go-graphql_galaxy/blob/main/internal/gorm/entities/secret.go), GraphQL schema: [Secret](https://github.com/kalalannn/go-graphql_galaxy/blob/main/internal/graphql/schemas/secret.graphql#L1-L5)
<details>
<summary><a>Expand ... </a></summary>
   
```sql
   Column    |  Type   | NULLable?
-------------+---------+----------
 id          | integer | not null 
 secret_code | bigint  | not null 
 nemesis_id  | integer | not null 
----------------------------------
Foreign keys:
    "nemesis" FOREIGN KEY (nemesis_id) REFERENCES nemesis(id)
```
</details>

## Examples
### RootQuery
<details>
<summary>Expand ...</summary>

* Request:
```graphql
query RootQuery {
  server_time
  health_check
  characters_count
  average_age
  average_weight
  average_beer_consumption
  nemeses_count
  average_nemeses_years
  secrets_count
  genders {
    male
    female
    other
  }
  alive_nemeses {
    alive
    dead
  }
}
```
* Response:
```json
{
  "data": {
    "server_time": "2025-02-02T13:54:06Z",
    "health_check": true,
    "characters_count": 11,
    "average_age": 40.91,
    "average_weight": 104.03,
    "average_beer_consumption": 134527.91,
    "nemeses_count": 13,
    "average_nemeses_years": 113.75,
    "secrets_count": 25,
    "genders": {
      "male": 6,
      "female": 2,
      "other": 3
    },
    "alive_nemeses": {
      "alive": 11,
      "dead": 2
    }
  }
}
```
</details>

### GetOnes
<details>
<summary>Expand ...</summary>

* Request:
```graphql
query GetOnes {
  character(id: 2) {
    id
    ability
    beer_consumption
    born
    gender
    in_space_since
    knows_the_answer
    minimal_distance
    name
    weight
  }
  nemesis(id: 8) {
    id
    years
    is_alive
  }
  secret(id: 7) {
    id
    secret_code
  }
}
```
* Response:
```json
{
  "data": {
    "character": {
      "id": "2",
      "ability": "mathematician",
      "beer_consumption": 6704,
      "born": "1994-12-14T00:00:00Z",
      "gender": "female",
      "in_space_since": "2014-12-24T17:21:50Z",
      "knows_the_answer": true,
      "minimal_distance": 6.2,
      "name": "Trillian",
      "weight": 49
    },
    "nemesis": {
      "id": "8",
      "years": 2,
      "is_alive": true
    },
    "secret": {
      "id": "7",
      "secret_code": 9449428626
    }
  }
}
```
</details>

### NotFoundOnes
> Note: error message should be different (regarding to business logic)

<details>
<summary>Expand ...</summary>

* Request:
```graphql
query NotFoundOnes {
  character(id: 9999) {
    id
  }
  nemesis(id: 9999) {
    id
    years
    is_alive
  }
  secret(id: 9999) {
    id
  }
}
```
* Response:
```json
{
  "errors": [
    {
      "message": "record not found",
      "path": [
        "character"
      ]
    },
    {
      "message": "record not found",
      "path": [
        "secret"
      ]
    },
    {
      "message": "record not found",
      "path": [
        "nemesis"
      ]
    }
  ],
  "data": {
    "character": null,
    "nemesis": null,
    "secret": null
  }
}
```
</details>

### SortPaginationQuery
<details>
<summary>Expand ...</summary>

* Request:
```graphql
query SortPaginationQuery {
  characters(
    orderBy: { field: name, direction: ASC },
    pagination: {limit: 1})
  {
    id
    name
    beer_consumption
  }
  nemeses(
    orderBy: {field: years, direction: ASC},
    pagination: {limit: 2})
  {
    id
    years
  }
  secrets(
    orderBy: {field: secret_code, direction: DESC},
    pagination: {limit: 3})
  {
    id
    secret_code
  }
}
```
* Response:
```json
{
  "data": {
    "characters": [
      {
        "id": "9",
        "name": "Alice Beeblebrox",
        "beer_consumption": 64
      }
    ],
    "nemeses": [
      {
        "id": "8",
        "years": 2
      },
      {
        "id": "2",
        "years": 28
      }
    ],
    "secrets": [
      {
        "id": "7",
        "secret_code": 9449428626
      },
      {
        "id": "6",
        "secret_code": 9442445871
      },
      {
        "id": "24",
        "secret_code": 8424742058
      }
    ]
  }
}
```
</details>

### RecursiveChildren
<details>
<summary>Expand ...</summary>

* SQL:
```sql
SELECT
  c.id AS c_id, c.name,
  n.id AS n_id, n.is_alive, n.years,
  s.id AS s_id, s.secret_code AS s_code
FROM character c
  JOIN nemesis n ON c.id = n.character_id
  JOIN secret  s ON n.id = s.nemesis_id where c.id = 12;

 c_id |  name   | n_id | is_alive | years | s_id |   s_code
------+---------+------+----------+-------+------+------------
   12 | Frankie |    7 | t        |   953 |   13 | 5467717091
   12 | Frankie |    7 | t        |   953 |   14 | 4166492176
   12 | Frankie |    8 | t        |     2 |   15 | 6271440484
   12 | Frankie |    8 | t        |     2 |   16 | 6275689247
```
* Request:
```graphql
query RecursiveChildren {
  character(id: 12) {
    id
    nemeses{
      id
      secrets{
        id
        nemesis{
          id
          character{
            id
          }
        }
      }
    }
  }
}
```
* Response:
```json
{
  "data": {
    "character": {
      "id": "12",
      "nemeses": [
        {
          "id": "7",
          "secrets": [
            {
              "id": "13",
              "nemesis": {
                "id": "7",
                "character": {
                  "id": "12"
                }
              }
            },
            {
              "id": "14",
              "nemesis": {
                "id": "7",
                "character": {
                  "id": "12"
                }
              }
            }
          ]
        },
        {
          "id": "8",
          "secrets": [
            {
              "id": "15",
              "nemesis": {
                "id": "8",
                "character": {
                  "id": "12"
                }
              }
            },
            {
              "id": "16",
              "nemesis": {
                "id": "8",
                "character": {
                  "id": "12"
                }
              }
            }
          ]
        }
      ]
    }
  }
}
```
</details>

### MaxPossibleDepth
> Note: configurable via `config.server.gql_depth_limit`

<details>
<summary>Expand ...</summary>

* Request:
```graphql
query MaxPossibleDepth {
  character(id: 2) {
    id
    nemeses {
      id
      character{
        id
        nemeses{
          id
          character{
            id
            nemeses{
              id
              character{
                id
                nemeses{
                  id
                  character{
                    id
                  }
                }
              }
            }
          }
        }
      }
    }
  }
}
```
* Response:
```json
{
  "data": {
    "character": {
      "id": "2",
      "nemeses": [
        {
          "id": "1",
          "character": {
            "id": "2",
            "nemeses": [
              {
                "id": "1",
                "character": {
                  "id": "2",
                  "nemeses": [
                    {
                      "id": "1",
                      "character": {
                        "id": "2",
                        "nemeses": [
                          {
                            "id": "1",
                            "character": {
                              "id": "2"
                            }
                          }
                        ]
                      }
                    }
                  ]
                }
              }
            ]
          }
        }
      ]
    }
  }
}
```
</details>

### DepthLimitError
> Note: configurable via `config.server.gql_depth_limit`

<details>
<summary>Expand ...</summary>

* Request:
```graphql
query DepthLimitError {
  character(id: 2) {
    id
    nemeses {
      id
      character{
        id
        nemeses{
          id
          character{
            id
            nemeses{
              id
              character{
                id
                nemeses{
                  id
                  character{
                    id
                    nemeses{ // depth >= MaxDepthLimit
                      id
                    }
                  }
                }
              }
            }
          }
        }
      }
    }
  }
}
```
* Response:
```json
{
  "errors": [
    {
      "message": "Max depth limit exceeded >= 10",
      "extensions": {
        "code": "DEPTH_LIMIT_EXCEEDED"
      }
    }
  ],
  "data": null
}
```
</details>

<!--
### Template
<details>
<summary>Expand ...</summary>

* Request:
```graphql
```
* Response:
```json
```
</details>
-->

## New entities guide (howto add new entity?)
1. Create DB table
2. Create GraphQL schema into `internal/graphql/schemas`
3. Regenerate Go code (`models` && `resolvers`) with `make generate`
4. Implement GORM `entity model` && `service` in `internal/gorm`
5. Implement `<entity>_transformer.go`
6. Implement generated `<entity>.resolvers.go`

## Deployment guide
* Local build
  1. Build docker image with `make build_app_image`
  2. Tag and push image to your `Cloud Artifact Registry`
* Cloud build
  1. Setup Cloud image build (`CD/CI`)

Finally Deploy `App && Service` to `Kubernetes Cluster` using `kubectl apply -f .kube/`

## Future Development (TODOs)
* Unit tests
* Extend integration tests
* Mutations
