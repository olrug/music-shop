CREATE TABLE SubCategories (
    SubCategoryId serial primary key,
    CategoryId int not null,
    Name varchar(50)
);