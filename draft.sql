-- Create table for Users
CREATE TABLE User (
    UserID INTEGER PRIMARY KEY AUTOINCREMENT,
    Username TEXT NOT NULL UNIQUE,
    Email TEXT NOT NULL UNIQUE,
    Password TEXT NOT NULL,
    SessionToken TEXT,
    SessionExpiration DATETIME
);

-- Create table for Categories
CREATE TABLE Category (
    CategoryID INTEGER PRIMARY KEY AUTOINCREMENT,
    CategoryName TEXT NOT NULL UNIQUE
);

-- Create table for Posts
CREATE TABLE Post (
    PostID INTEGER PRIMARY KEY AUTOINCREMENT,
    Title TEXT NOT NULL,
    Content TEXT NOT NULL,
    CreatedDate DATETIME DEFAULT CURRENT_TIMESTAMP,
    UserID INTEGER,
    CategoryID INTEGER,
    Likes INTEGER DEFAULT 0,
    Dislikes INTEGER DEFAULT 0,
    FOREIGN KEY (UserID) REFERENCES User(UserID),
    FOREIGN KEY (CategoryID) REFERENCES Category(CategoryID)
);

-- Create table for Comments
CREATE TABLE Comment (
    CommentID INTEGER PRIMARY KEY AUTOINCREMENT,
    Content TEXT NOT NULL,
    CreatedDate DATETIME DEFAULT CURRENT_TIMESTAMP,
    PostID INTEGER,
    UserID INTEGER,
    Likes INTEGER DEFAULT 0,
    Dislikes INTEGER DEFAULT 0,
    FOREIGN KEY (PostID) REFERENCES Post(PostID),
    FOREIGN KEY (UserID) REFERENCES User(UserID)
);

-- Create table for PostCategory (junction table for many-to-many relationship)
CREATE TABLE PostCategory (
    PostCategoryID INTEGER PRIMARY KEY AUTOINCREMENT,
    PostID INTEGER,
    CategoryID INTEGER,
    FOREIGN KEY (PostID) REFERENCES Post(PostID),
    FOREIGN KEY (CategoryID) REFERENCES Category(CategoryID)
);

-- Sample Insert Statements to populate the tables
INSERT INTO User (Username, Email, Password) VALUES ('john_doe', 'john@example.com', 'password123');
INSERT INTO User (Username, Email, Password) VALUES ('jane_doe', 'jane@example.com', 'password123');

INSERT INTO Category (CategoryName) VALUES ('General Discussion');
INSERT INTO Category (CategoryName) VALUES ('Announcements');
INSERT INTO Category (CategoryName) VALUES ('Feedback');

INSERT INTO Post (Title, Content, UserID, CategoryID) VALUES ('Welcome to the forum!', 'This is the first post.', 1, 1);
INSERT INTO Post (Title, Content, UserID, CategoryID) VALUES ('Forum Rules', 'Please read the rules.', 2, 2);

INSERT INTO Comment (Content, PostID, UserID) VALUES ('Thanks for the info!', 1, 2);
INSERT INTO Comment (Content, PostID, UserID) VALUES ('Will do!', 2, 1);

INSERT INTO PostCategory (PostID, CategoryID) VALUES (1, 1);
INSERT INTO PostCategory (PostID, CategoryID) VALUES (2, 2);
