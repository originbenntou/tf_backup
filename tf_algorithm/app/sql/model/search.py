from sqlalchemy.ext.declarative import declarative_base
from sqlalchemy.schema import Column
from sqlalchemy.types import String, DateTime, Enum
from sqlalchemy.dialects.mysql import INTEGER as Integer
from datetime import datetime, timedelta, timezone

Base = declarative_base()


class SearchModel(Base):
    __tablename__ = 'search'
    JST = timezone(timedelta(hours=+9), 'JST')

    id = Column('id', Integer(unsigned=True), primary_key=True, autoincrement=True, nullable=False)
    search_word = Column('search_word', String(255), nullable=False)
    date = Column('date', DateTime, default=datetime.now(JST), nullable=False)
    status = Column('status', Integer(unsigned=True), nullable=False)
    created_at = Column('created_at', DateTime, default=datetime.now(JST), nullable=True)
    updated_at = Column('updated_at', DateTime, default=datetime.now(JST), nullable=True)
