# Stage - Installing dep
FROM node:15-alpine3.13 as dependencies
WORKDIR /webapp
RUN npm set progress=false && npm config set depth 0
COPY package.json package-lock.json ./
RUN npm install

# Stage - build
FROM node:15-alpine3.13 as build
WORKDIR /webapp
COPY --from=dependencies /webapp/node_modules ./node_modules
COPY . .
RUN npm run build

# Stage - release
FROM nginx:stable-alpine as release
WORKDIR /opt/webapp
COPY --from=build /webapp/conf/nginx /etc/nginx/
COPY --from=build /webapp/generic/errors /usr/local/nginx/errors
COPY --from=build /webapp/build/. /usr/share/nginx/html
EXPOSE 8080