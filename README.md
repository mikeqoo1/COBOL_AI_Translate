# COBOL AI Translate

This project contains a minimal COBOL to Python/Java translation pipeline. It includes:

- A simple parser that builds an abstract syntax tree (AST).
- Conversion of the AST to a small intermediate representation (IR).
- Translators that emit Python or Java source code.
- Example COBOL programs with expected translations in `examples/`.
- Automated tests under `tests/` that verify the parser, IR conversion, and both translators.

## Running Tests

Ensure you have [Go](https://go.dev/) installed. Run all tests with:

```bash
go test ./...
```

The tests parse the sample COBOL program in `examples/` and confirm that the generated Python and Java code matches the expected outputs.

