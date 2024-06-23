# gentei-alerting

Gentei's runtime monitoring uses log-based metrics in GCP Operations, which unfortunately means that push notifications to my phone are just `Log match condition fired for Google Project with {project_id=member-gentei}.`.

This is the codebase for a [Google Cloud function](https://cloud.google.com/functions) that (soon) processes the alert and posts the actual logs that are causing problems to a Discord server + channel that only I have access to.
