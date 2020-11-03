// Create User
curl -i -X POST -H "Content-Type: application/json" -d '{
"username":"Budi",
"email":"budikurniawan238@gmail.com",
"phone":"081290858473",
"password":"payphone16"
}' https://mighty-meadow-64200.herokuapp.com/api/v1/users

// Login
curl -i -X POST -H "Content-Type: application/json" -d '{"email":"budikurniawan238@gmail.com","password":"payphone16"}' http://localhost:8080/login

//
https://git.heroku.com/stormy-tundra-37100.git
https://stormy-tundra-37100.herokuapp.com/
