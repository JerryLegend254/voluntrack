tailwind:
	@tailwindcss -i ui/static/css/main.css -o ui/static/css/styles.css --watch

air:
	@air


run:
	$(MAKE) tailwind & $(MAKE) air
