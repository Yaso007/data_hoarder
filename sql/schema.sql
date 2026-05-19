CREATE TABLE IF NOT EXISTS users (

    id SERIAL PRIMARY KEY,

    username VARCHAR(100)
    UNIQUE
    NOT NULL,

    password VARCHAR(255)
    NOT NULL
);

CREATE TABLE IF NOT EXISTS fir_records (

    id SERIAL PRIMARY KEY,

    police_station_code VARCHAR(100)
    UNIQUE
    NOT NULL,

    law_enforcement_agency VARCHAR(255),

    national_code VARCHAR(100),

    last_fir_no VARCHAR(100),

    last_fir_year VARCHAR(10),

    large_charge_sheet_date DATE,

    district VARCHAR(50)
    CHECK (
        district IN (
            'Gangtok',
            'Namchi',
            'Mangan',
            'Gyalshing',
            'Pakyong',
            'Soreng'
        )
    ),

    law_enforcement_type VARCHAR(100),

    display_status BOOLEAN
);