DO $$
BEGIN
    INSERT INTO vacancy (title, company, address, salary)
        VALUES ('$1', '$2', '$3', '$4');
END $$;
