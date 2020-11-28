import os

# default = local
host = "mysql"
db_name = "trend"
user = "2929"
password = "2929"
db_socket_dir = ""
cloud_sql_connection_name = ""

env = os.environ['ENV']

if env == "LOCAL":
    host = "mysql"
    db_name = "trend"
    user = "2929"
    password = "2929"
    db_socket_dir = ""
    cloud_sql_connection_name = ""
elif env == "PRD":
    host = "127.0.0.1"
    db_name = "trend"
    user = "2929"
    password = "dmeej1pu4Cos21gO"
    db_socket_dir = "/cloudsql"
    cloud_sql_connection_name = "aqueous-nebula-278307:asia-northeast1:trend-finder-mysql"
