# Build Stage 1
# This build created a staging docker image 
#
FROM node:10.15.2-alpine AS appbuild
WORKDIR /usr/src/app
COPY package.json ./
RUN npm install
COPY ./src ./src

# Build Stage 2
# This build takes the production build from staging build
#
FROM node:10.15.2-alpine as production
WORKDIR /usr/src/app
COPY package.json ./
RUN npm install
COPY --from=appbuild /usr/src/app/ ./
EXPOSE 4002
CMD npm run production