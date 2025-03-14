# Shiftr

A simple HTTP client and server which, periodically, will obtain payloads from multiple HTTP targets, save, and serve
them locally from a static file server.

Supports Basic Authentication with username and password. 

Custom request headers can also be specified on a per URL basis.

## Configuration
The application looks for a `config.yml` file is the local directory.

Example `config.yml` file

```yaml
---
port: "8080"
refresh_interval : 6
data_folder: "data"
sources:
  - name: "Test payload 1"
    url: "https://microsoftedge.github.io/Demos/json-dummy-data/512KB.json"
    server_filename: "data1.json"
    headers:
      accept: "application/json"
    basic_auth:
      username: "myuser"
      password: "mypassword"

  - name: "Test payload 2"
    url: "https://microsoftedge.github.io/Demos/json-dummy-data/64KB-min.json"
    server_filename: "data2.json"
    headers:
      accept: "application/json"
    basic_auth:
      username: "myuser"
      password: "mypassword"

```

In the above configuration example, the files will be saved in a `data` folder and served at these endpoints:

```shell
http://localhost:8080/data1.json
http://localhost:8080/data2.json
```