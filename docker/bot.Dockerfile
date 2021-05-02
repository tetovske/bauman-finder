FROM default AS builder

WORKDIR /baumanfinder

RUN chmod ugo+x .bin/bot

CMD [".bin/bot"]