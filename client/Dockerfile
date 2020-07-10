# FROM mhart/alpine-node:11 AS builder
# ARG REACT_APP_URI_API
# ENV REACT_APP_URI_API $REACT_APP_URI_API
# WORKDIR /app
# COPY . .
# RUN npm run build

FROM mhart/alpine-node
RUN npm install -g serve
WORKDIR /app
COPY ./build .
CMD ["serve", "-p", "3000", "-s", "."]doc
EXPOSE 3000