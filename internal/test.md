

```sql
CREATE DATABASE IF NOT EXISTS mss_test_db0
  CHARACTER SET utf8mb4
  COLLATE utf8mb4_unicode_ci;

CREATE DATABASE IF NOT EXISTS mss_test_db1
  CHARACTER SET utf8mb4
  COLLATE utf8mb4_unicode_ci;

GRANT ALL PRIVILEGES ON mss_test_db0.* TO 'mss_test'@'localhost' IDENTIFIED BY 'mss_test_pass';
GRANT ALL PRIVILEGES ON mss_test_db1.* TO 'mss_test'@'localhost' IDENTIFIED BY 'mss_test_pass';

FLUSH PRIVILEGES;
```