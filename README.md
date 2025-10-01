# Go Movie Rating Backend

Backend project built with Golang. I built mobile app with Flutter. If you want to see it, I'll leave the link below.
<a href="https://github.com/burakiscoding/flutter_movie_rating">Flutter Movie Rating App</a>

## Tech Stack

- Flutter for mobile
- Golang for backend
- MySQL for database

# Main features of the project

- Handling HTTP requests
- JWT authentication
- File upload & serve
- MySQL table creation & queries

## What could I have done better

- I could store the image paths in the database. I didn't do it because all images start with the same prefix(http://localhost:8080/movies/) and I wanted to save some space. But now, I have to create the path every time before I return to the frontend.

## Developer's note
I didn't complete some of the endpoints because I don't want to lose so much time on a portfolio project.

## How to install

Clone the repo
```
git clone https://github.com/burakiscoding/go-movie-rating.git
```
Move into the project directory
```
cd go-movie-rating
```
Setup the environment variables for database username and database password. This won't save your variables permanently. If you want to save them permanently you should save them in a place where you can set your environment variables. This place could be ~/.profile, ~/.bashrc, etc. But we don't need to save them permanently right now.
```
export DBUSER=<yourusername>
export DBPASS=<yourpassword>
```
Now we need to run MySQL. If you don't have MySQL in your machine go to <a href="https://dev.mysql.com/doc/mysql-installation-excerpt/5.7/en">MySQL Installation Guide</a>
After installing the MySQL, we need to create database for our application. Don't worry, I've already created sql scripts for initialization and seeding. Open a new terminal and move to the sql_scripts directory in our project. And run MySQL in this terminal and after that run the sql scripts.
To run MySQL:
```
mysql -u root -p
```
And then run the scripts
```
source init.sql
source insert.sql
```
Now you can run the app. Run this command in the project folder(go-movie-rating)
```
go run .
```
