FROM us.gcr.io/constellation-utils/golang-alp:latest

#Add any additional tools or libraries that your build image would need

#Install Cloud SQL Proxy
RUN curl -o cloud_sql_proxy https://dl.google.com/cloudsql/cloud_sql_proxy.linux.amd64 \
    && chmod +x cloud_sql_proxy \
    && mv cloud_sql_proxy /usr/local/bin