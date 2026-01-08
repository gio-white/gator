# Gator 🐊

Gator is a high-performance CLI RSS feed aggregator built in Go. It allows users to register, follow their favorite RSS feeds, and run a background worker to collect and store posts in a PostgreSQL database for offline browsing.

---

## 🛠 Prerequisites

To run Gator, you need the following installed:

* **Go**: [Installation Guide](https://go.dev/doc/install) (v1.21+)
* **PostgreSQL**: [Installation Guide](https://www.postgresql.org/download/) (v15+)
* **Goose**: Used for database migrations. Install it via:
  ```bash
  go install [github.com/pressly/goose/v3/cmd/goose@latest](https://github.com/pressly/goose/v3/cmd/goose@latest)