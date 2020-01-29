CREATE TABLE item (
  id int(11) NOT NULL AUTO_INCREMENT,
  name varchar(255) NOT NULL,
  price double(5,2) NOT NULL,
  category_id int(11) NOT NULL,
  in_sale tinyint(1) NOT NULL,
  added_time timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  last_updated timestamp NULL DEFAULT NULL,
  removed_time timestamp NULL DEFAULT NULL,
  PRIMARY KEY (id)
);