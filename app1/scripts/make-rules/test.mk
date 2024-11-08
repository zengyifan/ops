.PHONY: test.api
test.api:
#	@./tests/api.sh insecure::hello
	@./tests/api.sh insecure::user
#	@./tests/api.sh insecure::group
#	@./tests/api.sh insecure::auth



