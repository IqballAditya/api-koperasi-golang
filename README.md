# Cooperative Management RESTful API

A simple and fast backend REST API service for managing cooperative (Koperasi) member data. This project serves as a micro-portfolio demonstrating backend development using Golang, the Gin framework, and a relational MySQL database.

## Features
* **GET `/anggota`**: Retrieve a list of all cooperative members.
* **POST `/anggota`**: Add a new member to the cooperative system.

## Tech Stack
* **Language**: Golang (Go)
* **Framework**: [Gin](https://gin-gonic.com/) (HTTP web framework)
* **Database**: MySQL
* **Driver**: `go-sql-driver/mysql`

## Prerequisites
Before running this project, make sure you have the following installed:
* [Go](https://go.dev/dl/) (version 1.16 or higher)
* [XAMPP](https://www.apachefriends.org/index.html) (for local MySQL database)
* Git

## Database Setup
1. Start the **MySQL** module on your XAMPP Control Panel.
2. Open phpMyAdmin (`http://localhost/phpmyadmin`).
3. Create a new database named `db_koperasi`.
4. Run the following SQL query to create the `anggota` table:

```sql
CREATE TABLE `anggota` (
  `id_anggota` int(11) NOT NULL AUTO_INCREMENT,
  `nama` varchar(100) NOT NULL,
  `alamat` text NOT NULL,
  `total_simpanan` int(11) DEFAULT 0,
  PRIMARY KEY (`id_anggota`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;