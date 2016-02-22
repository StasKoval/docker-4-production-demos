FROM ruby:2.3.0-alpine

COPY . /app
WORKDIR /app
RUN bundle install
EXPOSE 3000
USER nobody
CMD rackup -o 0.0.0.0 -p 3000
