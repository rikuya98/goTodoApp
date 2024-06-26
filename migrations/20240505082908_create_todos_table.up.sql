CREATE TABLE todos (
  id INT AUTO_INCREMENT PRIMARY KEY,
  task VARCHAR(30) NOT NULL,
  due_date DATETIME NOT NULL,
  status TINYINT(1) NOT NULL,
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
);
