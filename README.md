# DevOps Go Lab  

A simple Go web application designed for learning and practicing DevOps concepts such as Docker, CI/CD, and cloud deployment.  

---

## Features  

- Basic HTTP server with endpoints:  
  - `/` – Request counter  
  - `/health` – Health check  
  - `/time` – Current server time  
  - `/env` – View environment variables  
- Dockerized with Dockerfile and docker-compose  
- Automated tests with `go test`  
- CI/CD using GitLab Pipelines  
- Ready for deployment to cloud platforms (AWS, Azure, etc.)  

---

## How to Run Locally  

### 1. Clone the repository  

```bash  
git clone https://github.com/DaniilPolskov/devops-go-lab.git  
cd devops-go-lab  
```  

### 2. Configure environment  

Create a `.env` file in the `configs/` folder:  

```  
APP_PORT=8080  
APP_ENV=dev  
```  

### 3. Run with Docker Compose  

```bash  
docker-compose --env-file configs/.env up --build  
```  

The application will be available at: [http://localhost:8080](http://localhost:8080)  

---

## CI/CD and DevOps Setup  

This project uses GitLab CI/CD to automate:  

- Build with `go build`  
- Unit testing with `go test`  
- Docker image build  
- (Optional) Deployment with systemd or Docker  

The `.gitlab-ci.yml` file defines stages for build and test and can be extended for deployment.  

[View CI/CD Pipeline on GitLab](https://gitlab.com/daniilpolskov-group/devops-go-lab)  

---

## Technologies Used  

- Go 1.22.2  
- Docker + Docker Compose  
- GitLab CI/CD  
- GitHub & GitLab integration  

---

## Future Improvements  

- Add database support (PostgreSQL, SQLite)  
- Add frontend with React or Vue  
- Automatic deployment to cloud (AWS EC2, Azure App Service)  
- Monitoring via Prometheus + Grafana
- systemd (for local service deployment)  
- NGINX (optional reverse proxy)  
- Certbot (SSL automation)  

## Links

* GitHub Repo: [https://github.com/DaniilPolskov/devops-go-lab](https://github.com/DaniilPolskov/devops-go-lab)
* GitLab CI/CD: [https://gitlab.com/daniilpolskov-group/devops-go-lab](https://gitlab.com/daniilpolskov-group/devops-go-lab)
