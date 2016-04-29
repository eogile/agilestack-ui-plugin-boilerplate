FROM centurylink/ca-certs
MAINTAINER EOGILE "agilestack@eogile.com"

ENV name plugin-ui-boilerplate

ENV workdir /plugin
ENV builddir $workdir/build

ENV baseUrl plugin-ui-boilerplate
EXPOSE 8081

LABEL SERVICE_TAGS="urlprefix-/$baseUrl" \
      SERVICE_CHECK_HTTP="/stats" \
      SERVICE_CHECK_INTERVAL="10s"

WORKDIR $workdir
ADD $name $workdir/$name
ADD build $builddir

ENTRYPOINT ["./plugin-ui-boilerplate"]
