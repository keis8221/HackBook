FROM node:19.8.1-alpine

WORKDIR /frontend
COPY . /frontend

ENV LANG C.UTF-8
ENV TZ Asia/Tokyo

ENV NODE_OPTIONS --max-old-space-size=4096
RUN yarn install

EXPOSE 3000

RUN chmod +x frontend/start.sh
CMD ["frontend/start.sh"]
