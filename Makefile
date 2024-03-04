
dev-go:
	air

dev-js:
	pnpm run dev

dev:
	/bin/bash -c "./node_modules/.bin/tailwindcss -i ./assets/postcss/main.css -o ./assets/style/main.css --watch" &
	/bin/bash -c "~/go/bin/air"
