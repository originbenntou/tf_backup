FROM python:3.8

ENV PYTHONUNBUFFERED True

ENV WORK_DIR /app
WORKDIR $WORK_DIR

COPY . .

RUN pip3 install -r ./requirements.txt

CMD exec gunicorn --bind :$PORT --workers 1 --threads 8 --timeout 0 --reload main:app
