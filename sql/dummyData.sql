-- Insert Users
INSERT INTO User (Username, FirstName, LastName, Email, Password, SessionToken, SessionExpiration) VALUES
('john_doe', 'John', 'Doe', 'john.doe@example.com', 'password123', NULL, NULL),
('jane_smith', 'Jane', 'Smith', 'jane.smith@example.com', 'password123', NULL, NULL),
('alice_jones', 'Alice', 'Jones', 'alice.jones@example.com', 'password123', NULL, NULL);

-- Insert Categories
INSERT INTO Category (CategoryName) VALUES
('General'),
('Announcements'),
('Feedback');

-- Insert Posts
INSERT INTO Post (Title, Content, UserID) VALUES
('Welcome to the forum', 'This is the first post', 1),
('Forum Guidelines', 'Please read the rules', 2),
('New Features', 'Check out the new features', 3),
('General Discussion', 'Feel free to talk about anything', 1),
('Site Updates', 'We have updated the site', 2),
('Suggestions', 'We welcome your suggestions', 3),
('Hello everyone', 'Nice to meet you all', 1),
('Important Update', 'Please read this important update', 2),
('Community Feedback', 'Share your thoughts', 3);

-- Insert Comments
INSERT INTO Comment (Content, PostID, UserID) VALUES
('Great post!', 1, 2),
('Very informative', 2, 3),
('I like the new features', 3, 1),
('Thanks for the info', 4, 3),
('Good to know', 5, 1),
('I have a suggestion', 6, 2),
('Hello!', 7, 2),
('Important indeed', 8, 3),
('I think this could be improved', 9, 1);

-- Insert PostCategory
INSERT INTO PostCategory (PostID, CategoryID) VALUES
(1, 1),
(2, 2),
(3, 3),
(4, 1),
(5, 2),
(6, 3),
(7, 1),
(8, 2),
(9, 3);

-- Insert Interactions
INSERT INTO Interaction (PostID, UserID, Kind) VALUES
(1, 2, 1), -- User 2 likes Post 1
(1, 3, 0), -- User 3 dislikes Post 1
(2, 1, 1), -- User 1 likes Post 2
(2, 3, 1), -- User 3 likes Post 2
(3, 1, 0), -- User 1 dislikes Post 3
(3, 2, 1), -- User 2 likes Post 3
(4, 3, 1), -- User 3 likes Post 4
(4, 2, 0), -- User 2 dislikes Post 4
(5, 1, 1), -- User 1 likes Post 5
(5, 3, 1), -- User 3 likes Post 5
(6, 1, 0), -- User 1 dislikes Post 6
(6, 2, 1); -- User 2 likes Post 6
