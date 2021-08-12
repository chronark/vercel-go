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
    - [x] **v4** Remove a domain by name
    - [x] **v4** Check a domain availability
    - [x] **v4** Check the price of a domain
    - [x] **v4** Purchase a domain

- [ ] DNS    
    - [ ] List all the DNS records of a domain
    - [ ] Create a new DNS record
    - [ ] Remove a DNS record        

- [ ] Certificates   
    - [ ] List cerfiticates
    - [ ] Get a single certificate
    - [ ] Create a new certificate
    - [ ] Submit a certificate
    - [ ] Delete a certificate

- [ ] Aliases    
    - [ ] List all the aliases
    - [ ] Get a Single Alias
    - [ ] Remove an Alias
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
    - [ ] Remove user from team

- [ ] Projects
    - [ ] Create a Project
    - [ ] Get Projects
    - [ ] Get a Single Project
    - [ ] Update a Single Project
    - [ ] Remove a Single Project
    - [ ] Get Project Environment Variables
    - [ ] Create a Project Environment Variable
    - [ ] Edit a Project Environment Variable
    - [ ] Delete a Specific Environment Variable
    - [ ] Add a Domain to a Project
    - [ ] Get Project Domains
    - [ ] Get a Single Project Domain
    - [ ] Set Redirect for a Project Domain
    - [ ] Delete a Specific Production Domain of a Project

- [ ] Authentication
    - [ ] Request a login
    - [ ] Verify login