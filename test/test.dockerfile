
FROM golang:1.20.4-bullseye AS testing

WORKDIR /setup
COPY /setup/ ./
RUN ./install_bazel.sh

WORKDIR /
COPY . .
RUN bazel test //test/...
RUN echo success > /out

FROM scratch
COPY --from=testing /out /out
