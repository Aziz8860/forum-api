# Forum API

### Packages
- Gin
- Go-MySQL-Driver
- Viper (for manage configs)

### Includes
- Middleware
- JWT (access token and refresh token)

### Endpoints
1. /memberships/sign-up
2. /memberships/login
3. /memberships/refresh
4. /posts/create
5. /posts/comment/:postID
6. /posts/user_activity/:postID
7. /posts/?pageIndex=X&pageSize=X
8. /posts/:postID

### ERD
<img src="https://github.com/user-attachments/assets/c9d97a20-1c6b-43f4-8d19-3f6f6e2b364c" alt="ERD" width="800">

### Get Started
1. Run docker

   ```bash
   docker-compose up -d
   ```

2. Run main

   ```bash
    go run cmd/main.go
   ```
