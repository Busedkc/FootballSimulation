# FootballSimulation 

This project emulates a football league where teams such as Liverpool, Manchester United, Manchester City, 
and Chelsea compete in random matches. The teams' scores are updated according to the match outcomes, and the 
league standings are determined accordingly. The project is implemented using the Go programming language and a MySQL database.

# File Descriptions 

## 1. main.go
The main function in this code initializes the random number generator to ensure consistent randomness throughout the application, sets up the database connection, and configures the HTTP server by defining available endpoints and their handlers. The server then starts running on http://localhost:8080, and any errors encountered during this process are logged and cause the application to terminate.

## 2. logic.go
This code defines a function simulateMatch that simulates a football match between two teams and updates their scores and goal differences. The function generates random scores for the home and away teams based on their respective strengths. After determining the scores, it updates the goal differences. If the home team wins, the home team's goal difference increases, and the away team's goal difference decreases by the same margin. Conversely, if the away team wins, the home team's goal difference decreases, and the away team's goal difference increases accordingly.

## 3. database.go
This code initializes a connection to a MySQL database using the go-sql-driver/mysql driver and defines several functions for interacting with the database. The initDB function establishes the database connection using credentials and connection details specified in the sql.Open function. It verifies the connection with db.Ping() and logs any errors if the connection fails. The getTeams function retrieves all teams from the teams table, scans each row into a Team struct, and returns a slice of Team structs. Similarly, the getMatches function retrieves all matches from the matches table, scans each row into a Match struct, and includes detailed information for each team by calling the getTeamByID function, which fetches a team's details by its ID. The saveMatch function inserts a new match result into the matches table, using the team IDs and match scores provided in the Match struct. Error handling is included in each function to ensure any issues are logged or returned appropriately.

## 4. models.go
The code defines three main structures to represent elements of a football league. The Team structure represents a football team with attributes such as ID, name, strength, and goal difference. The Match structure represents a football match between two teams, including the respective scores for the home and away teams. The League structure encapsulates a collection of teams and matches, providing a comprehensive model of a football league.

## 5. handlers.go
The provided code defines three HTTP handlers for a football league application. The getTeamsHandler function handles requests to retrieve all teams, fetching them from the database, setting the response content type to JSON, and encoding the teams as JSON before sending the response. The getMatchesHandler function performs a similar role for retrieving all matches, fetching them from the database, setting the response content type to JSON, and encoding the matches as JSON before sending the response. The simulateMatchHandler function handles requests to simulate a match; it decodes the request body into a Match struct, simulates the match, saves the match result to the database, sets the response content type to JSON, and encodes the match result as JSON before sending the response.

## 6. routes.go
The setupRoutes function configures the HTTP routes for the web server. It maps the URL path "/teams" to the getTeamsHandler function, which handles requests for retrieving all teams. The URL path "/matches" is mapped to the getMatchesHandler function, which handles requests for retrieving all matches. Additionally, the URL path "/simulate" is mapped to the simulateMatchHandler function, which handles requests for simulating a match. These route configurations ensure that the appropriate handler functions are called when the respective endpoints are accessed.

# Installation and Usage

## 1. MySQL Database Setup

CREATE TABLE teams (
    id INT AUTO_INCREMENT PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    strength INT NOT NULL,
    goal_difference INT NOT NULL DEFAULT 0
);

CREATE TABLE matches (
    id INT AUTO_INCREMENT PRIMARY KEY,
    home_team_id INT NOT NULL,
    away_team_id INT NOT NULL,
    home_score INT NOT NULL,
    away_score INT NOT NULL,
    FOREIGN KEY (home_team_id) REFERENCES teams(id),
    FOREIGN KEY (away_team_id) REFERENCES teams(id)
);

INSERT INTO teams (id, name, strength, goal_difference) VALUES (1, 'Chelsea', 95, 0);
INSERT INTO teams (id, name, strength, goal_difference) VALUES (2, 'Arsenal', 97, 0);
INSERT INTO teams (id, name, strength, goal_difference) VALUES (3, 'Manchester City', 87, 0);
INSERT INTO teams (id, name, strength, goal_difference) VALUES (4, 'Liverpool', 85, 0);


## 2. Download Project

git clone https://github.com/Busedkc/FootballSimulation.git
cd FootballSimulation

## 3. Start the Server

go run main.go database.go logic.go models.go handlers.go routes.go 

## 4. Usage
open your browser and go to http://localhost:8080 


