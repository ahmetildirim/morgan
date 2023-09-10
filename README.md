# morgan
Morgan is a social media focused on news and discussions

[Project board](https://github.com/users/ahmetildirim/projects/1)

[Repository](https://github.com/ahmetildirim/morgan)


## Architecture
    project/
    ├── cmd/
    ├── internal/
    │   ├── user/
    │   ├── post/
    │   ├── feed/
    │   ├── comment/
    │   ├── reaction/
    │   ├── platform/

Packages are structured based on the domain driven design. Each package has its own layers. For example, user package has 
user entity, repository, service, handler, and domain layers.


