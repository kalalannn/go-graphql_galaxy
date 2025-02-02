# GraphQL API microservice using GORM, gqlgen and genqlient

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
### 2. Docker with an external database `galaxy`. Config: <u>[compose-galaxy.yaml](https://github.com/kalalannn/go-graphql_galaxy/blob/main/config/compose-galaxy.yaml)</u>
> Note: possible unavailable (DB)
<details>
<summary>Click to expand</summary>

* Prerequisite: `docker` ideally with [buildx](https://github.com/docker/buildx)
* Start (Up)
<pre><code>make build_app_image up_galaxy</code></pre>
* Finish (Down)
<pre><code>make down_galaxy</code></pre>
</details>

### 3. Locally with an external database `galaxy`. Config: <u>[local-galaxy.yaml](https://github.com/kalalannn/go-graphql_galaxy/blob/main/config/local-galaxy.yaml)</u>
> Note: possible unavailable (DB)
<details>
<summary>Click to expand</summary>

* Prerequisite: `go >= 1.23`
* Start (Up)
<pre><code>make mod_download local_run_galaxy</code></pre>
* Finish (Down)
<pre><code>^C</code></pre>
</details>

### 4. Docker with an internal database `postgres`. Config: <u>[compose-local.yaml](https://github.com/kalalannn/go-graphql_galaxy/blob/main/config/compose-local.yaml)</u>
> Note: always available
<details>
<summary>Click to expand</summary>

* Prerequisite: `docker` ideally with [buildx](https://github.com/docker/buildx)
* Start (Up)
<pre><code class="language-bash">make up_local
make migrate_db # Run once
</code></pre>
* Finish (Down)
<pre><code class="language-bash">make down_local
make clean_db # if needed
</code></pre>
</details>

### 5. Locally with an internal database `postgres`. Config: <u>[local.yaml](https://github.com/kalalannn/go-graphql_galaxy/blob/main/config/local.yaml)</u>
> Note: always available
<details>
<summary>Click to expand</summary>

* Prerequisite: `go >= 1.23`, `docker` ideally with [buildx](https://github.com/docker/buildx)
* Start (Up)
<pre><code class="language-bash">make up_db_only
make migrate_db    # Run once
make local_run
</code></pre>
* Finish (Down)
<pre><code class="language-bash">^C
make down_db_only
make clean_db      # if needed
</code></pre>
</details>

## API Endpoints
- **`GET /`**: GraphQL Playground
- **`GET /ping`**: Server healthcheck
- **`GET /graphql`**: GraphQL endpoint

## GraphQL
> Note: Check introspection for better experience
### Schemas: <u>[internal/graphql/schemas](https://github.com/kalalannn/go-graphql_galaxy/tree/main/internal/graphql/schemas)</u>

### Entities
<details>
<summary>Click to expand</summary>
<pre><code class="language-graphql">type Character {
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
</code></pre>
</details>

### Queries
<details>
<summary>Click to expand</summary>
<pre><code class="language-graphql">type Query {
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
</code></pre>
</details>


## Project structure
<details>
<summary> Click to expand </summary>
<pre><code>├── cmd                                # entrypoints (main)
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
</code></pre>
</details>

## Data structure
### DB Table: `Character`, GORM model: <u>[Character](https://github.com/kalalannn/go-graphql_galaxy/blob/main/internal/gorm/entities/character.go)</u>, GraphQL schema: <u>[Character](https://github.com/kalalannn/go-graphql_galaxy/blob/main/internal/graphql/schemas/character.graphql#L1-L13)</u>

<details>
<summary>Click to expand</summary>
<pre><code class="language-sql">    Column        |         Type         | NULLable?
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
</code></pre>
</details>

### DB Table: `Nemesis`, GORM model: <u>[Nemesis](https://github.com/kalalannn/go-graphql_galaxy/blob/main/internal/gorm/entities/nemesis.go)</u>, GraphQL schema: <u>[Nemesis](https://github.com/kalalannn/go-graphql_galaxy/blob/main/internal/graphql/schemas/nemesis.graphql#L1-L7)</u>
<details>
<summary>Click to expand</summary>
<pre><code class="language-sql">    Column    |  Type   | NULLable?
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
</code></pre>
</details>

### DB Table: `Secret`, GORM model: <u>[Secret](https://github.com/kalalannn/go-graphql_galaxy/blob/main/internal/gorm/entities/secret.go)</u>, GraphQL schema: <u>[Secret](https://github.com/kalalannn/go-graphql_galaxy/blob/main/internal/graphql/schemas/secret.graphql#L1-L5)</u>
<details>
<summary>Click to expand</summary>
<pre><code class="language-sql">   Column    |  Type   | NULLable?
-------------+---------+----------
 id          | integer | not null 
 secret_code | bigint  | not null 
 nemesis_id  | integer | not null 
----------------------------------
Foreign keys:
    "nemesis" FOREIGN KEY (nemesis_id) REFERENCES nemesis(id)
</code></pre>
</details>

<!-- ## Examples
### RootQuery
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
}
``` -->
