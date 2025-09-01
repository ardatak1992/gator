````markdown
# 📰 Blog Aggregator App

This is a CLI-based **blog aggregator** built in **Go** with **PostgreSQL** for persistent storage. It allows users to register, follow RSS feeds, and browse aggregated blog content from feeds they follow.

---

## 🛠 Prerequisites

Before running the app, make sure you have the following installed:

### ✅ Go

Install the latest version of Go:  
🔗 https://golang.org/dl/

```bash
go version
```
````

### ✅ PostgreSQL

Install PostgreSQL:

- **macOS (Homebrew)**:

  ```bash
  brew install postgresql
  brew services start postgresql
  ```

- **Ubuntu/Debian**:

  ```bash
  sudo apt update
  sudo apt install postgresql postgresql-contrib
  sudo service postgresql start
  ```

- **Windows**:
  Download from: [https://www.postgresql.org/download/windows/](https://www.postgresql.org/download/windows/)

Create the database and user:

```bash
createdb blog_aggregator
createuser blog_user --pwprompt
```

Update your app’s DB connection string accordingly.

### ✅ Gator

If you're using policies with [`gator`](https://github.com/open-policy-agent/gator):

```bash
go install github.com/open-policy-agent/gator@latest
export PATH="$PATH:$(go env GOPATH)/bin"
gator version
```

---

## 🚀 Building and Running the App

Build by using:

```bash
go build
```

Building Run any command using:
go build

```bash
./gator [command]
```

Example:

```bash
./gator register username
```

---

## ⚙️ CLI Commands

| Command     | Description                                                | Usage                      |
| ----------- | ---------------------------------------------------------- | -------------------------- |
| `register`  | Register a new user account.                               | register \<username>       |
| `login`     | Log in to an existing user account.                        | login \<username>          |
| `reset`     | Reset (delete) the user table. **⚠️ Destructive**          | reset                      |
| `users`     | List all registered users.                                 | users                      |
| `agg`       | Trigger the feed aggregation (fetch latest blog posts).    | agg <time_between_reqs>    |
| `addfeed`   | Add a new RSS feed (requires login).                       | addfeed \<feedName> \<url> |
| `feeds`     | List all available feeds.                                  | feeds                      |
| `follow`    | Follow a feed (requires login).                            | follow \<feedUrl>          |
| `following` | View feeds you are following (requires login).             | following                  |
| `unfollow`  | Unfollow a feed (requires login).                          | unfollow \<feedUrl>        |
| `browse`    | Browse posts from feeds you're following (requires login). | browse <limit=2>           |

---

## 🧰 Configuration File

The CLI stores session/auth info in a config file in your home directory.

### 📌 Location

```bash
~/.gatorconfig.json
```

This is automatically created after you log in or register.

### 🧾 Example `.gatorconfig.json`

```json
{
  "db_url": "conntection_string",
  "current_user_name": "exampleuser"
}
```

### 🛠 Managing the Config

- To **view** it:

  ```bash
  cat ~/.gatorconfig.json
  ```

- To **delete/reset** it:

  ```bash
  rm ~/.gatorconfig.json
  go run main.go login
  ```

- To **secure** it:

  ```bash
  chmod 600 ~/.gatorconfig.json
  ```
