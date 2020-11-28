from sqlalchemy.ext.declarative import declarative_base
from sqlalchemy.schema import Column
from sqlalchemy.types import String, DateTime, Enum
from sqlalchemy.dialects.mysql import INTEGER as Integer
from datetime import datetime, timedelta, timezone

Base = declarative_base()


class ChildSuggestModel(Base):
    __tablename__ = 'child_suggest'
    JST = timezone(timedelta(hours=+9), 'JST')

    id = Column('id', Integer(unsigned=True), primary_key=True, autoincrement=True, nullable=False)
    suggest_id = Column('suggest_id', Integer(unsigned=True), nullable=False)
    child_suggest_word = Column('child_suggest_word', String(255), nullable=False)
    short = Column('short', Integer(unsigned=True), nullable=False)
    medium = Column('medium', Integer(unsigned=True), nullable=False)
    long = Column('long', Integer(unsigned=True), nullable=False)
    short_graphs = Column('short_graphs', String(2055), nullable=False)
    medium_graphs = Column('medium_graphs', String(2055), nullable=False)
    long_graphs = Column('long_graphs', String(2055), nullable=False)
    created_at = Column('created_at', DateTime, default=datetime.now(JST), nullable=True)
    updated_at = Column('updated_at', DateTime, default=datetime.now(JST), nullable=True)
