FROM mysql:8.0.32-1.2.11-server

RUN chown -R mysql:root /var/lib/mysql/

#ARG MYSQL_DATABASE
#ARG MYSQL_USER
#ARG MYSQL_PASSWORD
#ARG MYSQL_ROOT_PASSWORD

ENV MYSQL_DATABASE=FLORA
ENV MYSQL_USER=floraadm
ENV MYSQL_PASSWORD=floraadm
ENV MYSQL_ROOT_PASSWORD=totallynotadmin123

# Copy existing data from local file data.sql to container
# ADD data.sql /etc/mysql/data.sql
# RUN sed -i 's/MYSQL_DATABASE/'$MYSQL_DATABASE'/g' /etc/mysql/data.sql
# RUN cp /etc/mysql/data.sql /docker-entrypoint-initdb.d

EXPOSE 3306
EXPOSE 33060
EXPOSE 33061

# Create 
# docker run -d -p 3306:3306 --name mysql-docker-container -e MYSQL_ROOT_PASSWORD=totallynotadmin123 -e MYSQL_DATABASE=flora -e MYSQL_USER=floraadmin -e MYSQL_PASSWORD=floraadmin mysql/mysql-server:8.0.32-1.2.11-server

# Run container
# docker exec -it mysql-docker-container bash

# Access DB from bash terminal
# mysql -u floraadmin -p
