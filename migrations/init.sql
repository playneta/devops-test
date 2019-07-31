CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    name character varying(255) NOT NULL,
    email character varying(255) NOT NULL UNIQUE
);

INSERT INTO "users" ("id","name","email")
VALUES
(1, 'Pavel Makarenko', 'p.makarenko@playneta.gg'),
(2, 'Maksim Moskalik', 'm.moskalik@playneta.gg'),
(3, 'Sergey Fugin', 's.fugin@playneta.gg');