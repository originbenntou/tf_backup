import sqlalchemy
from sqlalchemy import create_engine
from sqlalchemy.orm import sessionmaker
from .config import *

pool = create_engine(
    # engine.url.URL(
    #     drivername="mysql+pymysql",
    #     username=user,
    #     password=password,
    #     database=db_name,
    #     query={
    #         "unix_socket": "{}/{}".format(
    #             db_socket_dir,
    #             cloud_sql_connection_name)
    #     },
    'mysql+pymysql://{user}:{passwd}@/{db}?charset=utf8&unix_socket={db_socket_dir}/{cloud_sql_connection_name}'.format(
        user=user,
        passwd=password,
        db=db_name,
        db_socket_dir=db_socket_dir,
        cloud_sql_connection_name=cloud_sql_connection_name
    ),
    encoding='utf-8',
    echo=False
)

Session = sessionmaker(autocommit=False, autoflush=False, bind=pool)
session = Session()
