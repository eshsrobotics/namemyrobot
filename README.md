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

## License

See [LICENSE](LICENSE).
