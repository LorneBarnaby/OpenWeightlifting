# OpenWeightlifting Makefile
# Shortcuts to the most common tools should be implemented here.

# Installs the python tools used to update the database
.PHONY: install_tools
install_tools:
	echo "Installing python tools"
	@cd python_tools && pipenv install

# Runs the python_tools used to update the database
.PHONY: update_db
update_db:
	echo "Updating the database"
	@cd python_tools && pipenv run python3 backend_cli.py --update all

# Stages and commits locally all the new csv files added to the event_data folder
.PHONY: stage_csv
stage_csv:
	echo "Staging csv files"
	@git add backend/event_data/\*.csv
	@git status --p --short | grep backend/event_data
	@git commit -m "Database Update"

.PHONY: check_db
DB ?= ""
check_db:
	@cd python_tools && pipenv run python3 check_db.py $(DB)

.PHONY: generate-docs
generate-docs:
	echo "Generating docs..."
	cd backend && swag init --parseDependency --parseInternal

# Removes build files.
.PHONY: clean
clean:
	rm -f backend/backend

# Removes build files plus cached dependencies.
.PHONY: veryclean
veryclean: clean
	rm -rf frontend/node_modules
