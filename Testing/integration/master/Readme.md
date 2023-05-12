# Golang Integration Test Example

This is a small example project demonstrating one strategy of separating unit and integration tests.

To run this example, first start the web server:

```bash
go run ./addserver/main.go
```

Then in a separate shell, you can run the following:

```bash
go test # Run all unit tests in our package
go test -tags=integration # Run all unit and integration tests in our package
go test -tags=integration -run ^TestAPISum$ # Run a specific integration test
```

## Notes for running with VSCode

To debug a specific integration test in VSCode:

1. Make sure VSCode is opened with this directory as the root (**golang_integration_testing**)
2. In **add_integration_test.go**, add a breakpoint somewhere in `TestAPISum`
3. Highlight `TestAPISum`
4. Run the **Debug Golang test** profile from the **Run** menu
