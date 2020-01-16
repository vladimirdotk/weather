CREATE TABLE WEATHER (
    id serial PRIMARY KEY,
    --@todo: use postgis
    lat varchar(20),
    long varchar(20),
    temp numeric,
    humidity numeric,
    pressure numeric
);