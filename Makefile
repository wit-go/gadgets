#a git remote add github git@github.com:wit-go/gadgets.git
all:
	@echo
	@echo gadgets are collections of widget primaties that work
	@echo in ways that are common enough they are generally useful
	@echo
	@echo The gadgets are themselves outside of the 'gui' binary tree
	@echo structure. They do not modify the fundamental nature of
	@echo how the 'gui' package communicates to the toolkits via the plugin
	@echo
	@echo in that sense, they are intended to be useful wrappers
	@echo

redomod:
	rm -f go.*
	go mod init
	go mod tidy

github:
	git push origin master
	git push origin devel
	git push origin --tags
	git push github master
	git push github devel
	git push github --tags

goproxy:
	GOPROXY=https://proxy.golang.org GO111MODULE=on go get go.wit.com/gui/gadgets@v0.10
