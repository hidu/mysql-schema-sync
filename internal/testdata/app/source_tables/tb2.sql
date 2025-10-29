CREATE TABLE tb2 (
  -- 整数类型
  id INT AUTO_INCREMENT PRIMARY KEY,
  tiny_int TINYINT,
  small_int SMALLINT,
  medium_int MEDIUMINT,
  normal_int INT,
  big_int BIGINT,

  -- 浮点与定点
  float_col FLOAT(10,2),
  double_col DOUBLE(16,4),
  decimal_col DECIMAL(18,6),

  -- 位与布尔
  bit_col BIT(8),
  bool_col BOOLEAN,

  -- 日期与时间
  date_col DATE,
  datetime_col DATETIME,
  timestamp_col TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  time_col TIME,
  year_col YEAR,

  -- 字符串类型
  char_col CHAR(10),
  varchar_col VARCHAR(255),

  -- 二进制字符串
  binary_col BINARY(8),
  varbinary_col VARBINARY(128),

  -- 文本类型
  tiny_text TINYTEXT,
  text_col TEXT,
  medium_text MEDIUMTEXT,
  long_text LONGTEXT,

  -- 二进制大对象
  tiny_blob TINYBLOB,
  blob_col BLOB,
  medium_blob MEDIUMBLOB,
  long_blob LONGBLOB,

  -- ENUM & SET
  enum_col ENUM('small','medium','large'),
  set_col SET('red','green','blue'),

  -- JSON
  json_col JSON,

  -- 空间类型（GIS）
  geom_col GEOMETRY,
  point_col POINT,
  linestring_col LINESTRING,
  polygon_col POLYGON,
  multipoint_col MULTIPOINT,
  multilinestring_col MULTILINESTRING,
  multipolygon_col MULTIPOLYGON,
  geomcollection_col GEOMETRYCOLLECTION,

  -- 其他辅助字段
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;