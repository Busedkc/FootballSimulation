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
