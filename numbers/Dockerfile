FROM python:3.7

ARG ENABLE_DEBUGGING
RUN if [ "x$ENABLE_DEBUGGING" != "x" ]; then pip install web_pdb; fi

ADD . /app
RUN cd /app && pip install -r requirements.txt

ENTRYPOINT ["./app.py"]
