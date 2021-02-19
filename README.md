# A GO web service

This is a sample GO web service created using GIN framework. The app run on port :3000

## Security

Implements Okta oauth security

## Endpoints

- **GET**: **"/"**
  - unprotected
- **GET**: **"/students"**
  - protected

## Note

You will need to get an access token to call protected endpoint. Add Okta oauth configs in **"config/config.properties"**