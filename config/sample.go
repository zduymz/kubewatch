package config

var yannotated = `# Handlers know how to send notifications to specific services.
handler:
  slack:
    # Slack "legacy" API token.
    token: ""
    # Slack channel.
    channel: ""
    # Title of the message.
    title: ""
# Resources to watch.
resource:
  deployment: false
  rc: false
  rs: false
  ds: false
  svc: false
  po: false
  job: false
  node: false
  clusterrole: false
  sa: false
  pv: false
  ns: false
  secret: false
  configmap: false
  ing: false
# For watching specific namespace, leave it empty for watching all.
# this config is ignored when watching namespaces
namespace: ""
`
