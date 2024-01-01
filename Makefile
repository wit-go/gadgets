# git remote add github git@github.com:wit-go/gadgets.git

github:
	git push -u github master
	git push -u github devel
	git push github --tags
