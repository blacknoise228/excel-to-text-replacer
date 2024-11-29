DO $$
BEGIN
    INSERT INTO users (name, age, email, address)
        VALUES (?, ?, ?, ?);
END $$;
