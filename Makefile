.PHONY: deploy update

update:
	gcloud functions deploy discord-alert --source .

# only call to reconstruct things
deploy:
	gcloud functions deploy \
		discord-alert \
		--gen2 \
		--runtime go122 \
		--region us-central1 \
		--service-account=alerting-discord@member-gentei.iam.gserviceaccount.com \
		--trigger-topic=${TRIGGER_TOPIC} \
		--set-secrets=DISCORD_WEBHOOK_URL=alerting-discord-webhook-url:latest \
		--source . \
		--entry-point HandlePubSubAlert