CREATE DATABASE simulation;
USE simulation;

CREATE TABLE Teams (team_id INT AUTO_INCREMENT PRIMARY KEY,
                    team_name VARCHAR(100),
                    strength INT
   );
   
   SELECT * FROM Teams;
   
   CREATE TABLE Matches (match_id INT AUTO_INCREMENT PRIMARY KEY,
						 week INT,
                         home_team_id INT,
                         away_team_id INT,
                         home_goals INT,
                         away_goals INT,
                         match_result ENUM('home_win', 'draw', 'away_win'),
                         FOREIGN KEY (home_team_id) REFERENCES Teams(team_id),
                         FOREIGN KEY (away_team_id) REFERENCES Teams(team_id)
   
   );
   
   SELECT * FROM Matches;
   
   CREATE TABLE Standings (team_id INT,
                           played INT DEFAULT 0,
                           wins INT DEFAULT 0,
                           draws INT DEFAULT 0,
                           losses INT DEFAULT 0,
                           goals_for INT DEFAULT 0,
                           goals_against INT DEFAULT 0,
                           goal_difference INT DEFAULT 0,
                           points INT DEFAULT 0,
                           PRIMARY KEY (team_id),
                           FOREIGN KEY (team_id) REFERENCES Teams(team_id)
   );
   
   SELECT * FROM Standings;