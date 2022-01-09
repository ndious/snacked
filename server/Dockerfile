FROM ruby:3.1.0-slim-bullseye

RUN apt-get update && apt-get install -y --no-install-recommends \
    build-essential \
    libpq-dev \
    postgresql-client \
    git \
  && rm -rf /var/lib/apt/lists/* /tmp/* /var/tmp/*

WORKDIR /snacked
COPY ./Gemfile /snacked/Gemfile
COPY ./Gemfile.lock /snacked/Gemfile.lock

RUN bundle install

COPY ./entrypoint.sh /usr/bin/
RUN echo $(ls -1 /snacked)
RUN chmod +x /usr/bin/entrypoint.sh
ENTRYPOINT ["entrypoint.sh"]
EXPOSE 3000

CMD ["bundle", "exec", "rails", "s", "-b", "0.0.0.0"]
