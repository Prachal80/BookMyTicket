# **Running ShowAPI**

-Explaination:
-ShowAPI uses two collections.
    1.One stores Show information:ShowID,TheatreID,MovieID
    2.Second Collection stores which user is stored in which show.

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

# **API Endpoints**

### 1. Ping

Check API status.

```
/bookings
```

### 2. Get Show

Get show by name.

```
/show/{id}
```

### 3. Get all Theater

Get all shows in the database.

```
/shows
```

### 4. Create Show

Create new Show entry in database from the details provided in body.

```
/show
```

### 5. Update Show details

Update existing show details.

```
/show/{id}
```

### 6. Delete Show

Delete show from database if exists.

```
/show/{id}
```
### 7. Get Booking of a Show

Get show booking by name.

```
/book/{id}
```

### 8. Get all Show Booking

Get all show bookings in the database.

```
/bookings
```
### 9. Create Show Booking

Create new Show Booking entry in database from the details provided in body.

```
/createbook/{id}
```
### 10. Add User to Show Booking

Create new Show Booking entry in database from the details provided in body.

```
/book/{id}
```
