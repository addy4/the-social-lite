# the-social-lite
Enabling seamless user connections for gaming applications through a tailored backend service for a networking platform.

### Demo
- [RESTful APIs](https://drive.google.com/file/d/11eMSiD0eE_2GMojzqMrpEmgV_y37IOgg/view?usp=sharing)
- [Live WebSocket Notifications For Online Status](https://drive.google.com/file/d/1l8xuHOPP9v_aBJ4fNLtcePOkp6Egs2GJ/view?usp=sharing)

## How To Run
- Development/Local
  - docker run -p 6379:6379 redis
  - go mod download
  - go build services/main.go

- Production/Docker
  - docker compose up

## Some design insights! 
- Creating RESTful APIs with well defined structs as per the feature requirement. 
- Using redis as a database.
  - Light weight
  - Low latency
  - In memory
- WebSocket server to which the user can connect.
- Add basic Authentication to websocket server, this lets us know which username has come online. 
- Notification to websocket clients for new connections.

## Features
- Create __*Users*__
- Add __*Friend*__ To User's Network
- Create a __*Party*__ (Short sessions as group of Users)
- Add __*Friend*__ to __*Party*__ as __*Member*__

## Endpoints

### 1. Create User

- **Endpoint:** `/api/register`
- **Method:** `POST`
- **Input:**
  ```json
  {
      "register": {
          "username": "example_user",
          "password": "secure_password"
      }
  }

### 1. Add Friend
Current User Follows Username

- **Endpoint:** `/api/friend`
- **Method:** `POST`
- **Input:**
  ```json
  {
    "follow": {
        "username": "to_bo_followed",
        "currentuser": "follower"
    }
}

### 3. Get Friends Of Users

- **Endpoint:** `/api/friends`
- **Method:** `GET`
- **Input:**
  ```json
  {
      "getfriends": {
          "currentuser": "logged_in_user"
      }
  }

### 4. Add Party

- **Endpoint:** `/api/party`
- **Method:** `POST`
- **Input:**
  ```json
  {
      "createparty": {
          "partytitle": "playo"
      }
  }

### 5. Add Members of Party

- **Endpoint:** `/api/party/membership`
- **Method:** `POST`
- **Input:**
  ```json
  {
    "addtoparty": {
        "partytitle": "playo",
        "username": "any_registered_user"
    }
}

### 6. Get Members of Party

- **Endpoint:** `/api/party/members`
- **Method:** `GET`
- **Input:**
  ```json
  {
    "getpartymembers": {
        "partytitle": "playo"
    }
  }
