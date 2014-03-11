# Name My Robot

Inspired by [namemydaughter.com](http://www.namemydaughter.com/), Team 1759 is
reaching out to the internet to generate ideas for the name of our robot. 

## Getting Started

To run the local server, you'll need to set some environment variables.

* `DB_DSN` - The destination of the database. For SQLite in development mode,
  this'll probably be a path to the database on your filesystem. Example:
  `/tmp/my.db`. If it doesn't already exist, it'll be created when ran. In
  production, this should be a Postgres URL. Example:
  `postgres://username:password@host:port/db`

After you have that set, install dependencies.

    $ go get

And run the server:

    $ go run server.go

## API Docs

### GET /api/adjectives

#### Action

Gets all existing adjectives.

#### Parameters

None.

#### Response

JSON array of all adjectives.

Example:

```
[
  {
    "id": 1,
    "createdAt": "2014-03-11T09: 14: 52.16135225Z",
    "content": "Fluffy",
    "votes": 0
  },
  {
    "id": 2,
    "createdAt": "2014-03-11T09: 14: 53.766193996Z",
    "content": "Dogelike",
    "votes": 2
  },
  {
    "id": 3,
    "createdAt": "2014-03-11T09: 14: 53.942152545Z",
    "content": "Awesome",
    "votes": 0
  }
]
```

### POST /api/adjectives

#### Action

Creates a new adjective.

#### Parameters

JSON of new adjective.

Example:

```
{
  "content": "Fluffy"
}
```

#### Response

JSON of new adjective as stored in the database.

Example:

```
{
  "id": 1,
  "createdAt": "2014-03-11T09: 14: 52.16135225Z",
  "content": "Fluffy",
  "votes": 0
},
```

### PUT /api/adjectives/:id/vote

#### Action

Votes for an adjective

#### Parameters

In URL:

* `:id` - ID of the adjective to vote on.

In request:

None.

#### Response

JSON of adjective voted for.

Example:

```
{
  "id": 1,
  "createdAt": "2014-03-11T09: 14: 52.16135225Z",
  "content": "Fluffy",
  "votes": 1
},
```

### GET /api/names

#### Action

Gets all existing names.

#### Parameters

None.

#### Response

JSON array of all names.

Example:

```
[
  {
    "id": 1,
    "createdAt": "2014-03-11T09: 14: 52.16135225Z",
    "content": "Spud",
    "votes": 0
  },
  {
    "id": 2,
    "createdAt": "2014-03-11T09: 14: 53.766193996Z",
    "content": "Jimmy",
    "votes": 2
  },
  {
    "id": 3,
    "createdAt": "2014-03-11T09: 14: 53.942152545Z",
    "content": "LabVIEW",
    "votes": 0
  }
]
```

### POST /api/names

#### Action

Creates a new name.

#### Parameters

JSON of new name.

Example:

```
{
  "content": "LabVIEW"
}
```

#### Response

JSON of new name as stored in the database.

Example:

```
{
  "id": 1,
  "createdAt": "2014-03-11T09: 14: 52.16135225Z",
  "content": "LabVIEW",
  "votes": 0
},
```

### PUT /api/names/:id/vote

#### Action

Votes for an name

#### Parameters

In URL:

* `:id` - ID of the name to vote on.

In request:

None.

#### Response

JSON of name voted for.

Example:

```
{
  "id": 1,
  "createdAt": "2014-03-11T09: 14: 52.16135225Z",
  "content": "LabVIEW",
  "votes": 1
},
```

## License

See [LICENSE](LICENSE).
