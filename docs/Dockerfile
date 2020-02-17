FROM jekyll/jekyll

LABEL maintainer="Jeremías Casteglione <jrmsdev@gmail.com>"
LABEL version="2020.02.17"

RUN gem install github-pages

COPY --chown=jekyll:jekyll Gemfile /srv/jekyll/

RUN (cd /srv/jekyll && jekyll build) && rm -vf /srv/jekyll/Gemfile
