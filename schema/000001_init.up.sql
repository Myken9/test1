
CREATE TABLE author(
                       id INT PRIMARY KEY AUTO_INCREMENT,
                       name VARCHAR(255)
);
CREATE TABLE book (
                      id INT PRIMARY KEY AUTO_INCREMENT,
                      title VARCHAR(255)
);

CREATE TABLE books_authors(
                              book_id INT,
                              author_id INT,
                              FOREIGN KEY (author_id)  REFERENCES author (id),
                              FOREIGN KEY (book_id)  REFERENCES book (id),
                              primary key (book_id, author_id)
);