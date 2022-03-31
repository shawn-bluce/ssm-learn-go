USE SSM;

DROP TABLE IF EXISTS student;

CREATE TABLE student (
    id                  int         NOT NULL AUTO_INCREMENT PRIMARY KEY,
    name                varchar(10) NOT NULL,
    operating_system    float       NOT NULL,
    networking          float       NOT NULL,
    data_struct         float       NOT NULL,
    golang              float       NOT NULL
) ENGINE=InnoDB;
