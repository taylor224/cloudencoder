INSERT INTO public.settings_option (id, name, description, title, secure) VALUES (1, 'S3_ACCESS_KEY', 'S3 Access Key', 'S3 Access Key', true);
INSERT INTO public.settings_option (id, name, description, title, secure) VALUES (2, 'S3_SECRET_KEY', 'S3 Secret Key', 'S3 Secret Key', true);
INSERT INTO public.settings_option (id, name, description, title, secure) VALUES (3, 'DIGITAL_OCEAN_ACCESS_TOKEN', 'Digital Ocean Access Token (Required for Machines)', 'Digital Ocean Access Token', true);
INSERT INTO public.settings_option (id, name, description, title, secure) VALUES (4, 'SLACK_WEBHOOK', 'Slack Webhook for notifications', 'Slack Webhook', true);
INSERT INTO public.settings_option (id, name, description, title, secure) VALUES (5, 'S3_INBOUND_BUCKET', 'S3 Inbound Bucket', 'S3 Inbound Bucket', false);
INSERT INTO public.settings_option (id, name, description, title, secure) VALUES (6, 'S3_OUTBOUND_BUCKET', 'S3 Outbound Bucket', 'S3 Outbound Bucket', false);
INSERT INTO public.settings_option (id, name, description, title, secure) VALUES (7, 'S3_OUTBOUND_BUCKET_REGION', 'S3 Outbound Bucket Region', 'S3 Outbound Bucket Region', false);
INSERT INTO public.settings_option (id, name, description, title, secure) VALUES (8, 'S3_INBOUND_BUCKET_REGION', 'S3 Inbound Bucket Region', 'S3 Inbound Bucket Region', false);
INSERT INTO public.settings_option (id, name, description, title, secure) VALUES (9, 'S3_PROVIDER', 'S3 Provider', 'S3 Provider', false);
INSERT INTO public.settings_option (id, name, description, title, secure) VALUES (10, 'S3_STREAMING', 'Enable this setting to enable streaming directly to FFmpeg from a pre-signed S3 URL, instead of downloading the file first if disk space is a concern. Please note this setting can impact performance.', 'Stream Encode from S3', false);
INSERT INTO public.settings_option (id, name, description, title, secure) VALUES (11, 'FTP_ADDR', 'FTP Connection', 'FTP Address', false);
INSERT INTO public.settings_option (id, name, description, title, secure) VALUES (12, 'FTP_USERNAME', 'FTP Username', 'FTP Username', true);
INSERT INTO public.settings_option (id, name, description, title, secure) VALUES (13, 'FTP_PASSWORD', 'FTP Password', 'FTP Password', true);
INSERT INTO public.settings_option (id, name, description, title, secure) VALUES (14, 'STORAGE_DRIVER', 'Storage Driver for input and output', 'Storage Driver', false);

SELECT setval('settings_option_id_seq', max(id)) FROM settings_option;