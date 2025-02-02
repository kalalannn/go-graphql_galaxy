SELECT c.id as character_id, c.name, n.id as nemesis_id, n.is_alive, n.years, s.id as secret_id, s.secret_code
FROM character c
JOIN nemesis n ON c.id = n.character_id
JOIN secret s ON n.id = s.nemesis_id;