# **Running Mail-API**

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

# **API ENDPOINTS**

####    1. Ping Check
Checks if API is up and running.
```
/
```
####    2. Add User booking details
User data provided in body in Json format.
```
/person
```
####    3. Get User by Email
User data for give email is returned if exist in database.
```
/person/{email}
```
####    4. Get all Users
Data of all the users in the database is returned
```
/person
```
####    5. Delete User by Email
Data of user with given email is deleted from database if it existed.
```
/person/{email}
```
####    6. Send Email to User
User with given Email must exist.
```
/person/{email}
```
####    7. Send SMS to User
User with given Email must exist.
```
/person/{email},{mobile_number}
```
