-- Seeds database

INSERT INTO movies 
(name, description, rating, duration_in_minutes, release_date, language_code, number_of_ratings) 
VALUES 
('Madame Web', 
'Forced to confront her past, Cassandra Webb, a Manhattan paramedic that may have clairvoyant abilities, forms a relationship with three young women destined for powerful futures, if they can survive their threatening present.',
4.1, 116, '2024-02-14', 'en', 107000),
('Batman Begins', 
'After witnessing his parents death, billionaire Bruce Wayne learns the art of fighting to confront injustice. When he returns to Gotham as Batman, he must stop a secret society that intends to destroy the city.',
8.2, 140, '2005-06-10', 'en', 1700000),
('Patch Adams', 
'The true story of the heroic Hunter "Patch" Adams whos determined to become a medical doctor because of his desire to help other people. He ventures where no doctor ever ventured before utilizing humor and pathos.',
6.9, 115, '1998-12-25', 'en', 128000),
('Angel-A', 
'A beautiful woman helps an inept scam artist get his game together.',
7.0, 91, '2005-12-21', 'fr', 37000);

INSERT INTO actors 
(first_name, last_name) 
VALUES 
('Dakota', 'Johnson'), ('Sydney', 'Sweeney'), ('Isabela', 'Merced'),
('Christian', 'Bale'), ('Michael', 'Caine'), ('Ken', 'Watanabe'),
('Robin', 'Williams'), ('Daniel', 'London'), ('Monica', 'Potter'),
('Rie', 'Rasmussen'), ('Jamel', 'Debbouze'), ('Gilbert', 'Melki');

INSERT INTO genres (name) 
VALUES 
('Superhero'), ('Action'), ('Thriller'), ('Action Epic'), ('Epic'), ('Tragedy'), ('Crime'), ('Drama'),
('Docudrama'), ('Medical Drama'), ('Biography'), ('Comedy'), ('Romance'), ('Fantasy');

INSERT INTO movie_genres 
(movie_id, genre_id) 
VALUES
(1, 1), (1, 2), (1, 3),
(2, 4), (2, 5), (2, 1), (2, 6), (2, 7), (2, 8), (2, 3),
(3, 9), (3, 10), (3, 11), (3, 12), (3, 8), (3, 13),
(4, 12), (4, 8), (4, 14), (4, 13);

INSERT INTO movie_actors 
(movie_id, actor_id) 
VALUES 
(1, 1), (1, 2), (1, 3), (2, 4), (2, 5), (2, 6), (3, 7), (3, 8), (3, 9), (4, 10), (4, 11), (4, 12);


INSERT INTO movie_medias 
(name, type, movie_id)
VALUES 
('madame-web-1.jpg', 2, 1),
('madame-web-2.jpg', 0, 1),
('madame-web-3.jpg', 0, 1),
('madame-web-4.jpg', 0, 1),
('madame-web-5.jpg', 0, 1),
('madame-web-6.jpg', 0, 1),
('madame-web-7.jpg', 0, 1),
('madame-web-8.jpg', 0, 1),
('madame-web-9.jpg', 0, 1),
('madame-web-clip-1.mp4', 1, 1),
('madame-web-clip-2.mp4', 1, 1),
('madame-web-trailer.mp4', 3, 1),
('batman-1.jpg', 2, 2),
('batman-2.jpg', 0, 2),
('batman-3.jpg', 0, 2),
('batman-4.jpg', 0, 2),
('batman-5.jpg', 0, 2),
('batman-6.jpg', 0, 2),
('batman-7.jpg', 0, 2),
('batman-8.jpg', 0, 2),
('batman-9.jpg', 0, 2),
('batman-10.jpg', 0, 2),
('batman-clip-1.mp4', 1, 2),
('batman-clip-2.mp4', 1, 2),
('batman-trailer.mp4', 3, 2),
('patch-adams-1.jpg', 0, 3),
('patch-adams-2.jpg', 0, 3),
('patch-adams-3.jpg', 0, 3),
('patch-adams-4.jpg', 0, 3),
('patch-adams-5.jpg', 0, 3),
('patch-adams-6.jpg', 0, 3),
('patch-adams-7.jpg', 0, 3),
('patch-adams-8.jpg', 0, 3),
('patch-adams-9.jpg', 0, 3),
('patch-adams-10.jpg', 2, 3),
('angel-a-1.jpg', 2, 4),
('angel-a-2.jpg', 0, 4),
('angel-a-3.jpg', 0, 4),
('angel-a-4.jpg', 0, 4),
('angel-a-5.jpg', 0, 4),
('angel-a-6.jpg', 0, 4),
('angel-a-7.jpg', 0, 4),
('angel-a-8.jpg', 0, 4),
('angel-a-9.jpg', 0, 4),
('angel-a-10.jpg', 0, 4),
('angel-a-trailer.mp4', 3, 4)