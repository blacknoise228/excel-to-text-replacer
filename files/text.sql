DO $$
BEGIN
    INSERT INTO user (name, age, email, country)
        VALUES ('$2', '$3', '$4', '$5');
END $$;
