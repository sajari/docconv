# Dockerfile which produces an AppEngine custom runtime containing docd and all
# of its runtime dependencies.
# https://cloud.google.com/appengine/docs/flexible/custom-runtimes/about-custom-runtimes
FROM sajari/docd:1.3.7
CMD ["-addr", ":8080", "-json-cloud-logging", "-error-reporting"]
