
# # production stage
FROM nginx:stable-alpine as production-stage
COPY ./build /usr/share/nginx/html
COPY nginx.conf /etc/nginx/conf.d/default.conf
COPY entrypoint.sh /usr/bin/

RUN chmod +x /usr/bin/entrypoint.sh

EXPOSE 80

CMD ["/usr/bin/entrypoint.sh"]
