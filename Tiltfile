# -*- mode: Python -*-

"""
* Frontend
  * Language: Go
  * Other notes: calls out to both `letters` and `numbers` microservices.
* Letters
  * Language: JavaScript
  * Other notes: Uses yarn. Does a `yarn install` for package dependencies iff they have changed.
* Numbers
    * Language: Python
    * Other notes: does a `pip install` for package dependencies. Re-installs dependencies iff they have changed.
"""

### NOTE: this Tiltfile / incremental onboarding experience currently only works on LOCAL k8s clusters
### that don't require pushing images -- so, e.g. k8s for Docker for Mac, Minikube + Docker

### Step 0: Hello World ###
# # Uncomment this to see Tilt do something!
# print("Welcome to Tilt! 👋")

### Step 1: Kubernetes YAML ###
# # Start with just your existing Kubernetes yaml -- your services will run (if the
# # images are available) and you can see their status and logs, but Tilt isn't
# # building anything for you, so code changes won't be propagated
# # NOTE: build these images first with `make build` 👀
# k8s_yaml([
#     'fe/deployments/fe.yaml',
#     'letters/deployments/letters.yaml',
#     'numbers/deployments/numbers.yaml',
# ])

### Step 2: Port Forwarding ###
# # Uncomment this line to port-forward your frontend so you can hit it locally
# # Once you've done this, you'll be able to access the 'fe' service in your browser at http://localhost:8000
# k8s_resource('fe', port_forwards='8000')

### Step 3: Tilt Image Build ###
# # Uncomment these lines to have Tilt build images for your code (and
# # smartly re-build whenever things change on disk)
# # After this and following steps, try changing the first `log.Println` message in `fe/main.go` and see how Tilt behaves.
# docker_build('abc123/fe', 'fe')            # == `docker build ./fe -t abc123/fe`
# docker_build('abc123/letters', 'letters')  # == `docker build ./letters -t abc123/letters`
# docker_build('abc123/numbers', 'numbers')  # == `docker build ./numbers -t abc123/numbers`

### Step 4: Live Update ###
# # The frontend builds pretty slowly, doesn't it? That's because every time you change it,
# # Docker re-builds the container, which means it has to tar up a really big build context
# # because of that pesky `fe/big_context` directory (16KB). Let's use Live Update instead,
# # so only the files you update get moved around. Uncomment the lines below
# # NOTE: comment out the `docker_build` for fe above 👀

# # Service: fe
# docker_build('abc123/fe', 'fe',
#              live_update=[
#                  sync('./fe', '/go/src/github.com/windmilleng/abc123/fe'),
#                  run('go install github.com/windmilleng/abc123/fe'),
#                  restart_container()
#              ])

### Step 5: Live Update ALL THE THINGS! ###
# # The other builds are pretty fast, but why not make them even faster? Uncomment
# # the rest of the Tiltfile to use Live Update for the other services as well.
# # NOTE: comment out the rest of the `docker_build` calls above 👀
#
# # Service: letters
# docker_build('abc123/letters', 'letters',
#              live_update=[
#                  sync('./letters/src', '/app'),
#                  sync('./letters/package.json', '/app/package.json'),
#                  sync('./letters/yarn.lock', '/app/yarn.lock'),
#                  run('cd /app && yarn install', trigger=['./letters/package.json', './letters/yarn.lock']),
#                  restart_container(),
#              ])
#
# # Service: numbers
# docker_build('abc123/numbers', 'numbers',
#              live_update=[
#                  sync('./numbers', '/app'),
#                  run('cd /app && pip install -r requirements.txt', trigger='numbers/requirements.txt'),
#              ])
