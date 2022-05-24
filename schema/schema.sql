CREATE TABLE author
(
    id   INT PRIMARY KEY AUTO_INCREMENT,
    name VARCHAR(255)
);
CREATE TABLE book
(
    id    INT PRIMARY KEY AUTO_INCREMENT,
    title VARCHAR(255)
);

CREATE TABLE books_authors
(
    book_id   INT,
    author_id INT,
    FOREIGN KEY (author_id) REFERENCES author (id),
    FOREIGN KEY (book_id) REFERENCES book (id),
    primary key (book_id, author_id)
);


INSERT INTO book (title)
VALUES ('Евгений Онегин');

INSERT INTO author (name)
VALUES ('Пушкин');

INSERT INTO books_authors (book_id, author_id)
VALUES (1, 1);

INSERT INTO book (title)
VALUES ('Метель');

INSERT INTO book (title)
VALUES ('Пометель');


INSERT INTO author (name)
VALUES ('Пупкин');

INSERT INTO books_authors (book_id, author_id)
VALUES (2, 1);

INSERT INTO books_authors (book_id, author_id)
VALUES (3, 2);