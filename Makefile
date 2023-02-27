install:
	@echo "Creating private/oauth file..."
	@mkdir -p private && touch private/oauth
	@if [ -z $(TOKEN) ]; then \
  		echo "Token not set, leaving oauth file empty."; \
	else \
	  echo "Adding token to file..."; \
	  echo $(TOKEN) > private/oauth; \
	fi
	@echo "Creating persistence files..."
	@mkdir -p persistence && touch persistence/rhinosfed

run:
	@go build
	@./rhino-bot.exe