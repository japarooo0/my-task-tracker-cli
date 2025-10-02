.PHONY: default
default:
	@echo Available commands:
	@echo list [todo inprogress done]
	@echo add task=[task_name]
	@echo update id=[id] task=[task_name]
	@echo delete id=[id]
	@echo mark-in-progress id=[id]
	@echo mark-done id=[id]

list::
	go run main.go list

add::
	go run main.go add "$(task)"

update::
	go run main.go update $(id) "$(task)"

mark-in-progress::
	go run main.go mark-in-progress $(id)

mark-done::
	go run main.go mark-done $(id)

delete::
	go run main.go delete $(id)
