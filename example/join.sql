create TABLE basket_a (
    id int PRIMARY KEY,
    fruit_name_a VARCHAR(100) NOT NULL
);

create TABLE basket_b (
    b INT PRIMARY KEY,
    fruit_name_b VARCHAR(100) NOT NULL
);

insert INTO basket_a (id, fruit_name_a)
VALUES
    (1, 'Apple'),
    (2, 'Orange'),
    (3, 'Banana'),
    (4, 'Cucumber');

insert INTO basket_b (b, fruit_name_b)
VALUES
    (1, 'Orange'),
    (2, 'Apple'),
    (3, 'Watermelon'),
    (4, 'Pear');

select id, fruit_name_a, b, fruit_name_b
FROM
    basket_a
inner JOIN basket_b
on fruit_name_a = fruit_name_b;

select 
    id as id,
    fruit_name_a as meva,
    b as id,
    fruit_name_b as meva2
FROM
    basket_a
left JOIN basket_b
on fruit_name_a = fruit_name_b;

SELECT *
FROM basket_a
right JOIN basket_b
ON id = b;

insert INTO basket_b (b, fruit_name_b) VALUES (5,'Olma');

SELECT
    id,
    fruit_name_a,
    b,
    fruit_name_b
FROM
    basket_a
RIGHT JOIN basket_b 
   ON fruit_name_a = fruit_name_b
WHERE id IS NULL;