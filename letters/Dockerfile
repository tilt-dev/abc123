FROM node:10

WORKDIR /app
ADD package.json package.json
ADD yarn.lock yarn.lock

RUN yarn install

ADD src .

ENTRYPOINT [ "node", "/app/index.js" ]
