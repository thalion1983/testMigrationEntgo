# Demo for migrations
This demo is based en on demo https://github.com/rotemtam/ent-versioned-migrations-demo

Schemas will be copied from there.

Tag: v0.0.1: First migration using atlas

Executed over a clean database
1. Removed the sentence that creates the schema in server.go
2. Created a migrate diff and stored in dir migrate/migrations
3. Applied such migrate diff
