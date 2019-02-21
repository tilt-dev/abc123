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
    * Other notes: does a `pip install` for package dependencies. Reinstalls dependencies iff they have changed.
"""
### Step 0: Hello World ###
# Uncomment this to see Tilt do something!
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
# k8s_resource('fe', port_forwards='8000')

### Step 3: Tilt Image Build ###
# # Uncomment these lines to have Tilt build images for your code (and
# # smartly re-build whenever things change on disk
# docker_build('abc123/fe', 'fe')            # == `docker build ./fe -t abc123/fe`
# docker_build('abc123/letters', 'letters')  # == `docker build ./letters -t abc123/letters`
# docker_build('abc123/numbers', 'numbers')  # == `docker build ./numbers -t abc123/numbers`
# k8s_resource('fe', port_forwards='8000') # HEY, LISTEN! delete the port-forward line above. Sorry ^_^'

### Step 4: Fast-Build ###
# # The frontend builds pretty slowly, doesn't it? That's because every time you change it,
# # Docker re-builds the container, which means it has to tar up a really big build context
# # because of that pesky `fe/big_context` directory (16KB). Let's use fast-build instead,
# # so only the files you update get moved around. Uncomment the lines below
# repo = local_git_repo('.')
# dockerfile_go = 'Dockerfile.go.base'
# fe_img = 'abc123/fe'
# fe_entrypt = '/go/bin/fe'
#
# # Service: fe
# fast_build(fe_img, dockerfile_go, fe_entrypt).\
#     add(repo.path('fe'), '/go/src/github.com/windmilleng/abc123/fe').\
#     run('go install github.com/windmilleng/abc123/fe')

### Step 5: Fast-Build ALL THE THINGS! ###
# # The other builds are pretty fast, but why not make them even faster? Uncomment
# # the rest of the Tiltfile to use fast-build for the other services as well.

# dockerfile_js = 'Dockerfile.js.base'
# dockerfile_py = 'Dockerfile.py.base'
#
# # Service: letters
# letters_img = 'abc123/letters'
# letters_entrypt = 'node /app/index.js'
#
# fast_build(letters_img, dockerfile_js, letters_entrypt).\
#     add(repo.path('letters/src'), '/app').\
#     add(repo.path('letters/package.json'), '/app/package.json').\
#     add(repo.path('letters/yarn.lock'), '/app/yarn.lock').\
#     run('cd /app && yarn install', trigger=['letters/package.json', 'letters/yarn.lock'])
#
# # Service: numbers
# numbers_img = 'abc123/numbers'
# numbers_entrypt = './app/app.py'
#
# fast_build(numbers_img, dockerfile_py). \
#     add(repo.path('numbers'), '/app'). \
#     run('cd /app && pip install -r requirements.txt', trigger='numbers/requirements.txt')
