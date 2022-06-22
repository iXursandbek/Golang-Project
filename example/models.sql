create TABLE authors (
    author_id serial primary key,
    name varchar(64) );

create TABLE books ( 
    book_id serial primary key,
    name varchar(60) NOT null, 
    author_id int NOT NULL references authors(author_id), 
    category_id int NOT NULL references categories(category_id) );
    
create TABLE categories(
    category_id serial primary key, 
    name VARCHAR(128) NOT NULL );

insert INTO authors (name) VALUES ('Xursandbek'),('Khurshid'),('Abdulaziz'); 
insert into categories (name) VALUES ('Iqtisod'),('Programming'),('Meditsina');
insert INTO books (name, author_id, category_id) VALUES 
    ('Tirilish',1,1),
    ('Ilm olish sirlari',2,2),
    ('Kecha va kunduz',3,1),
    ('Otgan kunlar',1,3);

SELECT *
FROM books AS b
JOIN authors AS a
    ON b.author_id = a.author_id;

SELECT b.book_id AS book_id,
         b.name AS book_name,
         a.name AS author_name
FROM books AS b
JOIN authors AS a
    ON b.author_id = a.author_id;
    
SELECT b.book_id AS book_id,
         b.name AS book_name,
         a.name AS author_name,
         c.name AS category_name
FROM books AS b
JOIN authors AS a USING(author_id)
JOIN categories AS c
    ON b.category_id = c.category_id
WHERE c.category_id = 2
    AND a.author_id = 2;
    
insert INTO books (name, author_id, category_id) VALUES ('Secret',3,3);