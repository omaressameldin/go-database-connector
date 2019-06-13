# Database Package
This package contains the firebase package that implements [database interface](../database) and [uploader interface](https://github.com/omaressameldin/lazy-panda-upload-service/blob/master/core/pkg/uploader/uploader.go).

## Structure
The file structure is as follows:
- [firebase.go](./firebase.go): contains functions to start (`StartConnection`) and close (`closeConnection`) the connection to firebase.
- [adapter.go](./adapter.go): contains all functions required for firebase to implement [database interface](../database).
- [uploader.go](./uploader.go): contains all functions required for firebase to implement [uploader interface](https://github.com/omaressameldin/lazy-panda-upload-service/blob/master/core/pkg/uploader/uploader.go).
- [utils.go](./utils.go): contains helper functions used by [adapter](./adapter.go) and [uploader](./uploader.go) packages