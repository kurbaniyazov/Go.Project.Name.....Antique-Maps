ALTER TABLE antiquemaps ADD CONSTRAINT antiquemaps_year_check CHECK (year >= 1);

ALTER TABLE antiquemaps ADD CONSTRAINT antiquemaps_title_check CHECK (title != '');
