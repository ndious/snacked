FROM ruby:3.1.0-slim-bullseye

RUN apt-get update && apt-get install -y --no-install-recommends \
    build-essential \
    libpq-dev \
    postgresql-client \
    git \
  && rm -rf /var/lib/apt/lists/* /tmp/* /var/tmp/*

# RUN npm install --global yarn tod-cli

WORKDIR /snacked
COPY code/Gemfile /snacked/Gemfile
COPY code/Gemfile.lock /snacked/Gemfile.lock

RUN bundle install

COPY code/entrypoint.sh /usr/bin/
RUN chmod +x /usr/bin/entrypoint.sh
ENTRYPOINT ["entrypoint.sh"]
EXPOSE 3000

CMD ["bundle", "exec", "rails", "s", "-b", "0.0.0.0"]
