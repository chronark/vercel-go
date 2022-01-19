# Vercel Go Client

[![codecov](https://codecov.io/gh/chronark/vercel-go/branch/main/graph/badge.svg?token=GYF8ON46AV)](https://codecov.io/gh/chronark/vercel-go)

This client implements the latest endpoints from the [vercel api](https://vercel.com/docs/api)

The main application of this library is in a [terraform provider](https://github.com/chronark/terraform-provider-vercel) but it might be useful to you in your projects.

## Endpoints:

- [x] User
    - [x] Get the authenticated user
- Deployments    
    - [ ] Create a New Deployment
    - [ ] Upload Deployment Files
    - [ ] List Deployments
    - [ ] Get a Single Deployment
    - [ ] Delete a Deployment
    - [ ] List Deployment Files
    - [ ] Get Single File Contents
    - [ ] List Builds
    - [ ] Cancel a Deployment
- [ ] Logs     
    - [ ] Get Build Logs
    - [ ] Stream Serverless Function Logs
    - [ ] Fetch Failed Requests for Serverless Function      
- [x] Domains
    - [x] **v5** List all the domains
    - [x] **v4** Add a New Domain
    - [x] **v4** Transfer In a Domain
    - [x] **v4** Verify a Domain
    - [x] **v4** Get Information for a Single Domain
    - [x] **v6** Get Auth Code for a Single Domain
    - [x] **v4** Delete a domain by name
    - [x] **v4** Check a domain availability
    - [x] **v4** Check the price of a domain
    - [x] **v4** Purchase a domain

- [x] DNS    
    - [x] **v4** List all the DNS records of a domain
    - [x] **v2** Create a new DNS record
    - [x] **v2** Delete a DNS record        

- [ ] Certificates   
    - [ ] List cerfiticates
    - [ ] Get a single certificate
    - [ ] Create a new certificate
    - [ ] Submit a certificate
    - [ ] Delete a certificate

- [ ] Aliases    
    - [ ] List all the aliases
    - [ ] Get a Single Alias
    - [ ] Delete an Alias
    - [ ] Purge Alias in the CDN
    - [ ] List aliases by deployment
    - [ ] Assign an alias to a deployment
    
- [ ] Secrets 
    - [x] **v3** List secrets
    - [x] **v3** Get a Single Secret
    - [x] **v2** Create a new secret
    - [x] **v2** Change secret name
    - [x] **v2** Delete a secret

- [ ] Teams    
    - [ ] Create a Team
    - [ ] Delete a Team
    - [ ] List all your teams
    - [x] **v1** Get single team information
    - [ ] Update team information
    - [ ] Get list of team members
    - [ ] Invite user to team
    - [ ] Change user role or status
    - [ ] Request to join to a team
    - [ ] Check status of a join request
    - [ ] Delete user from team

- [ ] Projects
    - [x] **v8** Create a Project
    - [x] **v8** Get Projects
    - [x] **v8** Get a Single Project
    - [x] **v8** Update a Single Project
    - [x] **v8** Delete a Single Project
    - [ ] Get Project Environment Variables
    - [x] **v8** Create a Project Environment Variable
    - [x] **v8** Edit a Project Environment Variable
    - [x] **v8** Delete a Specific Environment Variable
    - [x] **v8** Add a Domain to a Project
    - [x] **v8** Get Project Domains
    - [x] **v8** Get a Single Project Domain
    - [x] **v8** Update a Project Domain
    - [ ] Delete a Specific Production Domain of a Project

- [ ] Authentication
    - [ ] Request a login
    - [ ] Verify login