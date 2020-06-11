create table ta (id int not null auto_increment, xx char not null, yy double, primary key (id), key idx_ta(xx));
create table tb (id int not null auto_increment, xx char not null, zz char not null, ww double, primary key (id), key idx_tb(xx), constraint fk_tb foreign key (xx) REFERENCES ta (xx));
create table tc (id int not null auto_increment, pp char not null, qq double, primary key (id)) partition by hash(id);
