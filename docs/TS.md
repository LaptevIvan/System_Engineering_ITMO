
## 1. General Information

**Project Name:** Library-service  
**System Type:** CRU gRPC and REST API Service  
**Technology Stack:** Go, PostgreSQL, Docker


## 2. Non-Functional Requirements

### 2.1 Entities: 

#### 2.1.1 Author
    1) id 
    2) name

#### 2.1.2 Book:
    1) id
    2) name
    3) (optional) author_id (id of it authors)
    4) created_at
    5) updated_at

### 2.2 Performance
#### - API response time: < 200 ms
#### - Support for up to 200 rps

## 3. Functional Requirements

### 3.1 The service have to accepts the following requests

------------------------------

#### 3.1.1 Register author

In body of request define name of new author and service will return its id.

##### Constraints for name:

1) name must satisfy the regular expression ^[A-Za-z0-9]+( [A-Za-z0-9]+)*$
2) name's length must be in [1; 512] symbols.


------------------------------

#### 3.1.2 Get author info

Define id of required author and service will return info about him (his name),
if author with given id exists, else return code status 'not found'.

------------------------------

#### 3.1.3 Change author info

Define id of author for updating and his new name, and service will edit author,
if he exists, else return code status 'not found'.

##### New name must satisfy the same constraints that in request of creating author.

------------------------------

#### 3.1.4 Get author's books

Define author's id and service will find all books, which contains
this author in it list of authors.

##### If there is no given author in library, service will return empty list.

------------------------------

#### 3.1.5 Add book

Define name and id of authors of new book and service will return its id.

##### The book may not have authors, but if you specify them, each id of each specified author must be stored in the service.

------------------------------

#### 3.1.6 Get book info

Define id of required book and service will return info about it (name, authors),
if book with given id exists, else return code status 'not found'.

------------------------------

#### 3.1.7 Update book

Define id of book for updating and new info about him,
and service will edit book, if he exists, else return code status 'not found'.


##### The same restrictions apply to the IDs of book authors as in the add book request.

------------------------------

### 4. Configuration file (required environment variables)

#### For gRPC:
1) GRPC_PORT
2) GRPC_GATEWAY_PORT

#### For database
1) POSTGRES_HOST
2) POSTGRES_PORT
3) POSTGRES_DB
4) POSTGRES_USER
5) POSTGRES_PASSWORD
6) POSTGRES_MAX_CONN

