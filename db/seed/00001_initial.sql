-- +goose Up
INSERT INTO widgets (name, description) VALUES
  ('Widget Alpha', 'First example widget'),
  ('Widget Beta', 'Second example widget'),
  ('Widget Gamma', 'Third example widget'),
  ('Widget Delta', 'Fourth example widget'),
  ('Widget Epsilon', 'Fifth example widget'),
  ('Widget Zeta', 'Sixth example widget'),
  ('Widget Eta', 'Seventh example widget'),
  ('Widget Theta', 'Eighth example widget'),
  ('Widget Iota', 'Ninth example widget'),
  ('Widget Kappa', 'Tenth example widget');

-- Widget Alpha: 3 components
INSERT INTO components (widget_id, name, complexity) VALUES
  ((SELECT id FROM widgets WHERE name = 'Widget Alpha'), 'Red', 10),
  ((SELECT id FROM widgets WHERE name = 'Widget Alpha'), 'Green', 20),
  ((SELECT id FROM widgets WHERE name = 'Widget Alpha'), 'Blue', 30);

-- Widget Beta: 2 components
INSERT INTO components (widget_id, name, complexity) VALUES
  ((SELECT id FROM widgets WHERE name = 'Widget Beta'), 'Yellow', 15),
  ((SELECT id FROM widgets WHERE name = 'Widget Beta'), 'Purple', 25);

-- Widget Gamma: 5 components
INSERT INTO components (widget_id, name, complexity) VALUES
  ((SELECT id FROM widgets WHERE name = 'Widget Gamma'), 'Orange', 12),
  ((SELECT id FROM widgets WHERE name = 'Widget Gamma'), 'Cyan', 22),
  ((SELECT id FROM widgets WHERE name = 'Widget Gamma'), 'Magenta', 32),
  ((SELECT id FROM widgets WHERE name = 'Widget Gamma'), 'Lime', 42),
  ((SELECT id FROM widgets WHERE name = 'Widget Gamma'), 'Pink', 52);

-- Widget Delta: 0 components

-- Widget Epsilon: 1 component
INSERT INTO components (widget_id, name, complexity) VALUES
  ((SELECT id FROM widgets WHERE name = 'Widget Epsilon'), 'Teal', 18);

-- Widget Zeta: 4 components
INSERT INTO components (widget_id, name, complexity) VALUES
  ((SELECT id FROM widgets WHERE name = 'Widget Zeta'), 'Brown', 14),
  ((SELECT id FROM widgets WHERE name = 'Widget Zeta'), 'Black', 24),
  ((SELECT id FROM widgets WHERE name = 'Widget Zeta'), 'White', 34),
  ((SELECT id FROM widgets WHERE name = 'Widget Zeta'), 'Gray', 44);

-- Widget Eta: 2 components
INSERT INTO components (widget_id, name, complexity) VALUES
  ((SELECT id FROM widgets WHERE name = 'Widget Eta'), 'Violet', 19),
  ((SELECT id FROM widgets WHERE name = 'Widget Eta'), 'Indigo', 29);

-- Widget Theta: 5 components
INSERT INTO components (widget_id, name, complexity) VALUES
  ((SELECT id FROM widgets WHERE name = 'Widget Theta'), 'Maroon', 16),
  ((SELECT id FROM widgets WHERE name = 'Widget Theta'), 'Olive', 26),
  ((SELECT id FROM widgets WHERE name = 'Widget Theta'), 'Navy', 36),
  ((SELECT id FROM widgets WHERE name = 'Widget Theta'), 'Silver', 46),
  ((SELECT id FROM widgets WHERE name = 'Widget Theta'), 'Gold', 56);

-- Widget Iota: 0 components

-- Widget Kappa: 3 components
INSERT INTO components (widget_id, name, complexity) VALUES
  ((SELECT id FROM widgets WHERE name = 'Widget Kappa'), 'Turquoise', 21),
  ((SELECT id FROM widgets WHERE name = 'Widget Kappa'), 'Peach', 31),
  ((SELECT id FROM widgets WHERE name = 'Widget Kappa'), 'Lavender', 41);

-- +goose Down
TRUNCATE TABLE components RESTART IDENTITY CASCADE;
TRUNCATE TABLE widgets RESTART IDENTITY CASCADE;

