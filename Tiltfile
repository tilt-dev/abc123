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

k8s_resource_assembly_version(2)

username = str(local('whoami')).rstrip('\n')


def with_username(s):
    return '%s-%s' % (username, s)


def m4_yaml(file):
    read_file(file)
    return local('m4 -Dvarowner=%s %s' % (username, repr(file)))


k8s_yaml([
    m4_yaml('fe/deployments/fe.yaml'),
    m4_yaml('letters/deployments/letters.yaml'),
    m4_yaml('numbers/deployments/numbers.yaml'),
])

repo = local_git_repo('.')
dockerfile_go = 'Dockerfile.go.base'
dockerfile_js = 'Dockerfile.js.base'
dockerfile_py = 'Dockerfile.py.base'

# Service: frontend
fe_img = 'gcr.io/windmill-public-containers/abc123/fe'
fe_entrypt = '/go/bin/fe --owner ' + username

fast_build(fe_img, dockerfile_go, fe_entrypt).\
    add(repo.path('fe'), '/go/src/github.com/windmilleng/abc123/fe').\
    run('go install github.com/windmilleng/abc123/fe')
k8s_resource(with_username('fe'), new_name='fe', port_forwards=9000)

# Service: letters
letters_img = 'gcr.io/windmill-public-containers/abc123/letters'
letters_entrypt = 'node /app/index.js'

fast_build(letters_img, dockerfile_js, letters_entrypt).\
    add(repo.path('letters/src'), '/app').\
    add(repo.path('letters/package.json'), '/app/package.json').\
    add(repo.path('letters/yarn.lock'), '/app/yarn.lock').\
    run('cd /app && yarn install', trigger=['letters/package.json', 'letters/yarn.lock'])
k8s_resource(with_username('letters'), new_name='letters')

# Service: numbers
numbers_img = 'gcr.io/windmill-public-containers/abc123/numbers'
numbers_entrypt = 'node /app/index.js'

fast_build(numbers_img, dockerfile_py). \
    add(repo.path('numbers'), '/app'). \
    run('cd /app && pip install -r requirements.txt', trigger='numbers/requirements.txt')
k8s_resource(with_username('numbers'), new_name='numbers')
