FROM node:19.8.1-alpine

WORKDIR /app
COPY . /app

ENV LANG C.UTF-8
ENV TZ Asia/Tokyo

ENV NODE_OPTIONS --openssl-legacy-provider
RUN yarn install

EXPOSE 3000

RUN chmod +x frontend/start.sh
CMD ["frontend/start.sh"]
