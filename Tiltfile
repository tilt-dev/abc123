# -*- mode: Python -*-

"""
* Frontend
  * Language: Go
  * Other notes: presents a grid of the results of calling all of the other services
* Fortune
  * Language: Go
  * Other notes: Uses protobufs
* Hypothesizer
  * Language: Python
  * Other notes: does a `pip install` for package dependencies. Reinstalls dependencies, only if the dependencies have changed.
* Spoonerisms
  * Language: JavaScript
  * Other notes: Uses yarn. Does a `yarn install` for package dependencies, only if the dependencies have changed
"""


def get_username():
  return local('whoami').rstrip('\n')


def m4_yaml(file):
  read_file(file)
  return local('m4 -Dvarowner=%s %s' % (repr(get_username()), repr(file)))


def abc123():
  return composite_service([fe, letters, numbers])


def fe():
  yaml = m4_yaml('fe/deployments/fe.yaml')

  image_name = 'gcr.io/windmill-public-containers/abc123/fe'

  start_fast_build('Dockerfile.go.base', image_name, '/go/bin/fe --owner ' + get_username())
  # path = '%s/src/github.com/windmilleng/abc123/fe' % gopath
  path = '/go/src/github.com/windmilleng/abc123/fe'

  repo = local_git_repo('.')
  add(repo.path('fe'), path)

  run('go install github.com/windmilleng/abc123/fe')
  img = stop_build()
  img.cache('/root/.cache/go-build/')

  s = k8s_service(img, yaml=yaml)
  s.port_forward(9000)
  return s


def letters():
  yaml = m4_yaml('letters/deployments/letters.yaml')

  image_name = 'gcr.io/windmill-public-containers/abc123/letters'

  start_fast_build('Dockerfile.js.base', image_name, 'node /app/index.js')
  repo = local_git_repo('.')
  add(repo.path('letters/src'), '/app')
  add(repo.path('letters/package.json'), '/app/package.json')
  add(repo.path('letters/yarn.lock'), '/app/yarn.lock')

  run('cd /app && yarn install', trigger=['letters/package.json', 'letters/yarn.lock'])
  img = stop_build()

  s = k8s_service(img, yaml=yaml)
  s.port_forward(9001)
  return s


def numbers():
  yaml = m4_yaml('numbers/deployments/numbers.yaml')

  image_name = 'gcr.io/windmill-public-containers/abc123/numbers'

  start_fast_build('Dockerfile.py.base', image_name)
  repo = local_git_repo('.')
  add(repo.path('numbers'), "/app")

  run('cd /app && pip install -r requirements.txt', trigger='numbers/requirements.txt')
  img = stop_build()

  s = k8s_service(img, yaml=yaml)
  s.port_forward(9002)
  return s


