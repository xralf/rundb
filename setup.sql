drop table if exists suppliers;
create table suppliers(name text, address text);
select * from suppliers;

drop table if exists products;
create table products(name text, category text, sku text);
insert into products (category, sku, name) values
('paint', 'sku1', 'pink super-gloss'),
('paint', 'sku2', 'glossy yellow'),
('paint', 'sku3', 'outdoors yellow'),
('paint', 'sku4', 'feeling so indoorsy'),
('tiles', 'sku5', 'blue concrete'),
('tiles', 'sku6', 'pink porcelain');
select * from products;
