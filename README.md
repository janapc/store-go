<div align="center">
  <h1>Store Go</h1>
  <img alt="Last commit" src="https://img.shields.io/github/last-commit/janapc/store-go"/>
  <img alt="Language top" src="https://img.shields.io/github/languages/top/janapc/store-go"/>
  <img alt="Repo size" src="https://img.shields.io/github/repo-size/janapc/store-go"/>
</div>
<div align="center">
 <a href="#-project">Project</a>&nbsp;&nbsp;&nbsp;|&nbsp;&nbsp;&nbsp;
<a href="#-requirement">Requirement</a>&nbsp;&nbsp;&nbsp;|&nbsp;&nbsp;&nbsp;
<a href="#-runner">Requirement</a>&nbsp;&nbsp;&nbsp;|&nbsp;&nbsp;&nbsp;
<a href="#-technologies">Technologies</a>&nbsp;&nbsp;&nbsp;|&nbsp;&nbsp;&nbsp;
</div>

## ⭐️ Project

This project is a store manager, where you can insert, delete, and edit the products.

## 📍 Requirement

You will need to have installed golang on your machine, my version is 1.19. Docker to run of database or you can use another way.

## 🚀 run

In project root:

To initial the database run the command below:

```sh
docker-compose up -d
```
To create a table in the database you can use this command in your database manager:

```sql
CREATE TABLE products(
id serial primary key,
name varchar,
description varchar,
price decimal,
amount integer
)
```

To run the project use the command below:

```sh
go run main.go
```

The project uses port _8000_, on your browser access **http://localhost:8000/**

## ⚙️ Technologies

- golang
- postgres

<br>
<span align="center">

Made by Janapc 🤘 [Get in touch!](https://www.linkedin.com/in/janaina-pedrina/)

</span>
