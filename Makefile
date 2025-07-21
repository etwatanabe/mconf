# Variables
API_IMAGE = mconf/api:candidato-1
RUNNER_IMAGE = mconf/runner:candidato-1

QUERY = Lord of the Rings

build:
	cd apps/api && docker build -t $(API_IMAGE) .
	cd apps/runner && docker build -t $(RUNNER_IMAGE) .

api:
	docker run -ti --rm --name mconf-api -p 3000:3000 $(API_IMAGE)

runner:
	docker run -ti --rm --name mconf-runner -e API_PORT=3000 $(RUNNER_IMAGE) "$(QUERY)"

down:
	docker stop mconf-api

clean:
	docker stop mconf-api
	docker image rm $(API_IMAGE)
	docker image rm $(RUNNER_IMAGE)