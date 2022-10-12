# Front-end build stage
FROM alpine:3.16.2 as build-stage

# Install dependencies only
RUN apk add yarn
COPY package.json /frontend/
COPY yarn.lock /frontend
RUN cd /frontend && yarn install

# Setup the rest of the project to be executed
COPY . /frontend
WORKDIR /frontend
RUN yarn build

# Front-end execution stage
FROM nginx:1.23.1-alpine
COPY --from=build-stage /frontend/dist /app
COPY --from=build-stage /frontend/nginx.conf /etc/nginx/nginx.conf
