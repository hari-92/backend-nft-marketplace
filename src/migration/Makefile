diff-mysql:
	atlas migrate diff --env mysql

new-migration:
	@read -p "Enter migration name: " name; \
	atlas migrate new $$name --env mysql

diff-mysql-to:
	@read -p "Enter migration file name: " name; \
	atlas migrate diff --env mysql --to file://migrations/mysql/$$name.sql

diff-postgres:
	atlas migrate diff --env postgres

apply-mysql:
	atlas migrate apply --env mysql

hash:
	atlas migrate hash

reset-mysql:
	@echo "Cleaning up migration files..."
	@rm -f migrations/*.sql
	@echo "Done! Now you can run 'make diff-mysql' to create new migration"
