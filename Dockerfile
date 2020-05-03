FROM yanrishbe/gaming-website:v1

WORKDIR /bin

COPY .env .
COPY bin .

ENTRYPOINT [ "sh", "-c"]
CMD ["/bin/app"]