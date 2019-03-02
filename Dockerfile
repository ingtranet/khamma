FROM ubuntu:bionic

ENV LD_LIBRARY_PATH /usr/local/lib:${LD_LIBRARY_PATH}

RUN apt-get update -y && apt-get install -y python3-minimal cmake build-essential wget git golang-go locales
RUN locale-gen en_US.UTF-8 && update-locale LANG=en_US.UTF-8

WORKDIR /app
RUN git clone https://github.com/kakao/khaiii.git

WORKDIR /app/khaiii
RUN mkdir build

WORKDIR /app/khaiii/build
RUN cmake ..
RUN make all
RUN make large_resource
RUN make install

RUN rm -rf /root/.hunter
RUN rm -rf /app/khaiii

WORKDIR /app/khamma
COPY . .
RUN go test -bench=.

