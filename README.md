
# Environment Example in GoLang

>Example of how to implement a Go application with two environments.

When we are developing real world applications, we often use more than one environment, a production and development one, for example.

This is an approach of how to implement two environments using go, in a simple way.

Environment folder structure:


    ├── config
    │   └── config.go
    │   └── development
    │   │   └── database.go
    │   └── production
    │   │   └── database.go


each **database.go** has it's onw implementation of the function `GetSource()`.

**config.go**  defines a list of databaseSources, that uses the environment name as map key, and call it's onw `GetSource` function.


------------



To make it work, you need to create a localhost postgres database named `database_test` with user `juniorbarros` and **without password**.

