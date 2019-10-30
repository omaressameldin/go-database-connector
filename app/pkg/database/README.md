# Database Package
This package contains the database interface that should be used with all database related tasks in services using the connector

## Structure
The file structure is as follows:
- [database_test.go](./database_test.go): contains all tests related to package
- [mock.go](./mock.go): contains the Mocked connector that should be used when you need to mock the connector for integration tests.
- [types.go](./types.go): this contains all needed structs and interfaces
  - structs:
    - `ValidationError`: contains `Field` for field that has error and `Message` for the error message related to that field
    - `Validator`: contains 'Field' for field that needs to be validated and `Error` for the field error if any
    - `Updated`: is used when the databse needs to update an object. Contains `Key` for field name and `Val` for field value
  - interfaces:
    - `Connector`: any object used in the project that interacts with a database should implement that interface.
  - [functions.go](./functions.go): contains all helper functions needed when implementing `Connector` interface such as creating validators (`CreateValidator`), valitonErors(`CombineValidationErrors`) and json object from validation errors (`GenerateJsonError`)