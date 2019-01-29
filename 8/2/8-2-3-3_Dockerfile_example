FROM node:10-alpine

RUN apk update && apk upgrade \
    && apk add --no-cache bash git openssh make gcc g++ python
WORKDIR /usr/app
EXPOSE 3000
COPY . .
RUN yarn
CMD ["yarn", "start"]
