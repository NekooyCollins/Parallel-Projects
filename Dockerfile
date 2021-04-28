FROM tudinfse/cds_server

# Add this line for test
# Add new path in cds_server.json
COPY ./cds_server.json /etc/cds_server.json

# Add C++ compilation dependencies
RUN apt-get -y update && \
    apt-get install -y && \
    apt-get -y install gcc g++ make gdb wget curl git golang

# # Copy source code and compile it
# COPY ./mopp-2018-t0-harmonic-progression-sum /mopp-2018-t0-harmonic-progression-sum

# RUN cd mopp-2018-t0-harmonic-progression-sum && \
#     make && \
#     cp ./harmonic-progression-sum /usr/bin/

RUN mkdir code
COPY ./mandelbrot code/mandelbrot
COPY ./levenshtein code/levenshtein


RUN cd /code/mandelbrot && \
    make && \
    cp ./mandelbrot /usr/bin/

# RUN cd /code/levenshtein && \
#     make && \
#     cp ./lev /usr/bin/

RUN cd /code/levenshtein && \
    go build -o lev_parallel . && \
    cp ./lev_parallel /usr/bin/




