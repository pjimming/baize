gh-pages:
	npm run build
	git add dist/
	git commit -m "gh-pages: push dist"
	git push
	cd ../ && git subtree push --prefix="frontend/dist" origin gh-pages