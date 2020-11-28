FROM mysql:5.7

COPY ./mysql/conf.d/my.cnf /etc/mysql/conf.d/my.cnf
RUN chmod 644 /etc/mysql/conf.d/*

COPY ./mysql/init/* /docker-entrypoint-initdb.d/

RUN apt-get update && \
  apt-get install -y locales && \
  rm -rf /var/lib/apt/lists/* && \
  echo "ja_JP.UTF-8 UTF-8" > /etc/locale.gen && \
  locale-gen ja_JP.UTF-8
ENV LC_ALL ja_JP.UTF-8
