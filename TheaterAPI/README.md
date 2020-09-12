# **Running TheaterAPI**

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
make go-run
```

# **API Endpoints**

### 1. Ping

Check API status.

```
/
```

### 2. Get Theater

Get theater by name.

```
/theater/{name}
```

### 3. Get all Theater

Get all theaters in the database.

```
/theaters
```

### 4. Create Theater

Create new theater entry in database from the details provided in body.

```
/theater
```

### 5. Update Theater details

Update existing theater details.

```
/theater/{name}
```

### 6. Delete Theater

Delete theater from database if exists.

```
/theater/{name}
```

