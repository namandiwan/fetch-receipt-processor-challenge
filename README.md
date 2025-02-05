# 🧾 Fetch Receipt Processor API

## 📌 Project Overview
The **Fetch Receipt Processor API** is a Go-based backend service that processes receipts and awards points based on predefined rules. The API allows users to:
- **Submit a receipt** (`POST /receipts/process`)
- **Retrieve points for a receipt** (`GET /receipts/{id}/points`)

This project is built using **Go 1.23.6** and supports **Docker** for containerized deployment.

---

# 🚀 **Installation and Setup**

## **1️⃣ Prerequisites**
Before running the project, ensure you have the following installed:

| Dependency      | Version                     | Installation Command |
|----------------|----------------------------|----------------------|
| **Go**        | `1.23.6`                     | `brew install go` (Mac) |
| **Docker**    | `27.4.0 (Docker Desktop 4.37.2)` | [Download & Install](https://www.docker.com/products/docker-desktop) |
| **jq (optional, for testing)** | Latest | `brew install jq` |

Verify installations:
```sh
go version      # Should output go1.23.6
docker --version # Should output Docker version 27.4.0
jq --version    # If installed, should show version info
```

## 2️⃣ Clone the Repository
```sh
git clone https://github.com/namandiwan/fetch-receipt-processor-challenge.git
cd fetch-receipt-processor-challenge
```

## 3️⃣ Install Dependencies
```sh
go mod tidy
```

# 🚀 Running the Project

## 1️⃣ Run Locally with Go
```sh
go run main.go
```
The server will start on http://localhost:8080

## 2️⃣ Run with Docker
📌 Build the Docker Image
```sh
docker build -t receipt-processor .
```

📌 Run the Docker Container
```sh
docker run -p 8080:8080 receipt-processor
```
💡 If port 8080 is in use, run:
```sh
docker run -p 9090:8080 receipt-processor
```
Then access the API at http://localhost:9090

# 📡 API Endpoints and Usage

📌 Submit a Receipt
```sh
curl -X POST "http://localhost:8080/receipts/process" \
     -H "Content-Type: application/json" \
     --data-binary "@examples/morning-receipt.json"
```

✅ Expected Response:
```sh
{"id": "adb6b560-0eef-42bc-9d16-df48f30e89b2"}
```

📌 Get Points for a Receipt

Replace <id> with the actual receipt ID:
```sh
curl -X GET "http://localhost:8080/receipts/<id>/points"
```

✅ Expected Response:
```sh
{"points": 31}
```

# 🐳 Docker Setup and Usage

📌 Verify Docker Installation
```sh
docker --version
```

✅ Expected output:
```sh
Docker version 27.4.0, build bde2b89
```

📌 Check Running Containers
```sh
docker ps
```

📌 Stop a Running Container
```sh
docker stop <container_id>
```

📌 Remove a Container
```sh
docker rm <container_id>
```

📌 Remove an Image
```sh
docker rmi receipt-processor
```

📌 Push Docker Image to DockerHub
1️⃣ Login to DockerHub
```sh
git clone https://github.com/namandiwan/fetch-receipt-processor-challenge.git
cd fetch-receipt-processor-challenge
```

## 2️⃣ Clone the Repository
```sh
git clone https://github.com/namandiwan/fetch-receipt-processor-challenge.git
cd fetch-receipt-processor-challenge
```

## 2️⃣ Clone the Repository
```sh
git clone https://github.com/namandiwan/fetch-receipt-processor-challenge.git
cd fetch-receipt-processor-challenge
```

## 2️⃣ Clone the Repository
```sh
git clone https://github.com/namandiwan/fetch-receipt-processor-challenge.git
cd fetch-receipt-processor-challenge
```

## 2️⃣ Clone the Repository
```sh
git clone https://github.com/namandiwan/fetch-receipt-processor-challenge.git
cd fetch-receipt-processor-challenge
```
