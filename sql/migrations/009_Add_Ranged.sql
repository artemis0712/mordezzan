-- +goose Up
-- Create ranged weapon properties table
CREATE TABLE ranged_weapon_properties (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    symbol TEXT NOT NULL,
    description TEXT NOT NULL
);

-- Insert ranged weapon properties
INSERT INTO
    ranged_weapon_properties (symbol, description)
VALUES
    ('⤢', 'Thrown weapon'),
    ('↵', 'Shield bypass'),
    ('⤤', 'Two-handed weapon');

-- Create ranged weapons table
CREATE TABLE ranged_weapons (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    name TEXT NOT NULL,
    weapon_type TEXT NOT NULL CHECK (weapon_type IN ('Hurled', 'Launched')),
    cost_gp INTEGER,
    weight INTEGER NOT NULL,
    rate_of_fire TEXT NOT NULL,
    range_short INTEGER,
    range_medium INTEGER,
    range_long INTEGER,
    damage TEXT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Create ranged weapon property links table
CREATE TABLE ranged_weapon_property_links (
    weapon_id INTEGER NOT NULL,
    property_id INTEGER NOT NULL,
    FOREIGN KEY (weapon_id) REFERENCES ranged_weapons (id),
    FOREIGN KEY (property_id) REFERENCES ranged_weapon_properties (id),
    PRIMARY KEY (weapon_id, property_id)
);

-- Insert hurled weapons
INSERT INTO
    ranged_weapons (
        name,
        weapon_type,
        cost_gp,
        weight,
        rate_of_fire,
        range_short,
        range_medium,
        range_long,
        damage
    )
VALUES
    ('Bola', 'Hurled', 3, 2, '1/1', 15, 30, 45, '1d2'),
    (
        'Boomerang',
        'Hurled',
        1,
        1,
        '1/1',
        50,
        100,
        150,
        '1d4'
    ),
    ('Dart', 'Hurled', 1, 1, '2/1', 15, 30, 45, '1d3'),
    (
        'Hooked Throwing Knife',
        'Hurled',
        20,
        2,
        '1/1',
        30,
        60,
        90,
        '1d6'
    ),
    (
        'Lasso',
        'Hurled',
        3,
        3,
        '1/2',
        20,
        NULL,
        NULL,
        NULL
    ),
    (
        'Net, Fighting',
        'Hurled',
        5,
        7,
        '1/2',
        10,
        NULL,
        NULL,
        NULL
    ),
    ('Stone', 'Hurled', 0, 1, '2/1', 30, 60, 90, '1'),
    (
        'Sling',
        'Hurled',
        2,
        1,
        '1/1',
        50,
        100,
        150,
        '1d4'
    );

-- Insert launched weapons
INSERT INTO
    ranged_weapons (
        name,
        weapon_type,
        cost_gp,
        weight,
        rate_of_fire,
        range_short,
        range_medium,
        range_long,
        damage
    )
VALUES
    (
        'Blowgun',
        'Launched',
        5,
        1,
        '1/1',
        30,
        60,
        90,
        '1'
    ),
    (
        'Bow, Long-',
        'Launched',
        60,
        3,
        '3/2',
        70,
        140,
        210,
        '1d6'
    ),
    (
        'Bow, Long-, Composite',
        'Launched',
        100,
        3,
        '3/2',
        80,
        160,
        240,
        '1d6'
    ),
    (
        'Bow, Short',
        'Launched',
        20,
        2,
        '3/2',
        50,
        100,
        150,
        '1d6'
    ),
    (
        'Bow, Short, Composite',
        'Launched',
        50,
        2,
        '3/2',
        60,
        120,
        180,
        '1d6'
    ),
    (
        'Crossbow, Heavy',
        'Launched',
        25,
        10,
        '1/2',
        80,
        160,
        240,
        '1d6+2'
    ),
    (
        'Crossbow, Light',
        'Launched',
        15,
        5,
        '1/1',
        60,
        120,
        180,
        '1d6+1'
    ),
    (
        'Crossbow, Repeating',
        'Launched',
        100,
        6,
        '3/1',
        50,
        100,
        150,
        '1d6'
    );

-- Link properties to weapons
INSERT INTO
    ranged_weapon_property_links
SELECT
    rw.id,
    rwp.id
FROM
    ranged_weapons rw,
    ranged_weapon_properties rwp
WHERE
    rw.name IN ('Bola', 'Boomerang', 'Dart', 'Stone', 'Sling')
    AND rwp.symbol = '⤢';

INSERT INTO
    ranged_weapon_property_links
SELECT
    rw.id,
    rwp.id
FROM
    ranged_weapons rw,
    ranged_weapon_properties rwp
WHERE
    rw.name = 'Hooked Throwing Knife'
    AND rwp.symbol IN ('⤢', '↵');

INSERT INTO
    ranged_weapon_property_links
SELECT
    rw.id,
    rwp.id
FROM
    ranged_weapons rw,
    ranged_weapon_properties rwp
WHERE
    rw.name IN ('Bow, Long-', 'Bow, Long-, Composite')
    AND rwp.symbol = '⤤';

-- +goose Down
DROP TABLE IF EXISTS ranged_weapon_property_links;

DROP TABLE IF EXISTS ranged_weapons;

DROP TABLE IF EXISTS ranged_weapon_properties;
