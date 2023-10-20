# Idiomatic Clean Architecture with Golang

Have you ever read a blog post about an architecture, gotten excited, attempted to implement it, only to find it too abstract or disconnected from real-world scenarios? Often, the problems showcased are oversimplified, leaving readers without a comprehensive understanding.

My aim with this post is to present a more holistic view of an architecture suitable for both straightforward and complex projects.

TLWR: Here is the [GitHub repository][repo]

## Objectives

Before delving into the design, it's essential to set the architectural objectives:

- Scalability in terms of codebase and the number of teams collaborating.
- Not a one-size-fits-all solution.

## Case Study: A Social Media App

To make things tangible, we'll use a rudimentary social media application, similar to  ex Twitter (pun intented).

- User registration and login.
- Posting messages.
- Following other users.
- Liking and commenting on posts.
- Viewing a feed of posts from followed users.

The focus will be on the API's architecture. For the sake of clarity, some components like logging, configuration management, and integration tests are omitted.

## Anatomy of a Package

While individual packages may vary, most follow this general structure:

```markdown
📂package
 ┣ 📜entity.go
 ┣ 📜handler.go
 ┣ 📜repo.go
 ┗ 📜service.go
```

### handler

Handles HTTP requests by parsing the input, invoking the service, and dispatching the response.

### service

Houses the business logic. It liaises with the repository and processes data using entities.

### entity

Defines the data structure for the given problem - be it a user, a post, or a comment. It could be a domain or aggregate from Domain Driven Design (DDD) or a simple Data Transfer Object (DTO).

### repository

## Fundamental Reasoning Behind the Structure

Golang is idiomatic, making it imperative to structure the code in sync with its idioms.

The crux of our architecture revolves around Go's encapsulation method. Instead of classes found in many languages, Go employs packages to encapsulate both data and behavior. **This means we structure packages around the problems they address, not the component types they contain.**

For example, we'll have a 'user' package handling user-related use cases, rather than separate packages for user entities and repositories. Both the entity and repository play roles in the user use case.

Have a look at the project structure:

## Project Structure

```markdown
📦morgan
 ┣ 📂cmd
 ┃ ┗ 📜main.go
 ┣ 📂config
 ┃ ┗ 📜config.go
 ┣ 📂internal
 ┃ ┣ 📂auth
 ┃ ┃ ┣ 📜handler.go
 ┃ ┃ ┣ 📜middleware.go
 ┃ ┃ ┣ 📜service.go
 ┃ ┃ ┗ 📜token.go
 ┃ ┣ 📂feed
 ┃ ┃ ┣ 📜feed.go
 ┃ ┃ ┣ 📜handler.go
 ┃ ┃ ┗ 📜service.go
 ┃ ┣ 📂follow
 ┃ ┃ ┣ 📜follow.go
 ┃ ┃ ┣ 📜handler.go
 ┃ ┃ ┣ 📜repo.go
 ┃ ┃ ┗ 📜service.go
 ┃ ┣ 📂post
 ┃ ┃ ┣ 📂comment
 ┃ ┃ ┃ ┣ 📜comment.go
 ┃ ┃ ┃ ┗ 📜repo.go
 ┃ ┃ ┣ 📂like
 ┃ ┃ ┃ ┣ 📜like.go
 ┃ ┃ ┃ ┗ 📜repo.go
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

Engages with the database to persist or retrieve entities.

## Dependency Rule

Each component depends on interfaces it defines. Dependencies are injected into the component. The order of the dependencies is as follows: Handler -> Service -> Repository.

## Cross-Domain Dependencies

We identify two categories of cross-domain dependencies:

1. Dependencies between components within the same bounded context.
2. Dependencies between components across different bounded contexts.

The concept of a bounded context originates from DDD. It signifies a perimeter around components sharing a common domain. For instance, 'post', 'comment', and 'like' can be seen as sharing a bounded context due to their strong coupling. If one changes, the others will likely change as well.

In the 'post' package, 'like' and 'comment' entities are accessed via repositories since they share the context. These sub-packages don't need their respective services since they are only used within the 'post' package.

Conversely, in the 'follow' package, the 'user' domain is accessed using the 'user' service, maintaining domain boundaries without direct repository imports.

---

## Conclusion

The architecture we've explored offers both simplicity and flexibility, making it an ideal choice for many Golang projects. It is definitely not a perfect solution, but it is a good starting point. As your project grows, you can adapt it to suit your needs.

Here is the GitHub repository for the project: [morgan-go][repo]

[repo]: https://github.com/ahmetildirim/morgan
