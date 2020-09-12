# **Running MovieApi**

- Open terminal.
- Use `Makefile` for
    1. Set GOPATH .
    2. Install all dependencies.
    3. Build API.
    4. Run API.

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
make go-build
```

#### 4. **Run API**
```
go-run
```

# **MovieAPI ENDPOINTS**

####    1. Ping Check
Checks if API is up and running.
```
/
```
####    2. Create Movie
Movie data provided in body in Json format.
```
/movie
```
####    3. Get Movie by movie name
Movie data for particular movie is returned if it exists in database.
```
/movie/{movie-name}
```
####    4. Get all Movies
Data of all the movies present in the database is returned
```
/movies
```
####    5. Delete Movie by name
Data of movie with given name is deleted from database if it exists.
```
/movie/{movie-name}
```
####    6. Update movie details by movie name
Data of movie will be updated as provided in the body with the request if it exists.
```
/movie/{movie-name}
```
