FROM node:10-slim AS run-node

RUN npm install npm@6.4.1 -g

WORKDIR /home/app

COPY . .

RUN npm install

CMD ["node", "app.js"]