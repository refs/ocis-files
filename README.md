## ocis-files

This example showcases how a main module could "bundle" a few go modules, with the intention of provide a sort of "umbrella" project.

## Structure

```
.
├── go.mod
├── ocis-ocs
    ├── go.mod
│   └── pkg
│       └── server
│           └── server.go
├── ocis-phoenix
│   ├── go.mod
│   └── pkg
│       └── server
│           └── server.go
└── ocis-webdav
    ├── go.mod
    └── pkg
        └── server
            └── server.go
```
