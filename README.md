
# Test


```go
func DeleteUser(id int, user User, body DeleteUserBody, userRepository UserRepository, r golang.Request) (golang.Response, error) {
  err := userRepository.Delete(id)
  if err != nil {
    return err // Throw 500 (with error message if enabled)
  }

  return golang.Success("User successfully deleted")
}

func AuthMiddleware(r goland.Request, auth AuthService, container Container) (goland.Request, error) {
  authHeader := r.Headers().Get("Authorization")

  user := auth.Authorize(authHeader)
  if user != nil {
    container.Provide(di.T[User], user)
    return r
  }

  return nil, MyError
}

func main() {
  userRouter := golang.NewRouter().
    Middleware(AuthMiddleware).
    Get("/", GetUsers).
    Post("/", CreateUser).
    Delete("/{id}", DeleteUser).
    Route(nil, "/", PatchUser)

  app := golang.NewRouter().
    Service(di.T[UserRepository], NewUserRepository)
    Service(di.T[AuthService], NewAuthService)
    Router("/users", userRouter)
}
```

- Path patterns
- Routing system (with subrouting)
- Autowiring
  - Services
  - Request extraction (query params, body, etc)
- User friendly response system
- User friendly request system
- Middleware system
- Error handling system
