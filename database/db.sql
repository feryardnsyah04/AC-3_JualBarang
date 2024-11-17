CREATE DATABASE shopping_cart;
USE shopping_cart;

CREATE TABLE cart_items (
  id INT AUTO_INCREMENT PRIMARY KEY,
  product VARCHAR(255) NOT NULL,
  category VARCHAR(255) NOT NULL,
  price INT NOT NULL,
  quantity INT NOT NULL
);