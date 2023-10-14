# Clean Architecture with Golang

Ever see a blog post about an architecture, get excited, sit down to implement it, and then realize it detached from reality? In many case, problems they solve are so simple and they fail to draw whole picture.

I hope to address this issue by providing a more complete picture of the architecture that can be used in in both simple and complex projects.

## Objectives

We need to define the objectives of the architecture before we can start designing it:

- Scalable in terms of code and number of teams working on the project
- Not a one size fits all solution

## Case Study

We will be using a social media application as a case study. It will be similar to Twitter, but with basic use cases:

- Users can register and login
- Users can post messages
- Users can follow other users
- Users can view a feed of messages from users they follow
- Users can like and comment on messages

We will implement on the API and only focus on the architecture of the API.
Many necessary components will be omitted for simplicity, like logging, config management, integration tests, etc.

## Fundamental Reasoning

Go is idiomatic, so it is better to structure the code in a way that complies with the language's idioms.

What affects our architecture is how encapsulation is done in Go. Unlike classes in other languages, we use packages to encapsulate data and behavior. **That means packages will be created based on the problem they solve, not based on the type of components they contain.**

For example, we will have a package for the user use case, not a package for the user entity and another for the user repository. This is because the user entity and repository are both part of the user use case.

## Project Structure

``` plain
📦morgan
 ┣ 📂cmd
 ┣ 📂config
 ┣ 📂internal
 ┃ ┣ 📂auth
 ┃ ┃ ┣ 📜handler.go
 ┃ ┃ ┣ 📜middleware.go
 ┃ ┃ ┣ 📜service.go
 ┃ ┃ ┗ 📜token.go
 ┃ ┣ 📂comment
 ┃ ┃ ┣ 📜comment.go
 ┃ ┃ ┣ 📜handler.go
 ┃ ┃ ┣ 📜repo.go
 ┃ ┃ ┗ 📜service.go
 ┃ ┣ 📂feed
 ┃ ┃ ┣ 📜feed.go
 ┃ ┃ ┣ 📜handler.go
 ┃ ┃ ┗ 📜service.go
 ┃ ┣ 📂follow
 ┃ ┃ ┣ 📜follow.go
 ┃ ┃ ┣ 📜handler.go
 ┃ ┃ ┣ 📜repo.go
 ┃ ┃ ┗ 📜service.go
 ┃ ┣ 📂like
 ┃ ┃ ┣ 📜like.go
 ┃ ┃ ┣ 📜repo.go
 ┃ ┃ ┗ 📜service.go
 ┃ ┣ 📂post
 ┃ ┃ ┣ 📜handler.go
 ┃ ┃ ┣ 📜post.go
 ┃ ┃ ┣ 📜repo.go
 ┃ ┃ ┗ 📜service.go
 ┃ ┗ 📂user
 ┃ ┃ ┣ 📜handler.go
 ┃ ┃ ┣ 📜repo.go
 ┃ ┃ ┣ 📜service.go
 ┃ ┃ ┗ 📜user.go
 ┣ 📂migrations
 ┣ 📂scripts
 ┣ 📂test
```

## Package Anatomy

Most packages will have the following structure, but some may have more or less files depending on the use case:

``` plain
📂package
 ┣ 📜entity.go
 ┣ 📜handler.go
 ┣ 📜repo.go
 ┗ 📜service.go
```

### Handler

The handler is the component that handles the HTTP request. It is responsible for parsing the request, calling the service, and returning the response.

### Service

The service is the component that contains the application logic. It is responsible for interacting with the repository and performing the necessary business logic using the entity.

### Entity

The entity is the data structure that represents the problem we are trying to solve. It can be a user, a post, a comment, etc.

It can be a domain or an aggregate from Domain Driven Design (DDD), but it doesn't have to be. It can be a simple data transfer object (DTO).

### Repository

Interacts with the database. It is responsible for persisting and retrieving entities.

## Dependency Rule

Each component depends on interfaces it defines. Dependencies are injected into the component. The order of the dependencies is as follows: Handler -> Service -> Repository.

## Cross-Domain Dependencies

Cross-domain dependencies are not done over services. So, if the post service needs the likes domain, it will import the likes service.

Check the following example from post service:

``` go
    func (s *Service) AddLike(ctx context.Context, postID, userID uuid.UUID) error {
        exists, err := s.repo.Exists(ctx, postID)
        if err != nil {
            return err
        }

        err = s.likeService.Create(ctx, postID, userID)
        if err != nil {
            return err
        }

        if !exists {
            return ErrPostNotFound
        }

        err = s.repo.AddLike(ctx, postID)
        if err != nil {
            return err
        }

        return nil
    }
```

The **AddLike** service imports the **likeService** and uses it to create a like. This is a cross-domain dependency. The **likeService** interface is defined in the **post** package. The **likeService** implementation is defined in the **like** package.

