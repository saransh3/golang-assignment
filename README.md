# Golang Assignment: Import and Manage Data

### 🖍 Objective  
This project is a Golang-based web application to **upload an Excel file**, **store its data in a MySQL database**, and **cache it in Redis**. It also provides a **CRUD API** for managing the imported data.

---

## ⚙️ Features
- **Upload Excel File** (`POST /upload`): Asynchronous file processing in the background.  
- **Fetch Records** (`GET /records`): Fetches data from Redis cache (if available) or MySQL.  
- **Edit Record** (`PUT /records/:id`): Updates a specific record and refreshes the cache.  
- **Redis Caching**: Caches data for 5 minutes to reduce database load.  

---

## 🔧 Technologies Used
- **Golang** (`gin` framework) – REST API  
- **MySQL** – Relational database for storing records  
- **Redis** – In-memory data store for caching  
- **`excelize`** – Library for parsing `.xlsx` files  
- **Docker (Optional)** – For easier setup  

---

## 🖂 Project Structure
```
golang-assignment/
  ├── config/           # Database and Redis setup
  ├── handlers/         # API handlers
  ├── models/           # Data models and database functions
  ├── utils/            # Utility functions (Excel parsing, record insertion)
  └── main.go           # Main entry point of the application
```

---

## 🚀 Setup Instructions

### 1️⃣ Prerequisites
- **Golang** (>= 1.16)  
- **MySQL** (running on port `3306`)  
- **Redis** (running on port `6379`)  
- **Postman** (for testing APIs)  

### 2️⃣ Clone the Repository
```bash
git clone https://github.com/your-repo/golang-assignment.git
cd golang-assignment
```

### 3️⃣ Install Dependencies
```bash
go mod tidy
```

### 4️⃣ Configure MySQL and Redis
- **MySQL**: Create a database called `golang_assignment` and a `records` table:
  ```sql
  CREATE DATABASE golang_assignment;
  USE golang_assignment;

  CREATE TABLE records (
      id INT AUTO_INCREMENT PRIMARY KEY,
      first_name VARCHAR(100),
      last_name VARCHAR(100),
      company_name VARCHAR(100),
      address VARCHAR(200),
      city VARCHAR(100),
      country VARCHAR(100),
      postal VARCHAR(20),
      phone VARCHAR(20),
      email VARCHAR(100),
      web VARCHAR(100)
  );
  ```
- **Redis**: Ensure Redis is running on `127.0.0.1:6379`.

### 5️⃣ Run the Application
```bash
go run main.go
```

---

## 🔍 API Documentation

### 1. Upload Excel File (`POST /upload`)
- **Description**: Uploads an Excel file and processes it in the background.  
- **Request**:  
  - **Method**: `POST`  
  - **Body**: `form-data`, Key: `file` (upload a `.xlsx` file)

**Example Response:**
```json
{ "message": "File is being processed" }
```

---

### 2. Fetch Records (`GET /records`)
- **Description**: Fetches all records from Redis (if cached) or MySQL.  
- **Request**:  
  - **Method**: `GET`

**Example Response:**
```json
[
  {
    "id": 1,
    "first_name": "John",
    "last_name": "Doe",
    "company_name": "ABC Corp",
    "address": "123 St",
    "city": "NYC",
    "country": "USA",
    "postal": "10001",
    "phone": "1234567",
    "email": "john@abc.com",
    "web": "abc.com"
  }
]
```

---

### 3. Edit Record (`PUT /records/:id`)
- **Description**: Updates a specific record in MySQL and refreshes the cache.  
- **Request**:  
  - **Method**: `PUT`  
  - **Body**: `JSON`
```json
{
  "first_name": "John",
  "last_name": "Updated",
  "company_name": "New Corp",
  "address": "789 St",
  "city": "Boston",
  "country": "USA",
  "postal": "02101",
  "phone": "5551234",
  "email": "john.updated@newcorp.com",
  "web": "newcorp.com"
}
```

**Example Response:**
```json
{ "message": "Record updated successfully" }
```

---

## ⚠️ Error Handling
- **400 Bad Request**: Invalid input or missing file.  
- **500 Internal Server Error**: Database or cache errors.  

---

## 🛡️ Best Practices Implemented
- **Asynchronous File Processing**: File is processed in the background without blocking the client.  
- **Redis Caching**: Reduces database load by caching records for 5 minutes.  
- **Structured Error Handling**: Ensures graceful failure handling and logging.  

---

## 🤝 Contributing
Feel free to open an issue or submit a pull request!

---

