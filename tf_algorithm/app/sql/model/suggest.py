from sqlalchemy.ext.declarative import declarative_base
from sqlalchemy.schema import Column
from sqlalchemy.types import String, DateTime, Enum
from sqlalchemy.dialects.mysql import INTEGER as Integer
from datetime import datetime, timedelta, timezone

Base = declarative_base()


class SuggestModel(Base):
    __tablename__ = 'suggest'
    JST = timezone(timedelta(hours=+9), 'JST')

    id = Column('id', Integer(unsigned=True), primary_key=True, autoincrement=True, nullable=False)
    search_id = Column('search_id', Integer(unsigned=True), nullable=False)
    suggest_word = Column('suggest_word', String(255), nullable=False)
    created_at = Column('created_at', DateTime, default=datetime.now(JST), nullable=True)
    updated_at = Column('updated_at', DateTime, default=datetime.now(JST), nullable=True)
