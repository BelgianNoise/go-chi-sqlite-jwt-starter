package database

var Triggers = []string{
	`
		CREATE TRIGGER update_user_timestamp
		AFTER UPDATE ON user
		FOR EACH ROW
		BEGIN
			UPDATE user SET updated_at = CURRENT_TIMESTAMP WHERE id = OLD.id;
		END;
	`,
	`
		CREATE TRIGGER update_category_group_timestamp
		AFTER UPDATE ON category_group
		FOR EACH ROW
		BEGIN
			UPDATE category_group SET updated_at = CURRENT_TIMESTAMP WHERE id = OLD.id;
		END;
	`,
	`
		CREATE TRIGGER update_category_timestamp
		AFTER UPDATE ON category
		FOR EACH ROW
		BEGIN
			UPDATE category SET updated_at = CURRENT_TIMESTAMP WHERE id = OLD.id;
		END;
	`,
}
