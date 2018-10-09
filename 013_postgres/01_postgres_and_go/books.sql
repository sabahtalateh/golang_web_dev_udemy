CREATE TABLE books (
    isbn CHAR(14) PRIMARY KEY NOT NULL,
    title VARCHAR(255) NOT NULL,
    author VARCHAR(255) NOT NULL,
    price DECIMAL(5, 2) NOT NULL
);

INSERT INTO books (isbn, title, author, price)
VALUES
  ('172-1728278173', 'Emma', 'Jayne Austen', 9.44),
  ('172-1716216273', 'Zopa', 'Ivan Gog', 11.70),
  ('175-1823719263', 'Zopa 2', 'Petr Gog', 22.60);
