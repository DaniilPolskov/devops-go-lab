# DevOps Go Lab

A simple Go web application built to practice and demonstrate DevOps concepts such as Docker, CI/CD, cloud infrastructure, monitoring, and environment management.

---

##  Features

- Basic HTTP server with multiple endpoints:
  - `/` – Request counter
  - `/health` – Health check
  - `/time` – Current server time
  - `/env` – Displays environment variables
  - `/metrics` – Prometheus metrics endpoint
- Dockerized with `Dockerfile` and `docker-compose.yml`
- Prometheus-compatible metrics (`/metrics` endpoint)
- Unit testing with `go test`
- CI/CD pipelines (GitLab CI/CD and Jenkins)
- Infrastructure as Code using Terraform
- Ready for deployment to AWS (EC2)

---

## Project Structure

```

devops-go-lab/
├── cmd/
│   └── server/             # Main Go application
├── configs/                # Environment configuration
│   └── .env                # Example environment file
├── infra/                  # Terraform files (EC2, VPC, etc.)
├── Jenkinsfile             # Jenkins pipeline definition
├── Dockerfile              # App Docker image build
├── docker-compose.yml      # Local orchestration
├── go.mod / go.sum         # Go modules
└── README.md               # You are here!

```

---

## How to Run Locally

### 1. Clone the Repository

```bash
git clone https://github.com/DaniilPolskov/devops-go-lab.git
cd devops-go-lab
```

### 2. Configure Environment

Create a `.env` file in `configs/`:

```env
APP_PORT=8080
APP_ENV=dev
```

### 3. Run with Docker Compose

```bash
docker-compose --env-file configs/.env up --build
```

Then visit: [http://localhost:8080](http://localhost:8080)

---

## API Endpoints

| Path       | Description                   |
| ---------- | ----------------------------- |
| `/`        | Root route, counts requests   |
| `/health`  | Health check                  |
| `/time`    | Current server time           |
| `/env`     | Shows env variables           |
| `/metrics` | Prometheus metrics exposition |

---

## CI/CD Pipelines

### GitLab CI/CD

Located in `.gitlab-ci.yml` and includes stages:

* Build with `go build`
* Unit testing with `go test`
* Docker build

### Jenkins

Pipeline defined in `Jenkinsfile`:

```groovy
pipeline {
  agent any

  stages {
    stage('Build') {
      steps {
        sh 'go build -o app ./cmd/server'
      }
    }
    stage('Test') {
      steps {
        sh 'go test ./...'
      }
    }
    stage('Docker Build') {
      steps {
        sh 'docker build -t daniilpolskov/devops-go-lab .'
      }
    }
  }
}
```

---

## Infrastructure as Code (Terraform)

Terraform files located in `infra/`:

* AWS provider configured
* EC2 instance creation
* Outputs include instance ID and public IP

### Example Terraform Commands:

```bash
cd infra/
terraform init
terraform apply
```

---

## Monitoring

* Exposes Prometheus metrics at `/metrics`
* Counts total HTTP requests
* Can be integrated with Grafana for dashboarding

---

## Technologies Used

* **Go** 1.22+
* **Docker** & Docker Compose
* **Prometheus** (metrics)
* **Jenkins** (CI/CD)
* **GitLab CI/CD**
* **Terraform** (AWS EC2 provisioning)

---

## Future Improvements

* Add database (PostgreSQL, SQLite)
* systemd support for local deployment
* NGINX reverse proxy + Certbot SSL
* Auto-deploy to AWS (via Ansible/Terraform)
* Frontend with React or Vue
* More Prometheus metrics (status codes, request durations)

---

## Links

* GitHub: [https://github.com/DaniilPolskov/devops-go-lab](https://github.com/DaniilPolskov/devops-go-lab)
* GitLab: [https://gitlab.com/daniilpolskov-group/devops-go-lab](https://gitlab.com/daniilpolskov-group/devops-go-lab)
