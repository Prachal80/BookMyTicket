# USERAPI
UserAPI is used to register and deregister a user using a unique Email address.

## **Running USERAPI**

- Use `Makefile` for
     * Set GOPATH .
     * Download and Install Dependencies.
     * Build.
     * Run.

#### 1. **Set GOPATH**
```
make set-path
```

#### 2. **Install all dependencies**
```
make go-get
```

#### 3. **Build API**
```
make build
```

#### 4. **Run API**
```
make run
```

## **API ENDPOINTS**

####    1. Ping Check
Checks if API is up and running.
```
/
```
####    2. Create User
User data provided in body in Json format.
```
/user
```
####    3. Get User by EmailID
User data for give email is returned if exist in database.
```
/user/{email}
```
####    4. Get all Users
Data of all the users in the database is returned
```
/user
```
####    5. Delete User by EmailID
Data of user with given email is deleted from database if it existed.
```
/user/{email}
```
####    6. Update user by EmailID
Data of user is updated as provided in the body with the request if it exist.
```
/user/{email}
```
